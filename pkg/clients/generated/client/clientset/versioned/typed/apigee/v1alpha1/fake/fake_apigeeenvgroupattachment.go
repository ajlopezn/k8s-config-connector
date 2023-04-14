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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/apigee/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApigeeEnvgroupAttachments implements ApigeeEnvgroupAttachmentInterface
type FakeApigeeEnvgroupAttachments struct {
	Fake *FakeApigeeV1alpha1
	ns   string
}

var apigeeenvgroupattachmentsResource = schema.GroupVersionResource{Group: "apigee.cnrm.cloud.google.com", Version: "v1alpha1", Resource: "apigeeenvgroupattachments"}

var apigeeenvgroupattachmentsKind = schema.GroupVersionKind{Group: "apigee.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ApigeeEnvgroupAttachment"}

// Get takes name of the apigeeEnvgroupAttachment, and returns the corresponding apigeeEnvgroupAttachment object, and an error if there is any.
func (c *FakeApigeeEnvgroupAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApigeeEnvgroupAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apigeeenvgroupattachmentsResource, c.ns, name), &v1alpha1.ApigeeEnvgroupAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApigeeEnvgroupAttachment), err
}

// List takes label and field selectors, and returns the list of ApigeeEnvgroupAttachments that match those selectors.
func (c *FakeApigeeEnvgroupAttachments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApigeeEnvgroupAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apigeeenvgroupattachmentsResource, apigeeenvgroupattachmentsKind, c.ns, opts), &v1alpha1.ApigeeEnvgroupAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApigeeEnvgroupAttachmentList{ListMeta: obj.(*v1alpha1.ApigeeEnvgroupAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApigeeEnvgroupAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apigeeEnvgroupAttachments.
func (c *FakeApigeeEnvgroupAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apigeeenvgroupattachmentsResource, c.ns, opts))

}

// Create takes the representation of a apigeeEnvgroupAttachment and creates it.  Returns the server's representation of the apigeeEnvgroupAttachment, and an error, if there is any.
func (c *FakeApigeeEnvgroupAttachments) Create(ctx context.Context, apigeeEnvgroupAttachment *v1alpha1.ApigeeEnvgroupAttachment, opts v1.CreateOptions) (result *v1alpha1.ApigeeEnvgroupAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apigeeenvgroupattachmentsResource, c.ns, apigeeEnvgroupAttachment), &v1alpha1.ApigeeEnvgroupAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApigeeEnvgroupAttachment), err
}

// Update takes the representation of a apigeeEnvgroupAttachment and updates it. Returns the server's representation of the apigeeEnvgroupAttachment, and an error, if there is any.
func (c *FakeApigeeEnvgroupAttachments) Update(ctx context.Context, apigeeEnvgroupAttachment *v1alpha1.ApigeeEnvgroupAttachment, opts v1.UpdateOptions) (result *v1alpha1.ApigeeEnvgroupAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apigeeenvgroupattachmentsResource, c.ns, apigeeEnvgroupAttachment), &v1alpha1.ApigeeEnvgroupAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApigeeEnvgroupAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApigeeEnvgroupAttachments) UpdateStatus(ctx context.Context, apigeeEnvgroupAttachment *v1alpha1.ApigeeEnvgroupAttachment, opts v1.UpdateOptions) (*v1alpha1.ApigeeEnvgroupAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apigeeenvgroupattachmentsResource, "status", c.ns, apigeeEnvgroupAttachment), &v1alpha1.ApigeeEnvgroupAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApigeeEnvgroupAttachment), err
}

// Delete takes name of the apigeeEnvgroupAttachment and deletes it. Returns an error if one occurs.
func (c *FakeApigeeEnvgroupAttachments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(apigeeenvgroupattachmentsResource, c.ns, name, opts), &v1alpha1.ApigeeEnvgroupAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApigeeEnvgroupAttachments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apigeeenvgroupattachmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApigeeEnvgroupAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched apigeeEnvgroupAttachment.
func (c *FakeApigeeEnvgroupAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApigeeEnvgroupAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apigeeenvgroupattachmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApigeeEnvgroupAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApigeeEnvgroupAttachment), err
}
