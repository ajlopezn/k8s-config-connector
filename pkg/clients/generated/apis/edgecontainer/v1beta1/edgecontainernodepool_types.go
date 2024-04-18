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

type NodepoolLocalDiskEncryption struct {
	/* The Cloud KMS CryptoKeyVersion currently in use for protecting node local disks. Only applicable if kmsKey is set. */
	// +optional
	KmsKeyActiveVersion *string `json:"kmsKeyActiveVersion,omitempty"`

	// +optional
	KmsKeyRef *v1alpha1.ResourceRef `json:"kmsKeyRef,omitempty"`

	/* Availability of the Cloud KMS CryptoKey. If not KEY_AVAILABLE, then nodes may go offline as they cannot access their local data.
	This can be caused by a lack of permissions to use the key, or if the key is disabled or deleted. */
	// +optional
	KmsKeyState *string `json:"kmsKeyState,omitempty"`
}

type NodepoolNodeConfig struct {
	/* "The Kubernetes node labels". */
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
}

type EdgeContainerNodePoolSpec struct {
	ClusterRef v1alpha1.ResourceRef `json:"clusterRef"`

	/* Local disk encryption options. This field is only used when enabling CMEK support. */
	// +optional
	LocalDiskEncryption *NodepoolLocalDiskEncryption `json:"localDiskEncryption,omitempty"`

	/* Immutable. The location of the resource. */
	Location string `json:"location"`

	/* Only machines matching this filter will be allowed to join the node pool.
	The filtering language accepts strings like "name=<name>", and is
	documented in more detail in [AIP-160](https://google.aip.dev/160). */
	// +optional
	MachineFilter *string `json:"machineFilter,omitempty"`

	/* Configuration for each node in the NodePool. */
	// +optional
	NodeConfig *NodepoolNodeConfig `json:"nodeConfig,omitempty"`

	/* The number of nodes in the pool. */
	NodeCount int64 `json:"nodeCount"`

	/* Immutable. Name of the Google Distributed Cloud Edge zone where this node pool will be created. For example: 'us-central1-edge-customer-a'. */
	NodeLocation string `json:"nodeLocation"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type EdgeContainerNodePoolStatus struct {
	/* Conditions represent the latest available observations of the
	   EdgeContainerNodePool's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The time when the node pool was created. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* The lowest release version among all worker nodes. */
	// +optional
	NodeVersion *string `json:"nodeVersion,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* The time when the node pool was last updated. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpedgecontainernodepool;gcpedgecontainernodepools
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// EdgeContainerNodePool is the Schema for the edgecontainer API
// +k8s:openapi-gen=true
type EdgeContainerNodePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EdgeContainerNodePoolSpec   `json:"spec,omitempty"`
	Status EdgeContainerNodePoolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EdgeContainerNodePoolList contains a list of EdgeContainerNodePool
type EdgeContainerNodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EdgeContainerNodePool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EdgeContainerNodePool{}, &EdgeContainerNodePoolList{})
}
