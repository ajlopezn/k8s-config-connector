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

package mockgkehub

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	v1betapb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkehub/v1beta"
	v1beta1pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkehub/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// MockService represents a mocked gkehubfeature service.
type MockService struct {
	*common.MockEnvironment
	storage storage.Storage

	operations *operations.Operations

	v1beta  *GKEHubFeature
	v1beta1 *GKEHubMembership
}

// New creates a MockService.
func New(env *common.MockEnvironment, storage storage.Storage) *MockService {
	s := &MockService{
		MockEnvironment: env,
		storage:         storage,
		operations:      operations.NewOperationsService(storage),
	}
	s.v1beta = &GKEHubFeature{MockService: s}
	s.v1beta1 = &GKEHubMembership{MockService: s}
	return s
}

func (s *MockService) ExpectedHost() string {
	return "gkehub.googleapis.com"
}

func (s *MockService) Register(grpcServer *grpc.Server) {
	v1betapb.RegisterGkeHubServer(grpcServer, s.v1beta)
	v1beta1pb.RegisterGkeHubMembershipServiceServer(grpcServer, s.v1beta1)
}

func (s *MockService) NewHTTPMux(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux := runtime.NewServeMux()

	if err := v1betapb.RegisterGkeHubHandler(ctx, mux, conn); err != nil {
		return nil, err
	}
	if err := v1beta1pb.RegisterGkeHubMembershipServiceHandler(ctx, mux, conn); err != nil {
		return nil, err
	}
	return mux, nil
}
