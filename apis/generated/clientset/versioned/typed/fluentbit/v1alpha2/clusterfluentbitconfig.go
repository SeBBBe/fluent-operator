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

// ClusterFluentBitConfigsGetter has a method to return a ClusterFluentBitConfigInterface.
// A group's client should implement this interface.
type ClusterFluentBitConfigsGetter interface {
	ClusterFluentBitConfigs() ClusterFluentBitConfigInterface
}

// ClusterFluentBitConfigInterface has methods to work with ClusterFluentBitConfig resources.
type ClusterFluentBitConfigInterface interface {
	Create(ctx context.Context, clusterFluentBitConfig *v1alpha2.ClusterFluentBitConfig, opts v1.CreateOptions) (*v1alpha2.ClusterFluentBitConfig, error)
	Update(ctx context.Context, clusterFluentBitConfig *v1alpha2.ClusterFluentBitConfig, opts v1.UpdateOptions) (*v1alpha2.ClusterFluentBitConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.ClusterFluentBitConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.ClusterFluentBitConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.ClusterFluentBitConfig, err error)
	ClusterFluentBitConfigExpansion
}

// clusterFluentBitConfigs implements ClusterFluentBitConfigInterface
type clusterFluentBitConfigs struct {
	client rest.Interface
}

// newClusterFluentBitConfigs returns a ClusterFluentBitConfigs
func newClusterFluentBitConfigs(c *FluentbitV1alpha2Client) *clusterFluentBitConfigs {
	return &clusterFluentBitConfigs{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterFluentBitConfig, and returns the corresponding clusterFluentBitConfig object, and an error if there is any.
func (c *clusterFluentBitConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.ClusterFluentBitConfig, err error) {
	result = &v1alpha2.ClusterFluentBitConfig{}
	err = c.client.Get().
		Resource("clusterfluentbitconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterFluentBitConfigs that match those selectors.
func (c *clusterFluentBitConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.ClusterFluentBitConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.ClusterFluentBitConfigList{}
	err = c.client.Get().
		Resource("clusterfluentbitconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterFluentBitConfigs.
func (c *clusterFluentBitConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusterfluentbitconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterFluentBitConfig and creates it.  Returns the server's representation of the clusterFluentBitConfig, and an error, if there is any.
func (c *clusterFluentBitConfigs) Create(ctx context.Context, clusterFluentBitConfig *v1alpha2.ClusterFluentBitConfig, opts v1.CreateOptions) (result *v1alpha2.ClusterFluentBitConfig, err error) {
	result = &v1alpha2.ClusterFluentBitConfig{}
	err = c.client.Post().
		Resource("clusterfluentbitconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterFluentBitConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterFluentBitConfig and updates it. Returns the server's representation of the clusterFluentBitConfig, and an error, if there is any.
func (c *clusterFluentBitConfigs) Update(ctx context.Context, clusterFluentBitConfig *v1alpha2.ClusterFluentBitConfig, opts v1.UpdateOptions) (result *v1alpha2.ClusterFluentBitConfig, err error) {
	result = &v1alpha2.ClusterFluentBitConfig{}
	err = c.client.Put().
		Resource("clusterfluentbitconfigs").
		Name(clusterFluentBitConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterFluentBitConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterFluentBitConfig and deletes it. Returns an error if one occurs.
func (c *clusterFluentBitConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterfluentbitconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterFluentBitConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusterfluentbitconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterFluentBitConfig.
func (c *clusterFluentBitConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.ClusterFluentBitConfig, err error) {
	result = &v1alpha2.ClusterFluentBitConfig{}
	err = c.client.Patch(pt).
		Resource("clusterfluentbitconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
