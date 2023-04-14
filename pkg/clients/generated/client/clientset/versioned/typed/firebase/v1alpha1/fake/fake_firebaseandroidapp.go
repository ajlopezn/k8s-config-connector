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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/firebase/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFirebaseAndroidApps implements FirebaseAndroidAppInterface
type FakeFirebaseAndroidApps struct {
	Fake *FakeFirebaseV1alpha1
	ns   string
}

var firebaseandroidappsResource = schema.GroupVersionResource{Group: "firebase.cnrm.cloud.google.com", Version: "v1alpha1", Resource: "firebaseandroidapps"}

var firebaseandroidappsKind = schema.GroupVersionKind{Group: "firebase.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "FirebaseAndroidApp"}

// Get takes name of the firebaseAndroidApp, and returns the corresponding firebaseAndroidApp object, and an error if there is any.
func (c *FakeFirebaseAndroidApps) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FirebaseAndroidApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(firebaseandroidappsResource, c.ns, name), &v1alpha1.FirebaseAndroidApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirebaseAndroidApp), err
}

// List takes label and field selectors, and returns the list of FirebaseAndroidApps that match those selectors.
func (c *FakeFirebaseAndroidApps) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FirebaseAndroidAppList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(firebaseandroidappsResource, firebaseandroidappsKind, c.ns, opts), &v1alpha1.FirebaseAndroidAppList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FirebaseAndroidAppList{ListMeta: obj.(*v1alpha1.FirebaseAndroidAppList).ListMeta}
	for _, item := range obj.(*v1alpha1.FirebaseAndroidAppList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested firebaseAndroidApps.
func (c *FakeFirebaseAndroidApps) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(firebaseandroidappsResource, c.ns, opts))

}

// Create takes the representation of a firebaseAndroidApp and creates it.  Returns the server's representation of the firebaseAndroidApp, and an error, if there is any.
func (c *FakeFirebaseAndroidApps) Create(ctx context.Context, firebaseAndroidApp *v1alpha1.FirebaseAndroidApp, opts v1.CreateOptions) (result *v1alpha1.FirebaseAndroidApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(firebaseandroidappsResource, c.ns, firebaseAndroidApp), &v1alpha1.FirebaseAndroidApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirebaseAndroidApp), err
}

// Update takes the representation of a firebaseAndroidApp and updates it. Returns the server's representation of the firebaseAndroidApp, and an error, if there is any.
func (c *FakeFirebaseAndroidApps) Update(ctx context.Context, firebaseAndroidApp *v1alpha1.FirebaseAndroidApp, opts v1.UpdateOptions) (result *v1alpha1.FirebaseAndroidApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(firebaseandroidappsResource, c.ns, firebaseAndroidApp), &v1alpha1.FirebaseAndroidApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirebaseAndroidApp), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFirebaseAndroidApps) UpdateStatus(ctx context.Context, firebaseAndroidApp *v1alpha1.FirebaseAndroidApp, opts v1.UpdateOptions) (*v1alpha1.FirebaseAndroidApp, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(firebaseandroidappsResource, "status", c.ns, firebaseAndroidApp), &v1alpha1.FirebaseAndroidApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirebaseAndroidApp), err
}

// Delete takes name of the firebaseAndroidApp and deletes it. Returns an error if one occurs.
func (c *FakeFirebaseAndroidApps) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(firebaseandroidappsResource, c.ns, name, opts), &v1alpha1.FirebaseAndroidApp{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFirebaseAndroidApps) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(firebaseandroidappsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FirebaseAndroidAppList{})
	return err
}

// Patch applies the patch and returns the patched firebaseAndroidApp.
func (c *FakeFirebaseAndroidApps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FirebaseAndroidApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(firebaseandroidappsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FirebaseAndroidApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirebaseAndroidApp), err
}
