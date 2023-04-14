// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	"net/http"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type ComputeV1alpha1Interface interface {
	RESTClient() rest.Interface
	ComputeAutoscalersGetter
	ComputeBackendBucketSignedURLKeysGetter
	ComputeBackendServiceSignedURLKeysGetter
	ComputeDiskResourcePolicyAttachmentsGetter
	ComputeGlobalNetworkEndpointsGetter
	ComputeGlobalNetworkEndpointGroupsGetter
	ComputeInstanceGroupNamedPortsGetter
	ComputeMachineImagesGetter
	ComputeManagedSSLCertificatesGetter
	ComputeNetworkEndpointsGetter
	ComputeNetworkPeeringRoutesConfigsGetter
	ComputeOrganizationSecurityPoliciesGetter
	ComputeOrganizationSecurityPolicyAssociationsGetter
	ComputeOrganizationSecurityPolicyRulesGetter
	ComputePerInstanceConfigsGetter
	ComputeRegionAutoscalersGetter
	ComputeRegionDiskResourcePolicyAttachmentsGetter
	ComputeRegionPerInstanceConfigsGetter
}

// ComputeV1alpha1Client is used to interact with features provided by the compute.cnrm.cloud.google.com group.
type ComputeV1alpha1Client struct {
	restClient rest.Interface
}

func (c *ComputeV1alpha1Client) ComputeAutoscalers(namespace string) ComputeAutoscalerInterface {
	return newComputeAutoscalers(c, namespace)
}

func (c *ComputeV1alpha1Client) ComputeBackendBucketSignedURLKeys(namespace string) ComputeBackendBucketSignedURLKeyInterface {
	return newComputeBackendBucketSignedURLKeys(c, namespace)
}

func (c *ComputeV1alpha1Client) ComputeBackendServiceSignedURLKeys(namespace string) ComputeBackendServiceSignedURLKeyInterface {
	return newComputeBackendServiceSignedURLKeys(c, namespace)
}

func (c *ComputeV1alpha1Client) ComputeDiskResourcePolicyAttachments(namespace string) ComputeDiskResourcePolicyAttachmentInterface {
	return newComputeDiskResourcePolicyAttachments(c, namespace)
}

func (c *ComputeV1alpha1Client) ComputeGlobalNetworkEndpoints(namespace string) ComputeGlobalNetworkEndpointInterface {
	return newComputeGlobalNetworkEndpoints(c, namespace)
}

func (c *ComputeV1alpha1Client) ComputeGlobalNetworkEndpointGroups(namespace string) ComputeGlobalNetworkEndpointGroupInterface {
	return newComputeGlobalNetworkEndpointGroups(c, namespace)
}

func (c *ComputeV1alpha1Client) ComputeInstanceGroupNamedPorts(namespace string) ComputeInstanceGroupNamedPortInterface {
	return newComputeInstanceGroupNamedPorts(c, namespace)
}

func (c *ComputeV1alpha1Client) ComputeMachineImages(namespace string) ComputeMachineImageInterface {
	return newComputeMachineImages(c, namespace)
}

func (c *ComputeV1alpha1Client) ComputeManagedSSLCertificates(namespace string) ComputeManagedSSLCertificateInterface {
	return newComputeManagedSSLCertificates(c, namespace)
}

func (c *ComputeV1alpha1Client) ComputeNetworkEndpoints(namespace string) ComputeNetworkEndpointInterface {
	return newComputeNetworkEndpoints(c, namespace)
}

func (c *ComputeV1alpha1Client) ComputeNetworkPeeringRoutesConfigs(namespace string) ComputeNetworkPeeringRoutesConfigInterface {
	return newComputeNetworkPeeringRoutesConfigs(c, namespace)
}

func (c *ComputeV1alpha1Client) ComputeOrganizationSecurityPolicies(namespace string) ComputeOrganizationSecurityPolicyInterface {
	return newComputeOrganizationSecurityPolicies(c, namespace)
}

func (c *ComputeV1alpha1Client) ComputeOrganizationSecurityPolicyAssociations(namespace string) ComputeOrganizationSecurityPolicyAssociationInterface {
	return newComputeOrganizationSecurityPolicyAssociations(c, namespace)
}

func (c *ComputeV1alpha1Client) ComputeOrganizationSecurityPolicyRules(namespace string) ComputeOrganizationSecurityPolicyRuleInterface {
	return newComputeOrganizationSecurityPolicyRules(c, namespace)
}

func (c *ComputeV1alpha1Client) ComputePerInstanceConfigs(namespace string) ComputePerInstanceConfigInterface {
	return newComputePerInstanceConfigs(c, namespace)
}

func (c *ComputeV1alpha1Client) ComputeRegionAutoscalers(namespace string) ComputeRegionAutoscalerInterface {
	return newComputeRegionAutoscalers(c, namespace)
}

func (c *ComputeV1alpha1Client) ComputeRegionDiskResourcePolicyAttachments(namespace string) ComputeRegionDiskResourcePolicyAttachmentInterface {
	return newComputeRegionDiskResourcePolicyAttachments(c, namespace)
}

func (c *ComputeV1alpha1Client) ComputeRegionPerInstanceConfigs(namespace string) ComputeRegionPerInstanceConfigInterface {
	return newComputeRegionPerInstanceConfigs(c, namespace)
}

// NewForConfig creates a new ComputeV1alpha1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*ComputeV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new ComputeV1alpha1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*ComputeV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &ComputeV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new ComputeV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ComputeV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ComputeV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *ComputeV1alpha1Client {
	return &ComputeV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ComputeV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
