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

type FunctionEventTrigger struct {
	/* Immutable. Required. The type of event to observe. For example:
	`providers/cloud.storage/eventTypes/object.change` and
	`providers/cloud.pubsub/eventTypes/topic.publish`.

	Event types match pattern `providers/* /eventTypes/*.*`.
	The pattern contains:

	1. namespace: For example, `cloud.storage` and
	`google.firebase.analytics`.
	2. resource type: The type of resource on which event occurs. For
	example, the Google Cloud Storage API includes the type `object`.
	3. action: The action that generates the event. For example, action for
	a Google Cloud Storage Object is 'change'.
	These parts are lower case. */
	EventType string `json:"eventType"`

	/* Immutable. Specifies policy for failed executions. */
	// +optional
	FailurePolicy *bool `json:"failurePolicy,omitempty"`

	/* Immutable. */
	ResourceRef v1alpha1.ResourceRef `json:"resourceRef"`

	/* Immutable. The hostname of the service that should be observed.

	If no string is provided, the default service implementing the API will
	be used. For example, `storage.googleapis.com` is the default for all
	event types in the `google.storage` namespace. */
	// +optional
	Service *string `json:"service,omitempty"`
}

type FunctionHttpsTrigger struct {
	/* Immutable. Both HTTP and HTTPS requests with URLs that match the handler succeed without redirects. The application can examine the request to determine which protocol was used and respond accordingly. Possible values: SECURITY_LEVEL_UNSPECIFIED, SECURE_ALWAYS, SECURE_OPTIONAL */
	// +optional
	SecurityLevel *string `json:"securityLevel,omitempty"`
}

type FunctionSourceRepository struct {
	/* Immutable. The URL pointing to the hosted repository where the function is defined.
	There are supported Cloud Source Repository URLs in the following
	formats:

	To refer to a specific commit:
	`https://source.developers.google.com/projects/* /repos/* /revisions/* /paths/*`
	To refer to a moveable alias (branch):
	`https://source.developers.google.com/projects/* /repos/* /moveable-aliases/* /paths/*`
	In particular, to refer to HEAD use `master` moveable alias.
	To refer to a specific fixed alias (tag):
	`https://source.developers.google.com/projects/* /repos/* /fixed-aliases/* /paths/*`

	You may omit `paths/*` if you want to use the main directory. */
	Url string `json:"url"`
}

type CloudFunctionsFunctionSpec struct {
	/* Memory (in MB), available to the function. Default value is 256MB. Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB. */
	// +optional
	AvailableMemoryMb *int64 `json:"availableMemoryMb,omitempty"`

	/* User-provided description of a function. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. The name of the function (as defined in source code) that will be
	executed. Defaults to the resource name suffix, if not specified. For
	backward compatibility, if function with given name is not found, then the
	system will try to use function named "function".
	For Node.js this is name of a function exported by the module specified
	in `source_location`. */
	// +optional
	EntryPoint *string `json:"entryPoint,omitempty"`

	/* Environment variables that shall be available during function execution. */
	// +optional
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty"`

	/* Immutable. A source that fires events in response to a condition in another service. */
	// +optional
	EventTrigger *FunctionEventTrigger `json:"eventTrigger,omitempty"`

	/* Immutable. An HTTPS endpoint type of source that can be triggered via URL. */
	// +optional
	HttpsTrigger *FunctionHttpsTrigger `json:"httpsTrigger,omitempty"`

	/* The ingress settings for the function, controlling what traffic can reach
	it. Possible values: INGRESS_SETTINGS_UNSPECIFIED, ALLOW_ALL, ALLOW_INTERNAL_ONLY, ALLOW_INTERNAL_AND_GCLB */
	// +optional
	IngressSettings *string `json:"ingressSettings,omitempty"`

	/* The limit on the maximum number of function instances that may coexist at a
	given time. */
	// +optional
	MaxInstances *int64 `json:"maxInstances,omitempty"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. The name of the Cloud Functions region of the function. */
	Region string `json:"region"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The runtime in which to run the function. Required when deploying a new
	function, optional when updating an existing function. For a complete
	list of possible choices, see the
	[`gcloud` command
	reference](/sdk/gcloud/reference/functions/deploy#--runtime). */
	Runtime string `json:"runtime"`

	/* Immutable. */
	// +optional
	ServiceAccountRef *v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`

	/* Immutable. The Google Cloud Storage URL, starting with gs://, pointing to the zip archive which contains the function. */
	// +optional
	SourceArchiveUrl *string `json:"sourceArchiveUrl,omitempty"`

	/* Immutable. Represents parameters related to source repository where a function is hosted. */
	// +optional
	SourceRepository *FunctionSourceRepository `json:"sourceRepository,omitempty"`

	/* The function execution timeout. Execution is considered failed and
	can be terminated if the function is not completed at the end of the
	timeout period. Defaults to 60 seconds. */
	// +optional
	Timeout *string `json:"timeout,omitempty"`

	/* The egress settings for the connector, controlling what traffic is diverted
	through it. Possible values: VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED, PRIVATE_RANGES_ONLY, ALL_TRAFFIC */
	// +optional
	VpcConnectorEgressSettings *string `json:"vpcConnectorEgressSettings,omitempty"`

	// +optional
	VpcConnectorRef *v1alpha1.ResourceRef `json:"vpcConnectorRef,omitempty"`
}

type FunctionHttpsTriggerStatus struct {
	/* Output only. The deployed url for the function. */
	// +optional
	Url *string `json:"url,omitempty"`
}

type FunctionSourceRepositoryStatus struct {
	/* Output only. The URL pointing to the hosted repository where the function
	were defined at the time of deployment. It always points to a specific
	commit in the format described above. */
	// +optional
	DeployedUrl *string `json:"deployedUrl,omitempty"`
}

type CloudFunctionsFunctionStatus struct {
	/* Conditions represent the latest available observations of the
	   CloudFunctionsFunction's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	// +optional
	HttpsTrigger *FunctionHttpsTriggerStatus `json:"httpsTrigger,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// +optional
	SourceRepository *FunctionSourceRepositoryStatus `json:"sourceRepository,omitempty"`

	/* Output only. Status of the function deployment. Possible values: CLOUD_FUNCTION_STATUS_UNSPECIFIED, ACTIVE, OFFLINE, DEPLOY_IN_PROGRESS, DELETE_IN_PROGRESS, UNKNOWN */
	// +optional
	Status *string `json:"status,omitempty"`

	/* Output only. The last update timestamp of a Cloud Function in RFC3339 UTC 'Zulu' format, with nanosecond resolution and up to nine fractional digits. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`

	/* Output only. The version identifier of the Cloud Function. Each deployment attempt
	results in a new version of a function being created. */
	// +optional
	VersionId *int64 `json:"versionId,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudfunctionsfunction;gcpcloudfunctionsfunctions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudFunctionsFunction is the Schema for the cloudfunctions API
// +k8s:openapi-gen=true
type CloudFunctionsFunction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudFunctionsFunctionSpec   `json:"spec,omitempty"`
	Status CloudFunctionsFunctionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudFunctionsFunctionList contains a list of CloudFunctionsFunction
type CloudFunctionsFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudFunctionsFunction `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudFunctionsFunction{}, &CloudFunctionsFunctionList{})
}
