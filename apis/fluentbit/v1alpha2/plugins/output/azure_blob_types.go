package output

import (
	"github.com/SeBBBe/fluent-operator/v3/apis/fluentbit/v1alpha2/plugins"
	"github.com/SeBBBe/fluent-operator/v3/apis/fluentbit/v1alpha2/plugins/params"
)

// +kubebuilder:object:generate:=true

// Azure Blob is the Azure Blob output plugin, allows to ingest your records into Azure Blob Storage. <br />
// **For full documentation, refer to https://docs.fluentbit.io/manual/pipeline/outputs/azure_blob**
type AzureBlob struct {
	// Azure Storage account name
	AccountName string `json:"accountName"`
	// Specify the Azure Storage Shared Key to authenticate against the storage account
	SharedKey *plugins.Secret `json:"sharedKey"`
	// Name of the container that will contain the blobs
	ContainerName string `json:"containerName"`
	// Specify the desired blob type. Must be `appendblob` or `blockblob`
	// +kubebuilder:validation:Enum:=appendblob;blockblob
	BlobType string `json:"blobType,omitempty"`
	// Creates container if ContainerName is not set.
	// +kubebuilder:validation:Enum:=on;off
	AutoCreateContainer string `json:"autoCreateContainer,omitempty"`
	// Optional path to store the blobs.
	Path string `json:"path,omitempty"`
	// Optional toggle to use an Azure emulator
	// +kubebuilder:validation:Enum:=on;off
	EmulatorMode string `json:"emulatorMode,omitempty"`
	// HTTP Service of the endpoint (if using EmulatorMode)
	Endpoint string `json:"endpoint,omitempty"`
	// Enable/Disable TLS Encryption. Azure services require TLS to be enabled.
	*plugins.TLS `json:"tls,omitempty"`
	// Include fluentbit networking options for this output-plugin
	*plugins.Networking `json:"networking,omitempty"`
}

// Name implement Section() method
func (_ *AzureBlob) Name() string {
	return "azure_blob"
}

// Params implement Section() method
func (o *AzureBlob) Params(sl plugins.SecretLoader) (*params.KVs, error) {
	kvs := params.NewKVs()
	if o.AccountName != "" {
		kvs.Insert("account_name", o.AccountName)
	}
	if o.SharedKey != nil {
		u, err := sl.LoadSecret(*o.SharedKey)
		if err != nil {
			return nil, err
		}
		kvs.Insert("shared_key", u)
	}
	if o.ContainerName != "" {
		kvs.Insert("container_name", o.ContainerName)
	}
	if o.BlobType != "" {
		kvs.Insert("blob_type", o.BlobType)
	}
	if o.AutoCreateContainer != "" {
		kvs.Insert("auto_create_container", o.AutoCreateContainer)
	}
	if o.Path != "" {
		kvs.Insert("path", o.Path)
	}
	if o.EmulatorMode != "" {
		kvs.Insert("emulator_mode", o.EmulatorMode)
	}
	if o.Endpoint != "" {
		kvs.Insert("endpoint", o.Endpoint)
	}
	if o.TLS != nil {
		tls, err := o.TLS.Params(sl)
		if err != nil {
			return nil, err
		}
		kvs.Merge(tls)
	}
	if o.Networking != nil {
		net, err := o.Networking.Params(sl)
		if err != nil {
			return nil, err
		}
		kvs.Merge(net)
	}
	return kvs, nil
}
