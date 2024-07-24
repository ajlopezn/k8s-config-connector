// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockbigquery

import (
	"context"
	"net/http"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/bigquery/v2"
	"github.com/golang/protobuf/ptypes/empty"
)

type tablesServer struct {
	*MockService
	pb.UnimplementedTablesServerServer
}

func (s *tablesServer) GetTable(ctx context.Context, req *pb.GetTableRequest) (*pb.Table, error) {
	name, err := s.buildTableName(req.GetProjectId(), req.GetDatasetId(), req.GetTableId())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Table{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: Table %s:%s", name.Project.ID, name.TableID)
		}
		return nil, err
	}

	return obj, nil
}

func (s *tablesServer) InsertTable(ctx context.Context, req *pb.InsertTableRequest) (*pb.Table, error) {
	name, err := s.buildTableName(req.GetProjectId(), req.GetDatasetId(), req.GetTable().GetTableReference().GetTableId())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetTable()).(*pb.Table)

	if obj.TableReference == nil {
		obj.TableReference = &pb.TableReference{}
	}
	if obj.GetTableReference().GetProjectId() == "" {
		obj.TableReference.ProjectId = req.ProjectId
	}
	obj.CreationTime = PtrTo(now.UnixMilli())
	obj.LastModifiedTime = PtrTo(uint64(now.UnixMilli()))
	obj.Id = PtrTo(obj.GetTableReference().GetProjectId() + ":" + obj.GetTableReference().GetTableId())
	obj.Kind = PtrTo("bigquery#table")
	if obj.Location == nil {
		obj.Location = PtrTo("US")
	}
	if obj.Type == nil {
		obj.Type = PtrTo("DEFAULT")
	}

	obj.SelfLink = PtrTo("https://bigquery.googleapis.com/bigquery/v2/" + name.String())

	//sortAccess(obj)

	obj.Etag = PtrTo(computeEtag(obj))

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating Table: %v", err)
	}

	return obj, nil
}

func (s *tablesServer) UpdateTable(ctx context.Context, req *pb.UpdateTableRequest) (*pb.Table, error) {
	name, err := s.buildTableName(req.GetProjectId(), req.GetDatasetId(), req.GetTableId())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Table{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	now := time.Now()

	updated := req.GetTable()
	updated.TableReference = existing.TableReference

	updated.CreationTime = existing.CreationTime
	updated.LastModifiedTime = PtrTo(uint64(now.UnixMilli()))
	updated.Id = PtrTo(existing.GetTableReference().GetProjectId() + ":" + existing.GetTableReference().GetTableId())
	updated.Kind = PtrTo("bigquery#Table")
	updated.Location = existing.Location
	updated.Type = existing.Type
	updated.SelfLink = PtrTo("https://bigquery.googleapis.com/bigquery/v2/" + name.String())

	//sortAccess(updated)

	updated.Etag = PtrTo(computeEtag(updated))

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, err
}

func (s *tablesServer) DeleteTable(ctx context.Context, req *pb.DeleteTableRequest) (*empty.Empty, error) {
	name, err := s.buildTableName(req.GetProjectId(), req.GetDatasetId(), req.GetTableId())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Table{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	httpmux.SetStatusCode(ctx, http.StatusNoContent)

	return &empty.Empty{}, nil
}

type tableName struct {
	Project   *projects.ProjectData
	DatasetID string
	TableID   string
}

func (n *tableName) String() string {
	return "projects/" + n.Project.ID + "/datasets/" + n.DatasetID + "/tables/" + n.TableID
}

// parseTableName parses a string into a tableName.
// The expected form is projects/<projectID>/datasets/<DatasetID>/tables/<TableID> --> TODO
func (s *MockService) parseTableName(name string) (*tableName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "datasets" && tokens[4] == "tables" {
		return s.buildTableName(tokens[1], tokens[3], tokens[5])
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *MockService) buildTableName(projectName string, datasetID string, tableID string) (*tableName, error) {
	project, err := s.Projects.GetProjectByID(projectName)
	if err != nil {
		return nil, err
	}

	name := &tableName{
		Project:   project,
		DatasetID: datasetID,
		TableID:   tableID,
	}

	return name, nil
}
