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

type EssentialContactsContactSpec struct {
	/* Immutable. The email address to send notifications to. This does not need to be a Google account. */
	Email string `json:"email"`

	/* The preferred language for notifications, as a ISO 639-1 language code. See Supported languages for a list of supported languages. */
	LanguageTag string `json:"languageTag"`

	/* The categories of notifications that the contact will receive communications for. */
	NotificationCategorySubscriptions []string `json:"notificationCategorySubscriptions"`

	/* Immutable. The resource to save this contact for. Format: organizations/{organization_id}, folders/{folder_id} or projects/{project_id}. */
	Parent string `json:"parent"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type EssentialContactsContactStatus struct {
	/* Conditions represent the latest available observations of the
	   EssentialContactsContact's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The identifier for the contact. Format: {resourceType}/{resource_id}/contacts/{contact_id}. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EssentialContactsContact is the Schema for the essentialcontacts API
// +k8s:openapi-gen=true
type EssentialContactsContact struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EssentialContactsContactSpec   `json:"spec,omitempty"`
	Status EssentialContactsContactStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EssentialContactsContactList contains a list of EssentialContactsContact
type EssentialContactsContactList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EssentialContactsContact `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EssentialContactsContact{}, &EssentialContactsContactList{})
}
