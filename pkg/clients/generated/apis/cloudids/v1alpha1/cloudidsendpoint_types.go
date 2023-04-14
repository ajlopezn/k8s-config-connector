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

type CloudIDSEndpointSpec struct {
	/* Immutable. An optional description of the endpoint. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. The location for the endpoint. */
	Location string `json:"location"`

	/* Immutable. Name of the VPC network that is connected to the IDS endpoint. This can either contain the VPC network name itself (like "src-net") or the full URL to the network (like "projects/{project_id}/global/networks/src-net"). */
	Network string `json:"network"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The minimum alert severity level that is reported by the endpoint. Possible values: ["INFORMATIONAL", "LOW", "MEDIUM", "HIGH", "CRITICAL"]. */
	Severity string `json:"severity"`

	/* Configuration for threat IDs excluded from generating alerts. Limit: 99 IDs. */
	// +optional
	ThreatExceptions []string `json:"threatExceptions,omitempty"`
}

type CloudIDSEndpointStatus struct {
	/* Conditions represent the latest available observations of the
	   CloudIDSEndpoint's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Creation timestamp in RFC 3339 text format. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* URL of the endpoint's network address to which traffic is to be sent by Packet Mirroring. */
	// +optional
	EndpointForwardingRule *string `json:"endpointForwardingRule,omitempty"`

	/* Internal IP address of the endpoint's network entry point. */
	// +optional
	EndpointIp *string `json:"endpointIp,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* Last update timestamp in RFC 3339 text format. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudIDSEndpoint is the Schema for the cloudids API
// +k8s:openapi-gen=true
type CloudIDSEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudIDSEndpointSpec   `json:"spec,omitempty"`
	Status CloudIDSEndpointStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudIDSEndpointList contains a list of CloudIDSEndpoint
type CloudIDSEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudIDSEndpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudIDSEndpoint{}, &CloudIDSEndpointList{})
}
