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

type WorkstationclusterPrivateClusterConfig struct {
	/* Additional project IDs that are allowed to attach to the workstation cluster's service attachment.
	By default, the workstation cluster's project and the VPC host project (if different) are allowed. */
	// +optional
	AllowedProjects []string `json:"allowedProjects,omitempty"`

	/* Hostname for the workstation cluster.
	This field will be populated only when private endpoint is enabled.
	To access workstations in the cluster, create a new DNS zone mapping this domain name to an internal IP address and a forwarding rule mapping that address to the service attachment. */
	// +optional
	ClusterHostname *string `json:"clusterHostname,omitempty"`

	/* Immutable. Whether Workstations endpoint is private. */
	EnablePrivateEndpoint bool `json:"enablePrivateEndpoint"`

	/* Service attachment URI for the workstation cluster.
	The service attachment is created when private endpoint is enabled.
	To access workstations in the cluster, configure access to the managed service using (Private Service Connect)[https://cloud.google.com/vpc/docs/configure-private-service-connect-services]. */
	// +optional
	ServiceAttachmentUri *string `json:"serviceAttachmentUri,omitempty"`
}

type WorkstationsWorkstationClusterSpec struct {
	/* Client-specified annotations. This is distinct from labels. */
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`

	/* Human-readable name for this resource. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* Immutable. The location where the workstation cluster should reside. */
	Location string `json:"location"`

	/* Immutable. The relative resource name of the VPC network on which the instance can be accessed.
	It is specified in the following form: "projects/{projectNumber}/global/networks/{network_id}". */
	Network string `json:"network"`

	/* Configuration for private cluster. */
	// +optional
	PrivateClusterConfig *WorkstationclusterPrivateClusterConfig `json:"privateClusterConfig,omitempty"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The workstationClusterId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. Name of the Compute Engine subnetwork in which instances associated with this cluster will be created.
	Must be part of the subnetwork specified for this cluster. */
	Subnetwork string `json:"subnetwork"`
}

type WorkstationclusterDetailsStatus struct {
}

type WorkstationclusterResourceConditionsStatus struct {
	/* The status code, which should be an enum value of google.rpc.Code. */
	// +optional
	Code *int64 `json:"code,omitempty"`

	/* A list of messages that carry the error details. */
	// +optional
	Details []WorkstationclusterDetailsStatus `json:"details,omitempty"`

	/* Human readable message indicating details about the current status. */
	// +optional
	Message *string `json:"message,omitempty"`
}

type WorkstationsWorkstationClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   WorkstationsWorkstationCluster's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Time when this resource was created. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* Whether this resource is in degraded mode, in which case it may require user action to restore full functionality.
	Details can be found in the conditions field. */
	// +optional
	Degraded *bool `json:"degraded,omitempty"`

	/* Checksum computed by the server.
	May be sent on update and delete requests to ensure that the client has an up-to-date value before proceeding. */
	// +optional
	Etag *string `json:"etag,omitempty"`

	/* The name of the cluster resource. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* Status conditions describing the current resource state. */
	// +optional
	ResourceConditions []WorkstationclusterResourceConditionsStatus `json:"resourceConditions,omitempty"`

	/* The system-generated UID of the resource. */
	// +optional
	Uid *string `json:"uid,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpworkstationsworkstationcluster;gcpworkstationsworkstationclusters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=alpha";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// WorkstationsWorkstationCluster is the Schema for the workstations API
// +k8s:openapi-gen=true
type WorkstationsWorkstationCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkstationsWorkstationClusterSpec   `json:"spec,omitempty"`
	Status WorkstationsWorkstationClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkstationsWorkstationClusterList contains a list of WorkstationsWorkstationCluster
type WorkstationsWorkstationClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkstationsWorkstationCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WorkstationsWorkstationCluster{}, &WorkstationsWorkstationClusterList{})
}
