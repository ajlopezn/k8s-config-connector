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
	"context"
	"time"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/bigquery/v1alpha1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BigQueryDatasetAccessesGetter has a method to return a BigQueryDatasetAccessInterface.
// A group's client should implement this interface.
type BigQueryDatasetAccessesGetter interface {
	BigQueryDatasetAccesses(namespace string) BigQueryDatasetAccessInterface
}

// BigQueryDatasetAccessInterface has methods to work with BigQueryDatasetAccess resources.
type BigQueryDatasetAccessInterface interface {
	Create(ctx context.Context, bigQueryDatasetAccess *v1alpha1.BigQueryDatasetAccess, opts v1.CreateOptions) (*v1alpha1.BigQueryDatasetAccess, error)
	Update(ctx context.Context, bigQueryDatasetAccess *v1alpha1.BigQueryDatasetAccess, opts v1.UpdateOptions) (*v1alpha1.BigQueryDatasetAccess, error)
	UpdateStatus(ctx context.Context, bigQueryDatasetAccess *v1alpha1.BigQueryDatasetAccess, opts v1.UpdateOptions) (*v1alpha1.BigQueryDatasetAccess, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.BigQueryDatasetAccess, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.BigQueryDatasetAccessList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BigQueryDatasetAccess, err error)
	BigQueryDatasetAccessExpansion
}

// bigQueryDatasetAccesses implements BigQueryDatasetAccessInterface
type bigQueryDatasetAccesses struct {
	client rest.Interface
	ns     string
}

// newBigQueryDatasetAccesses returns a BigQueryDatasetAccesses
func newBigQueryDatasetAccesses(c *BigqueryV1alpha1Client, namespace string) *bigQueryDatasetAccesses {
	return &bigQueryDatasetAccesses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bigQueryDatasetAccess, and returns the corresponding bigQueryDatasetAccess object, and an error if there is any.
func (c *bigQueryDatasetAccesses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BigQueryDatasetAccess, err error) {
	result = &v1alpha1.BigQueryDatasetAccess{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bigquerydatasetaccesses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BigQueryDatasetAccesses that match those selectors.
func (c *bigQueryDatasetAccesses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BigQueryDatasetAccessList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BigQueryDatasetAccessList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bigquerydatasetaccesses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bigQueryDatasetAccesses.
func (c *bigQueryDatasetAccesses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bigquerydatasetaccesses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a bigQueryDatasetAccess and creates it.  Returns the server's representation of the bigQueryDatasetAccess, and an error, if there is any.
func (c *bigQueryDatasetAccesses) Create(ctx context.Context, bigQueryDatasetAccess *v1alpha1.BigQueryDatasetAccess, opts v1.CreateOptions) (result *v1alpha1.BigQueryDatasetAccess, err error) {
	result = &v1alpha1.BigQueryDatasetAccess{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bigquerydatasetaccesses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bigQueryDatasetAccess).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a bigQueryDatasetAccess and updates it. Returns the server's representation of the bigQueryDatasetAccess, and an error, if there is any.
func (c *bigQueryDatasetAccesses) Update(ctx context.Context, bigQueryDatasetAccess *v1alpha1.BigQueryDatasetAccess, opts v1.UpdateOptions) (result *v1alpha1.BigQueryDatasetAccess, err error) {
	result = &v1alpha1.BigQueryDatasetAccess{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bigquerydatasetaccesses").
		Name(bigQueryDatasetAccess.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bigQueryDatasetAccess).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *bigQueryDatasetAccesses) UpdateStatus(ctx context.Context, bigQueryDatasetAccess *v1alpha1.BigQueryDatasetAccess, opts v1.UpdateOptions) (result *v1alpha1.BigQueryDatasetAccess, err error) {
	result = &v1alpha1.BigQueryDatasetAccess{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bigquerydatasetaccesses").
		Name(bigQueryDatasetAccess.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bigQueryDatasetAccess).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the bigQueryDatasetAccess and deletes it. Returns an error if one occurs.
func (c *bigQueryDatasetAccesses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bigquerydatasetaccesses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bigQueryDatasetAccesses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bigquerydatasetaccesses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched bigQueryDatasetAccess.
func (c *bigQueryDatasetAccesses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BigQueryDatasetAccess, err error) {
	result = &v1alpha1.BigQueryDatasetAccess{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bigquerydatasetaccesses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
