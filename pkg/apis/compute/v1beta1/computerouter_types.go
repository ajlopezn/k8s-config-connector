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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RouterAdvertisedIpRanges struct {
	/* User-specified description for the IP range. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* The IP range to advertise. The value must be a
	CIDR-formatted string. */
	Range string `json:"range"`
}

type RouterBgp struct {
	/* User-specified flag to indicate which mode to use for advertisement. Default value: "DEFAULT" Possible values: ["DEFAULT", "CUSTOM"]. */
	// +optional
	AdvertiseMode *string `json:"advertiseMode,omitempty"`

	/* User-specified list of prefix groups to advertise in custom mode.
	This field can only be populated if advertiseMode is CUSTOM and
	is advertised to all peers of the router. These groups will be
	advertised in addition to any specified prefixes. Leave this field
	blank to advertise no custom groups.

	This enum field has the one valid value: ALL_SUBNETS. */
	// +optional
	AdvertisedGroups []string `json:"advertisedGroups,omitempty"`

	/* User-specified list of individual IP ranges to advertise in
	custom mode. This field can only be populated if advertiseMode
	is CUSTOM and is advertised to all peers of the router. These IP
	ranges will be advertised in addition to any specified groups.
	Leave this field blank to advertise no custom IP ranges. */
	// +optional
	AdvertisedIpRanges []RouterAdvertisedIpRanges `json:"advertisedIpRanges,omitempty"`

	/* Local BGP Autonomous System Number (ASN). Must be an RFC6996
	private ASN, either 16-bit or 32-bit. The value will be fixed for
	this router resource. All VPN tunnels that link to this router
	will have the same local ASN. */
	Asn int `json:"asn"`
}

type ComputeRouterSpec struct {
	/* BGP information specific to this router. */
	// +optional
	Bgp *RouterBgp `json:"bgp,omitempty"`

	/* An optional description of this resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. Field to indicate if a router is dedicated to use with encrypted
	Interconnect Attachment (IPsec-encrypted Cloud Interconnect feature).

	Not currently available publicly. */
	// +optional
	EncryptedInterconnectRouter *bool `json:"encryptedInterconnectRouter,omitempty"`

	/* A reference to the network to which this router belongs. */
	NetworkRef v1alpha1.ResourceRef `json:"networkRef"`

	/* Immutable. Region where the router resides. */
	Region string `json:"region"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type ComputeRouterStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeRouter's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Creation timestamp in RFC3339 text format. */
	CreationTimestamp string `json:"creationTimestamp,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/*  */
	SelfLink string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeRouter is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeRouter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeRouterSpec   `json:"spec,omitempty"`
	Status ComputeRouterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeRouterList contains a list of ComputeRouter
type ComputeRouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeRouter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeRouter{}, &ComputeRouterList{})
}
