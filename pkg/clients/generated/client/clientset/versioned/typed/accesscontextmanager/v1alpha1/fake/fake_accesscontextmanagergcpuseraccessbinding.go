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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/accesscontextmanager/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAccessContextManagerGCPUserAccessBindings implements AccessContextManagerGCPUserAccessBindingInterface
type FakeAccessContextManagerGCPUserAccessBindings struct {
	Fake *FakeAccesscontextmanagerV1alpha1
	ns   string
}

var accesscontextmanagergcpuseraccessbindingsResource = schema.GroupVersionResource{Group: "accesscontextmanager.cnrm.cloud.google.com", Version: "v1alpha1", Resource: "accesscontextmanagergcpuseraccessbindings"}

var accesscontextmanagergcpuseraccessbindingsKind = schema.GroupVersionKind{Group: "accesscontextmanager.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "AccessContextManagerGCPUserAccessBinding"}

// Get takes name of the accessContextManagerGCPUserAccessBinding, and returns the corresponding accessContextManagerGCPUserAccessBinding object, and an error if there is any.
func (c *FakeAccessContextManagerGCPUserAccessBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AccessContextManagerGCPUserAccessBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(accesscontextmanagergcpuseraccessbindingsResource, c.ns, name), &v1alpha1.AccessContextManagerGCPUserAccessBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessContextManagerGCPUserAccessBinding), err
}

// List takes label and field selectors, and returns the list of AccessContextManagerGCPUserAccessBindings that match those selectors.
func (c *FakeAccessContextManagerGCPUserAccessBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AccessContextManagerGCPUserAccessBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(accesscontextmanagergcpuseraccessbindingsResource, accesscontextmanagergcpuseraccessbindingsKind, c.ns, opts), &v1alpha1.AccessContextManagerGCPUserAccessBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AccessContextManagerGCPUserAccessBindingList{ListMeta: obj.(*v1alpha1.AccessContextManagerGCPUserAccessBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.AccessContextManagerGCPUserAccessBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested accessContextManagerGCPUserAccessBindings.
func (c *FakeAccessContextManagerGCPUserAccessBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(accesscontextmanagergcpuseraccessbindingsResource, c.ns, opts))

}

// Create takes the representation of a accessContextManagerGCPUserAccessBinding and creates it.  Returns the server's representation of the accessContextManagerGCPUserAccessBinding, and an error, if there is any.
func (c *FakeAccessContextManagerGCPUserAccessBindings) Create(ctx context.Context, accessContextManagerGCPUserAccessBinding *v1alpha1.AccessContextManagerGCPUserAccessBinding, opts v1.CreateOptions) (result *v1alpha1.AccessContextManagerGCPUserAccessBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(accesscontextmanagergcpuseraccessbindingsResource, c.ns, accessContextManagerGCPUserAccessBinding), &v1alpha1.AccessContextManagerGCPUserAccessBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessContextManagerGCPUserAccessBinding), err
}

// Update takes the representation of a accessContextManagerGCPUserAccessBinding and updates it. Returns the server's representation of the accessContextManagerGCPUserAccessBinding, and an error, if there is any.
func (c *FakeAccessContextManagerGCPUserAccessBindings) Update(ctx context.Context, accessContextManagerGCPUserAccessBinding *v1alpha1.AccessContextManagerGCPUserAccessBinding, opts v1.UpdateOptions) (result *v1alpha1.AccessContextManagerGCPUserAccessBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(accesscontextmanagergcpuseraccessbindingsResource, c.ns, accessContextManagerGCPUserAccessBinding), &v1alpha1.AccessContextManagerGCPUserAccessBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessContextManagerGCPUserAccessBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAccessContextManagerGCPUserAccessBindings) UpdateStatus(ctx context.Context, accessContextManagerGCPUserAccessBinding *v1alpha1.AccessContextManagerGCPUserAccessBinding, opts v1.UpdateOptions) (*v1alpha1.AccessContextManagerGCPUserAccessBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(accesscontextmanagergcpuseraccessbindingsResource, "status", c.ns, accessContextManagerGCPUserAccessBinding), &v1alpha1.AccessContextManagerGCPUserAccessBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessContextManagerGCPUserAccessBinding), err
}

// Delete takes name of the accessContextManagerGCPUserAccessBinding and deletes it. Returns an error if one occurs.
func (c *FakeAccessContextManagerGCPUserAccessBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(accesscontextmanagergcpuseraccessbindingsResource, c.ns, name, opts), &v1alpha1.AccessContextManagerGCPUserAccessBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAccessContextManagerGCPUserAccessBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(accesscontextmanagergcpuseraccessbindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AccessContextManagerGCPUserAccessBindingList{})
	return err
}

// Patch applies the patch and returns the patched accessContextManagerGCPUserAccessBinding.
func (c *FakeAccessContextManagerGCPUserAccessBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AccessContextManagerGCPUserAccessBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(accesscontextmanagergcpuseraccessbindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AccessContextManagerGCPUserAccessBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessContextManagerGCPUserAccessBinding), err
}
