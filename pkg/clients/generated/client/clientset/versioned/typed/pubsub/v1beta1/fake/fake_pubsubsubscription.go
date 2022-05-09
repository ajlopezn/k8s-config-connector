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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/pubsub/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePubSubSubscriptions implements PubSubSubscriptionInterface
type FakePubSubSubscriptions struct {
	Fake *FakePubsubV1beta1
	ns   string
}

var pubsubsubscriptionsResource = schema.GroupVersionResource{Group: "pubsub.cnrm.cloud.google.com", Version: "v1beta1", Resource: "pubsubsubscriptions"}

var pubsubsubscriptionsKind = schema.GroupVersionKind{Group: "pubsub.cnrm.cloud.google.com", Version: "v1beta1", Kind: "PubSubSubscription"}

// Get takes name of the pubSubSubscription, and returns the corresponding pubSubSubscription object, and an error if there is any.
func (c *FakePubSubSubscriptions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.PubSubSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pubsubsubscriptionsResource, c.ns, name), &v1beta1.PubSubSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PubSubSubscription), err
}

// List takes label and field selectors, and returns the list of PubSubSubscriptions that match those selectors.
func (c *FakePubSubSubscriptions) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.PubSubSubscriptionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pubsubsubscriptionsResource, pubsubsubscriptionsKind, c.ns, opts), &v1beta1.PubSubSubscriptionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.PubSubSubscriptionList{ListMeta: obj.(*v1beta1.PubSubSubscriptionList).ListMeta}
	for _, item := range obj.(*v1beta1.PubSubSubscriptionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pubSubSubscriptions.
func (c *FakePubSubSubscriptions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pubsubsubscriptionsResource, c.ns, opts))

}

// Create takes the representation of a pubSubSubscription and creates it.  Returns the server's representation of the pubSubSubscription, and an error, if there is any.
func (c *FakePubSubSubscriptions) Create(ctx context.Context, pubSubSubscription *v1beta1.PubSubSubscription, opts v1.CreateOptions) (result *v1beta1.PubSubSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pubsubsubscriptionsResource, c.ns, pubSubSubscription), &v1beta1.PubSubSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PubSubSubscription), err
}

// Update takes the representation of a pubSubSubscription and updates it. Returns the server's representation of the pubSubSubscription, and an error, if there is any.
func (c *FakePubSubSubscriptions) Update(ctx context.Context, pubSubSubscription *v1beta1.PubSubSubscription, opts v1.UpdateOptions) (result *v1beta1.PubSubSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pubsubsubscriptionsResource, c.ns, pubSubSubscription), &v1beta1.PubSubSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PubSubSubscription), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePubSubSubscriptions) UpdateStatus(ctx context.Context, pubSubSubscription *v1beta1.PubSubSubscription, opts v1.UpdateOptions) (*v1beta1.PubSubSubscription, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pubsubsubscriptionsResource, "status", c.ns, pubSubSubscription), &v1beta1.PubSubSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PubSubSubscription), err
}

// Delete takes name of the pubSubSubscription and deletes it. Returns an error if one occurs.
func (c *FakePubSubSubscriptions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(pubsubsubscriptionsResource, c.ns, name, opts), &v1beta1.PubSubSubscription{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePubSubSubscriptions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pubsubsubscriptionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.PubSubSubscriptionList{})
	return err
}

// Patch applies the patch and returns the patched pubSubSubscription.
func (c *FakePubSubSubscriptions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.PubSubSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pubsubsubscriptionsResource, c.ns, name, pt, data, subresources...), &v1beta1.PubSubSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PubSubSubscription), err
}