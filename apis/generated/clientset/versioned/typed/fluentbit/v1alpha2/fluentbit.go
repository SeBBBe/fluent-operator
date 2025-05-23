/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	"context"
	"time"

	v1alpha2 "github.com/SeBBBe/fluent-operator/v3/apis/fluentbit/v1alpha2"
	scheme "github.com/SeBBBe/fluent-operator/v3/apis/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FluentBitsGetter has a method to return a FluentBitInterface.
// A group's client should implement this interface.
type FluentBitsGetter interface {
	FluentBits(namespace string) FluentBitInterface
}

// FluentBitInterface has methods to work with FluentBit resources.
type FluentBitInterface interface {
	Create(ctx context.Context, fluentBit *v1alpha2.FluentBit, opts v1.CreateOptions) (*v1alpha2.FluentBit, error)
	Update(ctx context.Context, fluentBit *v1alpha2.FluentBit, opts v1.UpdateOptions) (*v1alpha2.FluentBit, error)
	UpdateStatus(ctx context.Context, fluentBit *v1alpha2.FluentBit, opts v1.UpdateOptions) (*v1alpha2.FluentBit, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.FluentBit, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.FluentBitList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.FluentBit, err error)
	FluentBitExpansion
}

// fluentBits implements FluentBitInterface
type fluentBits struct {
	client rest.Interface
	ns     string
}

// newFluentBits returns a FluentBits
func newFluentBits(c *FluentbitV1alpha2Client, namespace string) *fluentBits {
	return &fluentBits{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the fluentBit, and returns the corresponding fluentBit object, and an error if there is any.
func (c *fluentBits) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.FluentBit, err error) {
	result = &v1alpha2.FluentBit{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("fluentbits").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FluentBits that match those selectors.
func (c *fluentBits) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.FluentBitList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.FluentBitList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("fluentbits").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested fluentBits.
func (c *fluentBits) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("fluentbits").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a fluentBit and creates it.  Returns the server's representation of the fluentBit, and an error, if there is any.
func (c *fluentBits) Create(ctx context.Context, fluentBit *v1alpha2.FluentBit, opts v1.CreateOptions) (result *v1alpha2.FluentBit, err error) {
	result = &v1alpha2.FluentBit{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("fluentbits").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fluentBit).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a fluentBit and updates it. Returns the server's representation of the fluentBit, and an error, if there is any.
func (c *fluentBits) Update(ctx context.Context, fluentBit *v1alpha2.FluentBit, opts v1.UpdateOptions) (result *v1alpha2.FluentBit, err error) {
	result = &v1alpha2.FluentBit{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("fluentbits").
		Name(fluentBit.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fluentBit).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *fluentBits) UpdateStatus(ctx context.Context, fluentBit *v1alpha2.FluentBit, opts v1.UpdateOptions) (result *v1alpha2.FluentBit, err error) {
	result = &v1alpha2.FluentBit{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("fluentbits").
		Name(fluentBit.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fluentBit).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the fluentBit and deletes it. Returns an error if one occurs.
func (c *fluentBits) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("fluentbits").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *fluentBits) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("fluentbits").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched fluentBit.
func (c *fluentBits) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.FluentBit, err error) {
	result = &v1alpha2.FluentBit{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("fluentbits").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
