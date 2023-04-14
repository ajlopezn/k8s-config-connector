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

type ServicesplittrafficSplit struct {
	/* Mapping from version IDs within the service to fractional (0.000, 1] allocations of traffic for that version. Each version can be specified only once, but some versions in the service may not have any traffic allocation. Services that have traffic allocated cannot be deleted until either the service is deleted or their traffic allocation is removed. Allocations must sum to 1. Up to two decimal place precision is supported for IP-based splits and up to three decimal places is supported for cookie-based splits. */
	Allocations map[string]string `json:"allocations"`

	/* Mechanism used to determine which version a request is sent to. The traffic selection algorithm will be stable for either type until allocations are changed. Possible values: ["UNSPECIFIED", "COOKIE", "IP", "RANDOM"]. */
	// +optional
	ShardBy *string `json:"shardBy,omitempty"`
}

type AppEngineServiceSplitTrafficSpec struct {
	/* If set to true traffic will be migrated to this version. */
	// +optional
	MigrateTraffic *bool `json:"migrateTraffic,omitempty"`

	/* Immutable. */
	// +optional
	Project *string `json:"project,omitempty"`

	/* Immutable. Optional. The service of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Mapping that defines fractional HTTP traffic diversion to different versions within the service. */
	Split ServicesplittrafficSplit `json:"split"`
}

type AppEngineServiceSplitTrafficStatus struct {
	/* Conditions represent the latest available observations of the
	   AppEngineServiceSplitTraffic's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AppEngineServiceSplitTraffic is the Schema for the appengine API
// +k8s:openapi-gen=true
type AppEngineServiceSplitTraffic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppEngineServiceSplitTrafficSpec   `json:"spec,omitempty"`
	Status AppEngineServiceSplitTrafficStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AppEngineServiceSplitTrafficList contains a list of AppEngineServiceSplitTraffic
type AppEngineServiceSplitTrafficList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppEngineServiceSplitTraffic `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppEngineServiceSplitTraffic{}, &AppEngineServiceSplitTrafficList{})
}
