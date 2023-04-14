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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/orgpolicy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOrgPolicyCustomConstraints implements OrgPolicyCustomConstraintInterface
type FakeOrgPolicyCustomConstraints struct {
	Fake *FakeOrgpolicyV1alpha1
	ns   string
}

var orgpolicycustomconstraintsResource = schema.GroupVersionResource{Group: "orgpolicy.cnrm.cloud.google.com", Version: "v1alpha1", Resource: "orgpolicycustomconstraints"}

var orgpolicycustomconstraintsKind = schema.GroupVersionKind{Group: "orgpolicy.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "OrgPolicyCustomConstraint"}

// Get takes name of the orgPolicyCustomConstraint, and returns the corresponding orgPolicyCustomConstraint object, and an error if there is any.
func (c *FakeOrgPolicyCustomConstraints) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OrgPolicyCustomConstraint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(orgpolicycustomconstraintsResource, c.ns, name), &v1alpha1.OrgPolicyCustomConstraint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrgPolicyCustomConstraint), err
}

// List takes label and field selectors, and returns the list of OrgPolicyCustomConstraints that match those selectors.
func (c *FakeOrgPolicyCustomConstraints) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OrgPolicyCustomConstraintList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(orgpolicycustomconstraintsResource, orgpolicycustomconstraintsKind, c.ns, opts), &v1alpha1.OrgPolicyCustomConstraintList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OrgPolicyCustomConstraintList{ListMeta: obj.(*v1alpha1.OrgPolicyCustomConstraintList).ListMeta}
	for _, item := range obj.(*v1alpha1.OrgPolicyCustomConstraintList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested orgPolicyCustomConstraints.
func (c *FakeOrgPolicyCustomConstraints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(orgpolicycustomconstraintsResource, c.ns, opts))

}

// Create takes the representation of a orgPolicyCustomConstraint and creates it.  Returns the server's representation of the orgPolicyCustomConstraint, and an error, if there is any.
func (c *FakeOrgPolicyCustomConstraints) Create(ctx context.Context, orgPolicyCustomConstraint *v1alpha1.OrgPolicyCustomConstraint, opts v1.CreateOptions) (result *v1alpha1.OrgPolicyCustomConstraint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(orgpolicycustomconstraintsResource, c.ns, orgPolicyCustomConstraint), &v1alpha1.OrgPolicyCustomConstraint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrgPolicyCustomConstraint), err
}

// Update takes the representation of a orgPolicyCustomConstraint and updates it. Returns the server's representation of the orgPolicyCustomConstraint, and an error, if there is any.
func (c *FakeOrgPolicyCustomConstraints) Update(ctx context.Context, orgPolicyCustomConstraint *v1alpha1.OrgPolicyCustomConstraint, opts v1.UpdateOptions) (result *v1alpha1.OrgPolicyCustomConstraint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(orgpolicycustomconstraintsResource, c.ns, orgPolicyCustomConstraint), &v1alpha1.OrgPolicyCustomConstraint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrgPolicyCustomConstraint), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOrgPolicyCustomConstraints) UpdateStatus(ctx context.Context, orgPolicyCustomConstraint *v1alpha1.OrgPolicyCustomConstraint, opts v1.UpdateOptions) (*v1alpha1.OrgPolicyCustomConstraint, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(orgpolicycustomconstraintsResource, "status", c.ns, orgPolicyCustomConstraint), &v1alpha1.OrgPolicyCustomConstraint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrgPolicyCustomConstraint), err
}

// Delete takes name of the orgPolicyCustomConstraint and deletes it. Returns an error if one occurs.
func (c *FakeOrgPolicyCustomConstraints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(orgpolicycustomconstraintsResource, c.ns, name, opts), &v1alpha1.OrgPolicyCustomConstraint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOrgPolicyCustomConstraints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(orgpolicycustomconstraintsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.OrgPolicyCustomConstraintList{})
	return err
}

// Patch applies the patch and returns the patched orgPolicyCustomConstraint.
func (c *FakeOrgPolicyCustomConstraints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OrgPolicyCustomConstraint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(orgpolicycustomconstraintsResource, c.ns, name, pt, data, subresources...), &v1alpha1.OrgPolicyCustomConstraint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrgPolicyCustomConstraint), err
}
