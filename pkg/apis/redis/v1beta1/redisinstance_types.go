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

type RedisInstanceSpec struct {
	/* Immutable. Only applicable to STANDARD_HA tier which protects the instance
	against zonal failures by provisioning it across two zones.
	If provided, it must be a different zone from the one provided in
	[locationId]. */
	// +optional
	AlternativeLocationId *string `json:"alternativeLocationId,omitempty"`

	/* Optional. Indicates whether OSS Redis AUTH is enabled for the
	instance. If set to "true" AUTH is enabled on the instance.
	Default value is "false" meaning AUTH is disabled. */
	// +optional
	AuthEnabled *bool `json:"authEnabled,omitempty"`

	/* AUTH String set on the instance. This field will only be populated if auth_enabled is true. */
	// +optional
	AuthString *string `json:"authString,omitempty"`

	/* The network to which the instance is connected. If left
	unspecified, the default network will be used. */
	// +optional
	AuthorizedNetworkRef *v1alpha1.ResourceRef `json:"authorizedNetworkRef,omitempty"`

	/* Immutable. The connection mode of the Redis instance. Default value: "DIRECT_PEERING" Possible values: ["DIRECT_PEERING", "PRIVATE_SERVICE_ACCESS"]. */
	// +optional
	ConnectMode *string `json:"connectMode,omitempty"`

	/* An arbitrary and optional user-provided name for the instance. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* Immutable. The zone where the instance will be provisioned. If not provided,
	the service will choose a zone for the instance. For STANDARD_HA tier,
	instances will be created across two zones for protection against
	zonal failures. If [alternativeLocationId] is also provided, it must
	be different from [locationId]. */
	// +optional
	LocationId *string `json:"locationId,omitempty"`

	/* Redis memory size in GiB. */
	MemorySizeGb int `json:"memorySizeGb"`

	/* Redis configuration parameters, according to http://redis.io/topics/config.
	Please check Memorystore documentation for the list of supported parameters:
	https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs. */
	// +optional
	RedisConfigs map[string]string `json:"redisConfigs,omitempty"`

	/* The version of Redis software. If not provided, latest supported
	version will be used. Please check the API documentation linked
	at the top for the latest valid values. */
	// +optional
	RedisVersion *string `json:"redisVersion,omitempty"`

	/* Immutable. The name of the Redis region of the instance. */
	Region string `json:"region"`

	/* Immutable. The CIDR range of internal addresses that are reserved for this
	instance. If not provided, the service will choose an unused /29
	block, for example, 10.0.0.0/29 or 192.168.0.0/29. Ranges must be
	unique and non-overlapping with existing subnets in an authorized
	network. */
	// +optional
	ReservedIpRange *string `json:"reservedIpRange,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The service tier of the instance. Must be one of these values:

	- BASIC: standalone instance
	- STANDARD_HA: highly available primary/replica instances Default value: "BASIC" Possible values: ["BASIC", "STANDARD_HA"]. */
	// +optional
	Tier *string `json:"tier,omitempty"`

	/* Immutable. The TLS mode of the Redis instance, If not provided, TLS is disabled for the instance.

	- SERVER_AUTHENTICATION: Client to Server traffic encryption enabled with server authentcation Default value: "DISABLED" Possible values: ["SERVER_AUTHENTICATION", "DISABLED"]. */
	// +optional
	TransitEncryptionMode *string `json:"transitEncryptionMode,omitempty"`
}

type InstanceServerCaCertsStatus struct {
	/* Serial number, as extracted from the certificate. */
	Cert string `json:"cert,omitempty"`

	/* The time when the certificate was created. */
	CreateTime string `json:"createTime,omitempty"`

	/* The time when the certificate expires. */
	ExpireTime string `json:"expireTime,omitempty"`

	/* Serial number, as extracted from the certificate. */
	SerialNumber string `json:"serialNumber,omitempty"`

	/* Sha1 Fingerprint of the certificate. */
	Sha1Fingerprint string `json:"sha1Fingerprint,omitempty"`
}

type RedisInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   RedisInstance's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The time the instance was created in RFC3339 UTC "Zulu" format,
	accurate to nanoseconds. */
	CreateTime string `json:"createTime,omitempty"`
	/* The current zone where the Redis endpoint is placed.
	For Basic Tier instances, this will always be the same as the
	[locationId] provided by the user at creation time. For Standard Tier
	instances, this can be either [locationId] or [alternativeLocationId]
	and can change after a failover event. */
	CurrentLocationId string `json:"currentLocationId,omitempty"`
	/* Hostname or IP address of the exposed Redis endpoint used by clients
	to connect to the service. */
	Host string `json:"host,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/* Output only. Cloud IAM identity used by import / export operations
	to transfer data to/from Cloud Storage. Format is "serviceAccount:".
	The value may change over time for a given instance so should be
	checked before each import/export operation. */
	PersistenceIamIdentity string `json:"persistenceIamIdentity,omitempty"`
	/* The port number of the exposed Redis endpoint. */
	Port int `json:"port,omitempty"`
	/* List of server CA certificates for the instance. */
	ServerCaCerts []InstanceServerCaCertsStatus `json:"serverCaCerts,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedisInstance is the Schema for the redis API
// +k8s:openapi-gen=true
type RedisInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedisInstanceSpec   `json:"spec,omitempty"`
	Status RedisInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedisInstanceList contains a list of RedisInstance
type RedisInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RedisInstance{}, &RedisInstanceList{})
}
