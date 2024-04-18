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

type PrivateconnectionVpcPeeringConfig struct {
	/* Immutable. A free subnet for peering. (CIDR of /29). */
	Subnet string `json:"subnet"`

	/* Immutable. Fully qualified name of the VPC that Datastream will peer to.
	Format: projects/{project}/global/{networks}/{name}. */
	Vpc string `json:"vpc"`
}

type DatastreamPrivateConnectionSpec struct {
	/* Immutable. Display name. */
	DisplayName string `json:"displayName"`

	/* Immutable. The name of the location this private connection is located in. */
	Location string `json:"location"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The privateConnectionId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The VPC Peering configuration is used to create VPC peering
	between Datastream and the consumer's VPC. */
	VpcPeeringConfig PrivateconnectionVpcPeeringConfig `json:"vpcPeeringConfig"`
}

type PrivateconnectionErrorStatus struct {
	/* A list of messages that carry the error details. */
	// +optional
	Details map[string]string `json:"details,omitempty"`

	/* A message containing more information about the error that occurred. */
	// +optional
	Message *string `json:"message,omitempty"`
}

type DatastreamPrivateConnectionStatus struct {
	/* Conditions represent the latest available observations of the
	   DatastreamPrivateConnection's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The PrivateConnection error in case of failure. */
	// +optional
	Error []PrivateconnectionErrorStatus `json:"error,omitempty"`

	/* The resource's name. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* State of the PrivateConnection. */
	// +optional
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatastreamprivateconnection;gcpdatastreamprivateconnections
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=alpha";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DatastreamPrivateConnection is the Schema for the datastream API
// +k8s:openapi-gen=true
type DatastreamPrivateConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatastreamPrivateConnectionSpec   `json:"spec,omitempty"`
	Status DatastreamPrivateConnectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DatastreamPrivateConnectionList contains a list of DatastreamPrivateConnection
type DatastreamPrivateConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatastreamPrivateConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatastreamPrivateConnection{}, &DatastreamPrivateConnectionList{})
}
