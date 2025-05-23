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

// MultilineParsersGetter has a method to return a MultilineParserInterface.
// A group's client should implement this interface.
type MultilineParsersGetter interface {
	MultilineParsers(namespace string) MultilineParserInterface
}

// MultilineParserInterface has methods to work with MultilineParser resources.
type MultilineParserInterface interface {
	Create(ctx context.Context, multilineParser *v1alpha2.MultilineParser, opts v1.CreateOptions) (*v1alpha2.MultilineParser, error)
	Update(ctx context.Context, multilineParser *v1alpha2.MultilineParser, opts v1.UpdateOptions) (*v1alpha2.MultilineParser, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.MultilineParser, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.MultilineParserList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.MultilineParser, err error)
	MultilineParserExpansion
}

// multilineParsers implements MultilineParserInterface
type multilineParsers struct {
	client rest.Interface
	ns     string
}

// newMultilineParsers returns a MultilineParsers
func newMultilineParsers(c *FluentbitV1alpha2Client, namespace string) *multilineParsers {
	return &multilineParsers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the multilineParser, and returns the corresponding multilineParser object, and an error if there is any.
func (c *multilineParsers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.MultilineParser, err error) {
	result = &v1alpha2.MultilineParser{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("multilineparsers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MultilineParsers that match those selectors.
func (c *multilineParsers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.MultilineParserList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.MultilineParserList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("multilineparsers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested multilineParsers.
func (c *multilineParsers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("multilineparsers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a multilineParser and creates it.  Returns the server's representation of the multilineParser, and an error, if there is any.
func (c *multilineParsers) Create(ctx context.Context, multilineParser *v1alpha2.MultilineParser, opts v1.CreateOptions) (result *v1alpha2.MultilineParser, err error) {
	result = &v1alpha2.MultilineParser{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("multilineparsers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(multilineParser).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a multilineParser and updates it. Returns the server's representation of the multilineParser, and an error, if there is any.
func (c *multilineParsers) Update(ctx context.Context, multilineParser *v1alpha2.MultilineParser, opts v1.UpdateOptions) (result *v1alpha2.MultilineParser, err error) {
	result = &v1alpha2.MultilineParser{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("multilineparsers").
		Name(multilineParser.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(multilineParser).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the multilineParser and deletes it. Returns an error if one occurs.
func (c *multilineParsers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("multilineparsers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *multilineParsers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("multilineparsers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched multilineParser.
func (c *multilineParsers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.MultilineParser, err error) {
	result = &v1alpha2.MultilineParser{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("multilineparsers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
