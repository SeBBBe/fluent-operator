package plugins

import "github.com/SeBBBe/fluent-operator/v3/apis/fluentbit/v1alpha2/plugins/params"

// +kubebuilder:object:generate=false

// The Plugin interface defines methods for transferring input, filter
// and output plugins to textual section content.
type Plugin interface {
	Name() string
	Params(SecretLoader) (*params.KVs, error)
}

// The Namespaceable interface defines a method for adding a namespace
// to a plugins identifier.
type Namespaceable interface {
	MakeNamespaced(string)
}
