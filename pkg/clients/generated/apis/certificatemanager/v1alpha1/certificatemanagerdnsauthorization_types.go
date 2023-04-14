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

type CertificateManagerDNSAuthorizationSpec struct {
	/* A human-readable description of the resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. A domain which is being authorized. A DnsAuthorization resource covers a
	single domain and its wildcard, e.g. authorization for "example.com" can
	be used to issue certificates for "example.com" and "*.example.com". */
	Domain string `json:"domain"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type DnsauthorizationDnsResourceRecordStatus struct {
	/* Data of the DNS Resource Record. */
	// +optional
	Data *string `json:"data,omitempty"`

	/* Fully qualified name of the DNS Resource Record.
	E.g. '_acme-challenge.example.com'. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* Type of the DNS Resource Record. */
	// +optional
	Type *string `json:"type,omitempty"`
}

type CertificateManagerDNSAuthorizationStatus struct {
	/* Conditions represent the latest available observations of the
	   CertificateManagerDNSAuthorization's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The structure describing the DNS Resource Record that needs to be added
	to DNS configuration for the authorization to be usable by
	certificate. */
	// +optional
	DnsResourceRecord []DnsauthorizationDnsResourceRecordStatus `json:"dnsResourceRecord,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CertificateManagerDNSAuthorization is the Schema for the certificatemanager API
// +k8s:openapi-gen=true
type CertificateManagerDNSAuthorization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CertificateManagerDNSAuthorizationSpec   `json:"spec,omitempty"`
	Status CertificateManagerDNSAuthorizationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CertificateManagerDNSAuthorizationList contains a list of CertificateManagerDNSAuthorization
type CertificateManagerDNSAuthorizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateManagerDNSAuthorization `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CertificateManagerDNSAuthorization{}, &CertificateManagerDNSAuthorizationList{})
}
