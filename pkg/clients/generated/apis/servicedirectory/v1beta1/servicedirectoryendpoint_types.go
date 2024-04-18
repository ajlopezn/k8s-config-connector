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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ServiceDirectoryEndpointSpec struct {
	// +optional
	AddressRef *v1alpha1.ResourceRef `json:"addressRef,omitempty"`

	/* Only the `external` field is supported to configure the reference.

	Immutable. The Google Compute Engine network (VPC) of the endpoint in the format
	projects/<project number>/locations/global/networks/*.

	The project must be specified by project number (project id is rejected). Incorrectly formatted networks are
	rejected, but no other validation is performed on this field (ex. network or project existence,
	reachability, or permissions). */
	// +optional
	NetworkRef *v1alpha1.ResourceRef `json:"networkRef,omitempty"`

	/* Port that the endpoint is running on, must be in the
	range of [0, 65535]. If unspecified, the default is 0. */
	// +optional
	Port *int64 `json:"port,omitempty"`

	/* Immutable. Optional. The endpointId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The ServiceDirectoryService that this endpoint belongs to. */
	ServiceRef v1alpha1.ResourceRef `json:"serviceRef"`
}

type ServiceDirectoryEndpointStatus struct {
	/* Conditions represent the latest available observations of the
	   ServiceDirectoryEndpoint's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The resource name for the endpoint in the format
	'projects/* /locations/* /namespaces/* /services/* /endpoints/*'. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpservicedirectoryendpoint;gcpservicedirectoryendpoints
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ServiceDirectoryEndpoint is the Schema for the servicedirectory API
// +k8s:openapi-gen=true
type ServiceDirectoryEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceDirectoryEndpointSpec   `json:"spec,omitempty"`
	Status ServiceDirectoryEndpointStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceDirectoryEndpointList contains a list of ServiceDirectoryEndpoint
type ServiceDirectoryEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceDirectoryEndpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServiceDirectoryEndpoint{}, &ServiceDirectoryEndpointList{})
}
