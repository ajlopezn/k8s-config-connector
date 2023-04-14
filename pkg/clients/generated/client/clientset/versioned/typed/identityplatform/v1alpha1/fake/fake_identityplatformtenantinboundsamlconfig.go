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

package fake

import (
	"context"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/identityplatform/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIdentityPlatformTenantInboundSAMLConfigs implements IdentityPlatformTenantInboundSAMLConfigInterface
type FakeIdentityPlatformTenantInboundSAMLConfigs struct {
	Fake *FakeIdentityplatformV1alpha1
	ns   string
}

var identityplatformtenantinboundsamlconfigsResource = schema.GroupVersionResource{Group: "identityplatform.cnrm.cloud.google.com", Version: "v1alpha1", Resource: "identityplatformtenantinboundsamlconfigs"}

var identityplatformtenantinboundsamlconfigsKind = schema.GroupVersionKind{Group: "identityplatform.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "IdentityPlatformTenantInboundSAMLConfig"}

// Get takes name of the identityPlatformTenantInboundSAMLConfig, and returns the corresponding identityPlatformTenantInboundSAMLConfig object, and an error if there is any.
func (c *FakeIdentityPlatformTenantInboundSAMLConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IdentityPlatformTenantInboundSAMLConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(identityplatformtenantinboundsamlconfigsResource, c.ns, name), &v1alpha1.IdentityPlatformTenantInboundSAMLConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityPlatformTenantInboundSAMLConfig), err
}

// List takes label and field selectors, and returns the list of IdentityPlatformTenantInboundSAMLConfigs that match those selectors.
func (c *FakeIdentityPlatformTenantInboundSAMLConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IdentityPlatformTenantInboundSAMLConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(identityplatformtenantinboundsamlconfigsResource, identityplatformtenantinboundsamlconfigsKind, c.ns, opts), &v1alpha1.IdentityPlatformTenantInboundSAMLConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IdentityPlatformTenantInboundSAMLConfigList{ListMeta: obj.(*v1alpha1.IdentityPlatformTenantInboundSAMLConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.IdentityPlatformTenantInboundSAMLConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested identityPlatformTenantInboundSAMLConfigs.
func (c *FakeIdentityPlatformTenantInboundSAMLConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(identityplatformtenantinboundsamlconfigsResource, c.ns, opts))

}

// Create takes the representation of a identityPlatformTenantInboundSAMLConfig and creates it.  Returns the server's representation of the identityPlatformTenantInboundSAMLConfig, and an error, if there is any.
func (c *FakeIdentityPlatformTenantInboundSAMLConfigs) Create(ctx context.Context, identityPlatformTenantInboundSAMLConfig *v1alpha1.IdentityPlatformTenantInboundSAMLConfig, opts v1.CreateOptions) (result *v1alpha1.IdentityPlatformTenantInboundSAMLConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(identityplatformtenantinboundsamlconfigsResource, c.ns, identityPlatformTenantInboundSAMLConfig), &v1alpha1.IdentityPlatformTenantInboundSAMLConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityPlatformTenantInboundSAMLConfig), err
}

// Update takes the representation of a identityPlatformTenantInboundSAMLConfig and updates it. Returns the server's representation of the identityPlatformTenantInboundSAMLConfig, and an error, if there is any.
func (c *FakeIdentityPlatformTenantInboundSAMLConfigs) Update(ctx context.Context, identityPlatformTenantInboundSAMLConfig *v1alpha1.IdentityPlatformTenantInboundSAMLConfig, opts v1.UpdateOptions) (result *v1alpha1.IdentityPlatformTenantInboundSAMLConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(identityplatformtenantinboundsamlconfigsResource, c.ns, identityPlatformTenantInboundSAMLConfig), &v1alpha1.IdentityPlatformTenantInboundSAMLConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityPlatformTenantInboundSAMLConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIdentityPlatformTenantInboundSAMLConfigs) UpdateStatus(ctx context.Context, identityPlatformTenantInboundSAMLConfig *v1alpha1.IdentityPlatformTenantInboundSAMLConfig, opts v1.UpdateOptions) (*v1alpha1.IdentityPlatformTenantInboundSAMLConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(identityplatformtenantinboundsamlconfigsResource, "status", c.ns, identityPlatformTenantInboundSAMLConfig), &v1alpha1.IdentityPlatformTenantInboundSAMLConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityPlatformTenantInboundSAMLConfig), err
}

// Delete takes name of the identityPlatformTenantInboundSAMLConfig and deletes it. Returns an error if one occurs.
func (c *FakeIdentityPlatformTenantInboundSAMLConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(identityplatformtenantinboundsamlconfigsResource, c.ns, name, opts), &v1alpha1.IdentityPlatformTenantInboundSAMLConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIdentityPlatformTenantInboundSAMLConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(identityplatformtenantinboundsamlconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IdentityPlatformTenantInboundSAMLConfigList{})
	return err
}

// Patch applies the patch and returns the patched identityPlatformTenantInboundSAMLConfig.
func (c *FakeIdentityPlatformTenantInboundSAMLConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IdentityPlatformTenantInboundSAMLConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(identityplatformtenantinboundsamlconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.IdentityPlatformTenantInboundSAMLConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IdentityPlatformTenantInboundSAMLConfig), err
}
