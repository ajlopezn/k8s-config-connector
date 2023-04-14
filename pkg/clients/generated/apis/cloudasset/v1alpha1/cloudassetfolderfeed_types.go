// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type FolderfeedCondition struct {
	/* Description of the expression. This is a longer text which describes the expression,
	e.g. when hovered over it in a UI. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Textual representation of an expression in Common Expression Language syntax. */
	Expression string `json:"expression"`

	/* String indicating the location of the expression for error reporting, e.g. a file
	name and a position in the file. */
	// +optional
	Location *string `json:"location,omitempty"`

	/* Title for the expression, i.e. a short string describing its purpose.
	This can be used e.g. in UIs which allow to enter the expression. */
	// +optional
	Title *string `json:"title,omitempty"`
}

type FolderfeedFeedOutputConfig struct {
	/* Destination on Cloud Pubsub. */
	PubsubDestination FolderfeedPubsubDestination `json:"pubsubDestination"`
}

type FolderfeedPubsubDestination struct {
	/* Destination on Cloud Pubsub topic. */
	Topic string `json:"topic"`
}

type CloudAssetFolderFeedSpec struct {
	/* A list of the full names of the assets to receive updates. You must specify either or both of
	assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
	exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
	See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info. */
	// +optional
	AssetNames []string `json:"assetNames,omitempty"`

	/* A list of types of the assets to receive updates. You must specify either or both of assetNames
	and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
	the feed. For example: "compute.googleapis.com/Disk"
	See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
	supported asset types. */
	// +optional
	AssetTypes []string `json:"assetTypes,omitempty"`

	/* Immutable. The project whose identity will be used when sending messages to the
	destination pubsub topic. It also specifies the project for API
	enablement check, quota, and billing. */
	BillingProject string `json:"billingProject"`

	/* A condition which determines whether an asset update should be published. If specified, an asset
	will be returned only when the expression evaluates to true. When set, expression field
	must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
	expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
	condition are optional. */
	// +optional
	Condition *FolderfeedCondition `json:"condition,omitempty"`

	/* Asset content type. If not specified, no content but the asset name and type will be returned. Possible values: ["CONTENT_TYPE_UNSPECIFIED", "RESOURCE", "IAM_POLICY", "ORG_POLICY", "ACCESS_POLICY"]. */
	// +optional
	ContentType *string `json:"contentType,omitempty"`

	/* Immutable. This is the client-assigned asset feed identifier and it needs to be unique under a specific parent. */
	FeedId string `json:"feedId"`

	/* Output configuration for asset feed destination. */
	FeedOutputConfig FolderfeedFeedOutputConfig `json:"feedOutputConfig"`

	/* Immutable. The folder this feed should be created in. */
	Folder string `json:"folder"`

	/* The folder that this resource belongs to. */
	FolderRef v1alpha1.ResourceRef `json:"folderRef"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type CloudAssetFolderFeedStatus struct {
	/* Conditions represent the latest available observations of the
	   CloudAssetFolderFeed's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The ID of the folder where this feed has been created. Both [FOLDER_NUMBER]
	and folders/[FOLDER_NUMBER] are accepted. */
	// +optional
	FolderId *string `json:"folderId,omitempty"`

	/* The format will be folders/{folder_number}/feeds/{client-assigned_feed_identifier}. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudAssetFolderFeed is the Schema for the cloudasset API
// +k8s:openapi-gen=true
type CloudAssetFolderFeed struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudAssetFolderFeedSpec   `json:"spec,omitempty"`
	Status CloudAssetFolderFeedStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudAssetFolderFeedList contains a list of CloudAssetFolderFeed
type CloudAssetFolderFeedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudAssetFolderFeed `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudAssetFolderFeed{}, &CloudAssetFolderFeedList{})
}
