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

// FakeAccessContextManagerAccessLevelConditions implements AccessContextManagerAccessLevelConditionInterface
type FakeAccessContextManagerAccessLevelConditions struct {
	Fake *FakeAccesscontextmanagerV1alpha1
	ns   string
}

var accesscontextmanageraccesslevelconditionsResource = schema.GroupVersionResource{Group: "accesscontextmanager.cnrm.cloud.google.com", Version: "v1alpha1", Resource: "accesscontextmanageraccesslevelconditions"}

var accesscontextmanageraccesslevelconditionsKind = schema.GroupVersionKind{Group: "accesscontextmanager.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "AccessContextManagerAccessLevelCondition"}

// Get takes name of the accessContextManagerAccessLevelCondition, and returns the corresponding accessContextManagerAccessLevelCondition object, and an error if there is any.
func (c *FakeAccessContextManagerAccessLevelConditions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AccessContextManagerAccessLevelCondition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(accesscontextmanageraccesslevelconditionsResource, c.ns, name), &v1alpha1.AccessContextManagerAccessLevelCondition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessContextManagerAccessLevelCondition), err
}

// List takes label and field selectors, and returns the list of AccessContextManagerAccessLevelConditions that match those selectors.
func (c *FakeAccessContextManagerAccessLevelConditions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AccessContextManagerAccessLevelConditionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(accesscontextmanageraccesslevelconditionsResource, accesscontextmanageraccesslevelconditionsKind, c.ns, opts), &v1alpha1.AccessContextManagerAccessLevelConditionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AccessContextManagerAccessLevelConditionList{ListMeta: obj.(*v1alpha1.AccessContextManagerAccessLevelConditionList).ListMeta}
	for _, item := range obj.(*v1alpha1.AccessContextManagerAccessLevelConditionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested accessContextManagerAccessLevelConditions.
func (c *FakeAccessContextManagerAccessLevelConditions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(accesscontextmanageraccesslevelconditionsResource, c.ns, opts))

}

// Create takes the representation of a accessContextManagerAccessLevelCondition and creates it.  Returns the server's representation of the accessContextManagerAccessLevelCondition, and an error, if there is any.
func (c *FakeAccessContextManagerAccessLevelConditions) Create(ctx context.Context, accessContextManagerAccessLevelCondition *v1alpha1.AccessContextManagerAccessLevelCondition, opts v1.CreateOptions) (result *v1alpha1.AccessContextManagerAccessLevelCondition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(accesscontextmanageraccesslevelconditionsResource, c.ns, accessContextManagerAccessLevelCondition), &v1alpha1.AccessContextManagerAccessLevelCondition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessContextManagerAccessLevelCondition), err
}

// Update takes the representation of a accessContextManagerAccessLevelCondition and updates it. Returns the server's representation of the accessContextManagerAccessLevelCondition, and an error, if there is any.
func (c *FakeAccessContextManagerAccessLevelConditions) Update(ctx context.Context, accessContextManagerAccessLevelCondition *v1alpha1.AccessContextManagerAccessLevelCondition, opts v1.UpdateOptions) (result *v1alpha1.AccessContextManagerAccessLevelCondition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(accesscontextmanageraccesslevelconditionsResource, c.ns, accessContextManagerAccessLevelCondition), &v1alpha1.AccessContextManagerAccessLevelCondition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessContextManagerAccessLevelCondition), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAccessContextManagerAccessLevelConditions) UpdateStatus(ctx context.Context, accessContextManagerAccessLevelCondition *v1alpha1.AccessContextManagerAccessLevelCondition, opts v1.UpdateOptions) (*v1alpha1.AccessContextManagerAccessLevelCondition, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(accesscontextmanageraccesslevelconditionsResource, "status", c.ns, accessContextManagerAccessLevelCondition), &v1alpha1.AccessContextManagerAccessLevelCondition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessContextManagerAccessLevelCondition), err
}

// Delete takes name of the accessContextManagerAccessLevelCondition and deletes it. Returns an error if one occurs.
func (c *FakeAccessContextManagerAccessLevelConditions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(accesscontextmanageraccesslevelconditionsResource, c.ns, name, opts), &v1alpha1.AccessContextManagerAccessLevelCondition{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAccessContextManagerAccessLevelConditions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(accesscontextmanageraccesslevelconditionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AccessContextManagerAccessLevelConditionList{})
	return err
}

// Patch applies the patch and returns the patched accessContextManagerAccessLevelCondition.
func (c *FakeAccessContextManagerAccessLevelConditions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AccessContextManagerAccessLevelCondition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(accesscontextmanageraccesslevelconditionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AccessContextManagerAccessLevelCondition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessContextManagerAccessLevelCondition), err
}
