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

package v1alpha1

import (
	"net/http"

	v1alpha1 "github.com/SeBBBe/fluent-operator/v3/apis/fluentd/v1alpha1"
	"github.com/SeBBBe/fluent-operator/v3/apis/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type FluentdV1alpha1Interface interface {
	RESTClient() rest.Interface
	ClusterFiltersGetter
	ClusterFluentdConfigsGetter
	ClusterInputsGetter
	ClusterOutputsGetter
	FiltersGetter
	FluentdsGetter
	FluentdConfigsGetter
	InputsGetter
	OutputsGetter
}

// FluentdV1alpha1Client is used to interact with features provided by the fluentd.fluent.io group.
type FluentdV1alpha1Client struct {
	restClient rest.Interface
}

func (c *FluentdV1alpha1Client) ClusterFilters() ClusterFilterInterface {
	return newClusterFilters(c)
}

func (c *FluentdV1alpha1Client) ClusterFluentdConfigs() ClusterFluentdConfigInterface {
	return newClusterFluentdConfigs(c)
}

func (c *FluentdV1alpha1Client) ClusterInputs() ClusterInputInterface {
	return newClusterInputs(c)
}

func (c *FluentdV1alpha1Client) ClusterOutputs() ClusterOutputInterface {
	return newClusterOutputs(c)
}

func (c *FluentdV1alpha1Client) Filters(namespace string) FilterInterface {
	return newFilters(c, namespace)
}

func (c *FluentdV1alpha1Client) Fluentds(namespace string) FluentdInterface {
	return newFluentds(c, namespace)
}

func (c *FluentdV1alpha1Client) FluentdConfigs(namespace string) FluentdConfigInterface {
	return newFluentdConfigs(c, namespace)
}

func (c *FluentdV1alpha1Client) Inputs(namespace string) InputInterface {
	return newInputs(c, namespace)
}

func (c *FluentdV1alpha1Client) Outputs(namespace string) OutputInterface {
	return newOutputs(c, namespace)
}

// NewForConfig creates a new FluentdV1alpha1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*FluentdV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new FluentdV1alpha1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*FluentdV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &FluentdV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new FluentdV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *FluentdV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new FluentdV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *FluentdV1alpha1Client {
	return &FluentdV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FluentdV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
