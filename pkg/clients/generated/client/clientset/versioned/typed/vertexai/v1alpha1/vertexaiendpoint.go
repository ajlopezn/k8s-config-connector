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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/vertexai/v1alpha1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VertexAIEndpointsGetter has a method to return a VertexAIEndpointInterface.
// A group's client should implement this interface.
type VertexAIEndpointsGetter interface {
	VertexAIEndpoints(namespace string) VertexAIEndpointInterface
}

// VertexAIEndpointInterface has methods to work with VertexAIEndpoint resources.
type VertexAIEndpointInterface interface {
	Create(ctx context.Context, vertexAIEndpoint *v1alpha1.VertexAIEndpoint, opts v1.CreateOptions) (*v1alpha1.VertexAIEndpoint, error)
	Update(ctx context.Context, vertexAIEndpoint *v1alpha1.VertexAIEndpoint, opts v1.UpdateOptions) (*v1alpha1.VertexAIEndpoint, error)
	UpdateStatus(ctx context.Context, vertexAIEndpoint *v1alpha1.VertexAIEndpoint, opts v1.UpdateOptions) (*v1alpha1.VertexAIEndpoint, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.VertexAIEndpoint, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VertexAIEndpointList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VertexAIEndpoint, err error)
	VertexAIEndpointExpansion
}

// vertexAIEndpoints implements VertexAIEndpointInterface
type vertexAIEndpoints struct {
	client rest.Interface
	ns     string
}

// newVertexAIEndpoints returns a VertexAIEndpoints
func newVertexAIEndpoints(c *VertexaiV1alpha1Client, namespace string) *vertexAIEndpoints {
	return &vertexAIEndpoints{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the vertexAIEndpoint, and returns the corresponding vertexAIEndpoint object, and an error if there is any.
func (c *vertexAIEndpoints) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VertexAIEndpoint, err error) {
	result = &v1alpha1.VertexAIEndpoint{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vertexaiendpoints").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VertexAIEndpoints that match those selectors.
func (c *vertexAIEndpoints) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VertexAIEndpointList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VertexAIEndpointList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vertexaiendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vertexAIEndpoints.
func (c *vertexAIEndpoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("vertexaiendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a vertexAIEndpoint and creates it.  Returns the server's representation of the vertexAIEndpoint, and an error, if there is any.
func (c *vertexAIEndpoints) Create(ctx context.Context, vertexAIEndpoint *v1alpha1.VertexAIEndpoint, opts v1.CreateOptions) (result *v1alpha1.VertexAIEndpoint, err error) {
	result = &v1alpha1.VertexAIEndpoint{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("vertexaiendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vertexAIEndpoint).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a vertexAIEndpoint and updates it. Returns the server's representation of the vertexAIEndpoint, and an error, if there is any.
func (c *vertexAIEndpoints) Update(ctx context.Context, vertexAIEndpoint *v1alpha1.VertexAIEndpoint, opts v1.UpdateOptions) (result *v1alpha1.VertexAIEndpoint, err error) {
	result = &v1alpha1.VertexAIEndpoint{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vertexaiendpoints").
		Name(vertexAIEndpoint.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vertexAIEndpoint).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *vertexAIEndpoints) UpdateStatus(ctx context.Context, vertexAIEndpoint *v1alpha1.VertexAIEndpoint, opts v1.UpdateOptions) (result *v1alpha1.VertexAIEndpoint, err error) {
	result = &v1alpha1.VertexAIEndpoint{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vertexaiendpoints").
		Name(vertexAIEndpoint.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vertexAIEndpoint).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the vertexAIEndpoint and deletes it. Returns an error if one occurs.
func (c *vertexAIEndpoints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vertexaiendpoints").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vertexAIEndpoints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vertexaiendpoints").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched vertexAIEndpoint.
func (c *vertexAIEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VertexAIEndpoint, err error) {
	result = &v1alpha1.VertexAIEndpoint{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("vertexaiendpoints").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
