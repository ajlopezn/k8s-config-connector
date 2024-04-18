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

type SubscriptionAvroConfig struct {
	/* When true, write the subscription name, messageId, publishTime, attributes, and orderingKey as additional fields in the output. */
	// +optional
	WriteMetadata *bool `json:"writeMetadata,omitempty"`
}

type SubscriptionBigqueryConfig struct {
	/* When true and useTopicSchema is true, any fields that are a part of the topic schema that are not part of the BigQuery table schema are dropped when writing to BigQuery.
	Otherwise, the schemas must be kept in sync and any messages with extra fields are not written and remain in the subscription's backlog. */
	// +optional
	DropUnknownFields *bool `json:"dropUnknownFields,omitempty"`

	/* The name of the table to which to write data. */
	TableRef v1alpha1.ResourceRef `json:"tableRef"`

	/* When true, use the topic's schema as the columns to write to in BigQuery, if it exists. */
	// +optional
	UseTopicSchema *bool `json:"useTopicSchema,omitempty"`

	/* When true, write the subscription name, messageId, publishTime, attributes, and orderingKey to additional columns in the table.
	The subscription name, messageId, and publishTime fields are put in their own columns while all other message properties (other than data) are written to a JSON object in the attributes column. */
	// +optional
	WriteMetadata *bool `json:"writeMetadata,omitempty"`
}

type SubscriptionCloudStorageConfig struct {
	/* If set, message data will be written to Cloud Storage in Avro format. */
	// +optional
	AvroConfig *SubscriptionAvroConfig `json:"avroConfig,omitempty"`

	/* User-provided name for the Cloud Storage bucket. The bucket must be created by the user. The bucket name must be without any prefix like "gs://". */
	BucketRef v1alpha1.ResourceRef `json:"bucketRef"`

	/* User-provided prefix for Cloud Storage filename. */
	// +optional
	FilenamePrefix *string `json:"filenamePrefix,omitempty"`

	/* User-provided suffix for Cloud Storage filename. Must not end in "/". */
	// +optional
	FilenameSuffix *string `json:"filenameSuffix,omitempty"`

	/* The maximum bytes that can be written to a Cloud Storage file before a new file is created. Min 1 KB, max 10 GiB.
	The maxBytes limit may be exceeded in cases where messages are larger than the limit. */
	// +optional
	MaxBytes *int64 `json:"maxBytes,omitempty"`

	/* The maximum duration that can elapse before a new Cloud Storage file is created. Min 1 minute, max 10 minutes, default 5 minutes.
	May not exceed the subscription's acknowledgement deadline.
	A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s". */
	// +optional
	MaxDuration *string `json:"maxDuration,omitempty"`

	/* An output-only field that indicates whether or not the subscription can receive messages. */
	// +optional
	State *string `json:"state,omitempty"`
}

type SubscriptionDeadLetterPolicy struct {
	// +optional
	DeadLetterTopicRef *v1alpha1.ResourceRef `json:"deadLetterTopicRef,omitempty"`

	/* The maximum number of delivery attempts for any message. The value must be
	between 5 and 100.

	The number of delivery attempts is defined as 1 + (the sum of number of
	NACKs and number of times the acknowledgement deadline has been exceeded for the message).

	A NACK is any call to ModifyAckDeadline with a 0 deadline. Note that
	client libraries may automatically extend ack_deadlines.

	This field will be honored on a best effort basis.

	If this parameter is 0, a default value of 5 is used. */
	// +optional
	MaxDeliveryAttempts *int64 `json:"maxDeliveryAttempts,omitempty"`
}

type SubscriptionExpirationPolicy struct {
	/* Specifies the "time-to-live" duration for an associated resource. The
	resource expires if it is not active for a period of ttl.
	If ttl is set to "", the associated resource never expires.
	A duration in seconds with up to nine fractional digits, terminated by 's'.
	Example - "3.5s". */
	Ttl string `json:"ttl"`
}

type SubscriptionNoWrapper struct {
	/* When true, writes the Pub/Sub message metadata to
	'x-goog-pubsub-<KEY>:<VAL>' headers of the HTTP request. Writes the
	Pub/Sub message attributes to '<KEY>:<VAL>' headers of the HTTP request. */
	WriteMetadata bool `json:"writeMetadata"`
}

type SubscriptionOidcToken struct {
	/* Audience to be used when generating OIDC token. The audience claim
	identifies the recipients that the JWT is intended for. The audience
	value is a single case-sensitive string. Having multiple values (array)
	for the audience field is not supported. More info about the OIDC JWT
	token audience here: https://tools.ietf.org/html/rfc7519#section-4.1.3
	Note: if not specified, the Push endpoint URL will be used. */
	// +optional
	Audience *string `json:"audience,omitempty"`

	/* Service account email to be used for generating the OIDC token.
	The caller (for subscriptions.create, subscriptions.patch, and
	subscriptions.modifyPushConfig RPCs) must have the
	iam.serviceAccounts.actAs permission for the service account. */
	ServiceAccountEmail string `json:"serviceAccountEmail"`
}

type SubscriptionPushConfig struct {
	/* Endpoint configuration attributes.

	Every endpoint has a set of API supported attributes that can
	be used to control different aspects of the message delivery.

	The currently supported attribute is x-goog-version, which you
	can use to change the format of the pushed message. This
	attribute indicates the version of the data expected by
	the endpoint. This controls the shape of the pushed message
	(i.e., its fields and metadata). The endpoint version is
	based on the version of the Pub/Sub API.

	If not present during the subscriptions.create call,
	it will default to the version of the API used to make
	such call. If not present during a subscriptions.modifyPushConfig
	call, its value will not be changed. subscriptions.get
	calls will always return a valid version, even if the
	subscription was created without this attribute.

	The possible values for this attribute are:

	- v1beta1: uses the push format defined in the v1beta1 Pub/Sub API.
	- v1 or v1beta2: uses the push format defined in the v1 Pub/Sub API. */
	// +optional
	Attributes map[string]string `json:"attributes,omitempty"`

	/* When set, the payload to the push endpoint is not wrapped.Sets the
	'data' field as the HTTP body for delivery. */
	// +optional
	NoWrapper *SubscriptionNoWrapper `json:"noWrapper,omitempty"`

	/* If specified, Pub/Sub will generate and attach an OIDC JWT token as
	an Authorization header in the HTTP request for every pushed message. */
	// +optional
	OidcToken *SubscriptionOidcToken `json:"oidcToken,omitempty"`

	/* A URL locating the endpoint to which messages should be pushed.
	For example, a Webhook endpoint might use
	"https://example.com/push". */
	PushEndpoint string `json:"pushEndpoint"`
}

type SubscriptionRetryPolicy struct {
	/* The maximum delay between consecutive deliveries of a given message. Value should be between 0 and 600 seconds. Defaults to 600 seconds.
	A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". */
	// +optional
	MaximumBackoff *string `json:"maximumBackoff,omitempty"`

	/* The minimum delay between consecutive deliveries of a given message. Value should be between 0 and 600 seconds. Defaults to 10 seconds.
	A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". */
	// +optional
	MinimumBackoff *string `json:"minimumBackoff,omitempty"`
}

type PubSubSubscriptionSpec struct {
	/* This value is the maximum time after a subscriber receives a message
	before the subscriber should acknowledge the message. After message
	delivery but before the ack deadline expires and before the message is
	acknowledged, it is an outstanding message and will not be delivered
	again during that time (on a best-effort basis).

	For pull subscriptions, this value is used as the initial value for
	the ack deadline. To override this value for a given message, call
	subscriptions.modifyAckDeadline with the corresponding ackId if using
	pull. The minimum custom deadline you can specify is 10 seconds. The
	maximum custom deadline you can specify is 600 seconds (10 minutes).
	If this parameter is 0, a default value of 10 seconds is used.

	For push delivery, this value is also used to set the request timeout
	for the call to the push endpoint.

	If the subscriber never acknowledges the message, the Pub/Sub system
	will eventually redeliver the message. */
	// +optional
	AckDeadlineSeconds *int64 `json:"ackDeadlineSeconds,omitempty"`

	/* If delivery to BigQuery is used with this subscription, this field is used to configure it.
	Either pushConfig, bigQueryConfig or cloudStorageConfig can be set, but not combined.
	If all three are empty, then the subscriber will pull and ack messages using API methods. */
	// +optional
	BigqueryConfig *SubscriptionBigqueryConfig `json:"bigqueryConfig,omitempty"`

	/* If delivery to Cloud Storage is used with this subscription, this field is used to configure it.
	Either pushConfig, bigQueryConfig or cloudStorageConfig can be set, but not combined.
	If all three are empty, then the subscriber will pull and ack messages using API methods. */
	// +optional
	CloudStorageConfig *SubscriptionCloudStorageConfig `json:"cloudStorageConfig,omitempty"`

	/* A policy that specifies the conditions for dead lettering messages in
	this subscription. If dead_letter_policy is not set, dead lettering
	is disabled.

	The Cloud Pub/Sub service account associated with this subscription's
	parent project (i.e.,
	service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
	permission to Acknowledge() messages on this subscription. */
	// +optional
	DeadLetterPolicy *SubscriptionDeadLetterPolicy `json:"deadLetterPolicy,omitempty"`

	/* If 'true', Pub/Sub provides the following guarantees for the delivery
	of a message with a given value of messageId on this Subscriptions':

	- The message sent to a subscriber is guaranteed not to be resent before the message's acknowledgement deadline expires.

	- An acknowledged message will not be resent to a subscriber.

	Note that subscribers may still receive multiple copies of a message when 'enable_exactly_once_delivery'
	is true if the message was published multiple times by a publisher client. These copies are considered distinct by Pub/Sub and have distinct messageId values. */
	// +optional
	EnableExactlyOnceDelivery *bool `json:"enableExactlyOnceDelivery,omitempty"`

	/* Immutable. If 'true', messages published with the same orderingKey in PubsubMessage will be delivered to
	the subscribers in the order in which they are received by the Pub/Sub system. Otherwise, they
	may be delivered in any order. */
	// +optional
	EnableMessageOrdering *bool `json:"enableMessageOrdering,omitempty"`

	/* A policy that specifies the conditions for this subscription's expiration.
	A subscription is considered active as long as any connected subscriber
	is successfully consuming messages from the subscription or is issuing
	operations on the subscription. If expirationPolicy is not set, a default
	policy with ttl of 31 days will be used.  If it is set but ttl is "", the
	resource never expires.  The minimum allowed value for expirationPolicy.ttl
	is 1 day. */
	// +optional
	ExpirationPolicy *SubscriptionExpirationPolicy `json:"expirationPolicy,omitempty"`

	/* Immutable. The subscription only delivers the messages that match the filter.
	Pub/Sub automatically acknowledges the messages that don't match the filter. You can filter messages
	by their attributes. The maximum length of a filter is 256 bytes. After creating the subscription,
	you can't modify the filter. */
	// +optional
	Filter *string `json:"filter,omitempty"`

	/* How long to retain unacknowledged messages in the subscription's
	backlog, from the moment a message is published. If
	retain_acked_messages is true, then this also configures the retention
	of acknowledged messages, and thus configures how far back in time a
	subscriptions.seek can be done. Defaults to 7 days. Cannot be more
	than 7 days ('"604800s"') or less than 10 minutes ('"600s"').

	A duration in seconds with up to nine fractional digits, terminated
	by 's'. Example: '"600.5s"'. */
	// +optional
	MessageRetentionDuration *string `json:"messageRetentionDuration,omitempty"`

	/* If push delivery is used with this subscription, this field is used to
	configure it. An empty pushConfig signifies that the subscriber will
	pull and ack messages using API methods. */
	// +optional
	PushConfig *SubscriptionPushConfig `json:"pushConfig,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Indicates whether to retain acknowledged messages. If 'true', then
	messages are not expunged from the subscription's backlog, even if
	they are acknowledged, until they fall out of the
	messageRetentionDuration window. */
	// +optional
	RetainAckedMessages *bool `json:"retainAckedMessages,omitempty"`

	/* A policy that specifies how Pub/Sub retries message delivery for this subscription.

	If not set, the default retry policy is applied. This generally implies that messages will be retried as soon as possible for healthy subscribers.
	RetryPolicy will be triggered on NACKs or acknowledgement deadline exceeded events for a given message. */
	// +optional
	RetryPolicy *SubscriptionRetryPolicy `json:"retryPolicy,omitempty"`

	/* Reference to a PubSubTopic. */
	TopicRef v1alpha1.ResourceRef `json:"topicRef"`
}

type PubSubSubscriptionStatus struct {
	/* Conditions represent the latest available observations of the
	   PubSubSubscription's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcppubsubsubscription;gcppubsubsubscriptions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// PubSubSubscription is the Schema for the pubsub API
// +k8s:openapi-gen=true
type PubSubSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PubSubSubscriptionSpec   `json:"spec,omitempty"`
	Status PubSubSubscriptionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PubSubSubscriptionList contains a list of PubSubSubscription
type PubSubSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PubSubSubscription `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PubSubSubscription{}, &PubSubSubscriptionList{})
}
