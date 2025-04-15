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

package fake

import (
	"context"

	v1alpha1 "github.com/SeBBBe/fluent-operator/v3/apis/fluentd/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFluentds implements FluentdInterface
type FakeFluentds struct {
	Fake *FakeFluentdV1alpha1
	ns   string
}

var fluentdsResource = schema.GroupVersionResource{Group: "fluentd.fluent.io", Version: "v1alpha1", Resource: "fluentds"}

var fluentdsKind = schema.GroupVersionKind{Group: "fluentd.fluent.io", Version: "v1alpha1", Kind: "Fluentd"}

// Get takes name of the fluentd, and returns the corresponding fluentd object, and an error if there is any.
func (c *FakeFluentds) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Fluentd, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(fluentdsResource, c.ns, name), &v1alpha1.Fluentd{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Fluentd), err
}

// List takes label and field selectors, and returns the list of Fluentds that match those selectors.
func (c *FakeFluentds) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FluentdList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(fluentdsResource, fluentdsKind, c.ns, opts), &v1alpha1.FluentdList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FluentdList{ListMeta: obj.(*v1alpha1.FluentdList).ListMeta}
	for _, item := range obj.(*v1alpha1.FluentdList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fluentds.
func (c *FakeFluentds) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(fluentdsResource, c.ns, opts))

}

// Create takes the representation of a fluentd and creates it.  Returns the server's representation of the fluentd, and an error, if there is any.
func (c *FakeFluentds) Create(ctx context.Context, fluentd *v1alpha1.Fluentd, opts v1.CreateOptions) (result *v1alpha1.Fluentd, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(fluentdsResource, c.ns, fluentd), &v1alpha1.Fluentd{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Fluentd), err
}

// Update takes the representation of a fluentd and updates it. Returns the server's representation of the fluentd, and an error, if there is any.
func (c *FakeFluentds) Update(ctx context.Context, fluentd *v1alpha1.Fluentd, opts v1.UpdateOptions) (result *v1alpha1.Fluentd, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(fluentdsResource, c.ns, fluentd), &v1alpha1.Fluentd{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Fluentd), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFluentds) UpdateStatus(ctx context.Context, fluentd *v1alpha1.Fluentd, opts v1.UpdateOptions) (*v1alpha1.Fluentd, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(fluentdsResource, "status", c.ns, fluentd), &v1alpha1.Fluentd{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Fluentd), err
}

// Delete takes name of the fluentd and deletes it. Returns an error if one occurs.
func (c *FakeFluentds) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(fluentdsResource, c.ns, name, opts), &v1alpha1.Fluentd{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFluentds) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(fluentdsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FluentdList{})
	return err
}

// Patch applies the patch and returns the patched fluentd.
func (c *FakeFluentds) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Fluentd, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fluentdsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Fluentd{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Fluentd), err
}
