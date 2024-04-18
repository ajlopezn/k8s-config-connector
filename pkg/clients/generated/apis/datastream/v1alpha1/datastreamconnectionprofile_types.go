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

type ConnectionprofileBigqueryProfile struct {
}

type ConnectionprofileCaCertificate struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *ConnectionprofileValueFrom `json:"valueFrom,omitempty"`
}

type ConnectionprofileClientCertificate struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *ConnectionprofileValueFrom `json:"valueFrom,omitempty"`
}

type ConnectionprofileClientKey struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *ConnectionprofileValueFrom `json:"valueFrom,omitempty"`
}

type ConnectionprofileForwardSshConnectivity struct {
	/* Hostname for the SSH tunnel. */
	Hostname string `json:"hostname"`

	/* Immutable. SSH password. */
	// +optional
	Password *ConnectionprofilePassword `json:"password,omitempty"`

	/* Port for the SSH tunnel. */
	// +optional
	Port *int64 `json:"port,omitempty"`

	/* Immutable. SSH private key. */
	// +optional
	PrivateKey *ConnectionprofilePrivateKey `json:"privateKey,omitempty"`

	/* Username for the SSH tunnel. */
	Username string `json:"username"`
}

type ConnectionprofileGcsProfile struct {
	/* The Cloud Storage bucket name. */
	Bucket string `json:"bucket"`

	/* The root path inside the Cloud Storage bucket. */
	// +optional
	RootPath *string `json:"rootPath,omitempty"`
}

type ConnectionprofileMysqlProfile struct {
	/* Hostname for the MySQL connection. */
	Hostname string `json:"hostname"`

	/* Password for the MySQL connection. */
	Password ConnectionprofilePassword `json:"password"`

	/* Port for the MySQL connection. */
	// +optional
	Port *int64 `json:"port,omitempty"`

	/* SSL configuration for the MySQL connection. */
	// +optional
	SslConfig *ConnectionprofileSslConfig `json:"sslConfig,omitempty"`

	/* Username for the MySQL connection. */
	Username string `json:"username"`
}

type ConnectionprofileOracleProfile struct {
	/* Connection string attributes. */
	// +optional
	ConnectionAttributes map[string]string `json:"connectionAttributes,omitempty"`

	/* Database for the Oracle connection. */
	DatabaseService string `json:"databaseService"`

	/* Hostname for the Oracle connection. */
	Hostname string `json:"hostname"`

	/* Password for the Oracle connection. */
	Password ConnectionprofilePassword `json:"password"`

	/* Port for the Oracle connection. */
	// +optional
	Port *int64 `json:"port,omitempty"`

	/* Username for the Oracle connection. */
	Username string `json:"username"`
}

type ConnectionprofilePassword struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *ConnectionprofileValueFrom `json:"valueFrom,omitempty"`
}

type ConnectionprofilePostgresqlProfile struct {
	/* Database for the PostgreSQL connection. */
	Database string `json:"database"`

	/* Hostname for the PostgreSQL connection. */
	Hostname string `json:"hostname"`

	/* Password for the PostgreSQL connection. */
	Password ConnectionprofilePassword `json:"password"`

	/* Port for the PostgreSQL connection. */
	// +optional
	Port *int64 `json:"port,omitempty"`

	/* Username for the PostgreSQL connection. */
	Username string `json:"username"`
}

type ConnectionprofilePrivateConnectivity struct {
	/* A reference to a private connection resource. Format: 'projects/{project}/locations/{location}/privateConnections/{name}'. */
	PrivateConnection string `json:"privateConnection"`
}

type ConnectionprofilePrivateKey struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *ConnectionprofileValueFrom `json:"valueFrom,omitempty"`
}

type ConnectionprofileSslConfig struct {
	/* Immutable. PEM-encoded certificate of the CA that signed the source database
	server's certificate. */
	// +optional
	CaCertificate *ConnectionprofileCaCertificate `json:"caCertificate,omitempty"`

	/* Indicates whether the clientKey field is set. */
	// +optional
	CaCertificateSet *bool `json:"caCertificateSet,omitempty"`

	/* Immutable. PEM-encoded certificate that will be used by the replica to
	authenticate against the source database server. If this field
	is used then the 'clientKey' and the 'caCertificate' fields are
	mandatory. */
	// +optional
	ClientCertificate *ConnectionprofileClientCertificate `json:"clientCertificate,omitempty"`

	/* Indicates whether the clientCertificate field is set. */
	// +optional
	ClientCertificateSet *bool `json:"clientCertificateSet,omitempty"`

	/* Immutable. PEM-encoded private key associated with the Client Certificate.
	If this field is used then the 'client_certificate' and the
	'ca_certificate' fields are mandatory. */
	// +optional
	ClientKey *ConnectionprofileClientKey `json:"clientKey,omitempty"`

	/* Indicates whether the clientKey field is set. */
	// +optional
	ClientKeySet *bool `json:"clientKeySet,omitempty"`
}

type ConnectionprofileValueFrom struct {
	/* Reference to a value with the given key in the given Secret in the resource's namespace. */
	// +optional
	SecretKeyRef *v1alpha1.SecretKeyRef `json:"secretKeyRef,omitempty"`
}

type DatastreamConnectionProfileSpec struct {
	/* BigQuery warehouse profile. */
	// +optional
	BigqueryProfile *ConnectionprofileBigqueryProfile `json:"bigqueryProfile,omitempty"`

	/* Display name. */
	DisplayName string `json:"displayName"`

	/* Forward SSH tunnel connectivity. */
	// +optional
	ForwardSshConnectivity *ConnectionprofileForwardSshConnectivity `json:"forwardSshConnectivity,omitempty"`

	/* Cloud Storage bucket profile. */
	// +optional
	GcsProfile *ConnectionprofileGcsProfile `json:"gcsProfile,omitempty"`

	/* Immutable. The name of the location this connection profile is located in. */
	Location string `json:"location"`

	/* MySQL database profile. */
	// +optional
	MysqlProfile *ConnectionprofileMysqlProfile `json:"mysqlProfile,omitempty"`

	/* Oracle database profile. */
	// +optional
	OracleProfile *ConnectionprofileOracleProfile `json:"oracleProfile,omitempty"`

	/* PostgreSQL database profile. */
	// +optional
	PostgresqlProfile *ConnectionprofilePostgresqlProfile `json:"postgresqlProfile,omitempty"`

	/* Private connectivity. */
	// +optional
	PrivateConnectivity *ConnectionprofilePrivateConnectivity `json:"privateConnectivity,omitempty"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The connectionProfileId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type DatastreamConnectionProfileStatus struct {
	/* Conditions represent the latest available observations of the
	   DatastreamConnectionProfile's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The resource's name. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatastreamconnectionprofile;gcpdatastreamconnectionprofiles
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=alpha";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DatastreamConnectionProfile is the Schema for the datastream API
// +k8s:openapi-gen=true
type DatastreamConnectionProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatastreamConnectionProfileSpec   `json:"spec,omitempty"`
	Status DatastreamConnectionProfileStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DatastreamConnectionProfileList contains a list of DatastreamConnectionProfile
type DatastreamConnectionProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatastreamConnectionProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatastreamConnectionProfile{}, &DatastreamConnectionProfileList{})
}
