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

type BigQueryReservationReservationSpec struct {
	/* Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency will be automatically set based on the reservation size. */
	// +optional
	Concurrency *int `json:"concurrency,omitempty"`

	/* If false, any query using this reservation will use idle slots from other reservations within
	the same admin project. If true, a query using this reservation will execute with the slot
	capacity specified above at most. */
	// +optional
	IgnoreIdleSlots *bool `json:"ignoreIdleSlots,omitempty"`

	/* Immutable. The geographic location where the transfer config should reside.
	Examples: US, EU, asia-northeast1. The default value is US. */
	Location string `json:"location"`

	/* Applicable only for reservations located within one of the BigQuery multi-regions (US or EU).
	If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region. */
	// +optional
	MultiRegionAuxiliary *bool `json:"multiRegionAuxiliary,omitempty"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
	unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false. */
	SlotCapacity int `json:"slotCapacity"`
}

type BigQueryReservationReservationStatus struct {
	/* Conditions represent the latest available observations of the
	   BigQueryReservationReservation's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BigQueryReservationReservation is the Schema for the bigqueryreservation API
// +k8s:openapi-gen=true
type BigQueryReservationReservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryReservationReservationSpec   `json:"spec,omitempty"`
	Status BigQueryReservationReservationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BigQueryReservationReservationList contains a list of BigQueryReservationReservation
type BigQueryReservationReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryReservationReservation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigQueryReservationReservation{}, &BigQueryReservationReservationList{})
}
