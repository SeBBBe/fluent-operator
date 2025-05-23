package output

import (
	"fmt"
	"strings"

	"github.com/SeBBBe/fluent-operator/v3/apis/fluentbit/v1alpha2/plugins"
	"github.com/SeBBBe/fluent-operator/v3/apis/fluentbit/v1alpha2/plugins/params"
)

// +kubebuilder:object:generate:=true

// The loki output plugin, allows to ingest your records into a Loki service. <br />
// **For full documentation, refer to https://docs.fluentbit.io/manual/pipeline/outputs/loki**
type Loki struct {
	// Loki hostname or IP address.
	Host string `json:"host"`
	// Loki TCP port
	// +kubebuilder:validation:Minimum:=1
	// +kubebuilder:validation:Maximum:=65535
	Port *int32 `json:"port,omitempty"`
	// Specify a custom HTTP URI. It must start with forward slash.
	Uri string `json:"uri,omitempty"`
	// Set HTTP basic authentication user name.
	HTTPUser *plugins.Secret `json:"httpUser,omitempty"`
	// Password for user defined in HTTP_User
	// Set HTTP basic authentication password
	HTTPPasswd *plugins.Secret `json:"httpPassword,omitempty"`
	// Set bearer token authentication token value.
	// Can be used as alterntative to HTTP basic authentication
	BearerToken *plugins.Secret `json:"bearerToken,omitempty"`
	// Tenant ID used by default to push logs to Loki.
	// If omitted or empty it assumes Loki is running in single-tenant mode and no X-Scope-OrgID header is sent.
	TenantID *plugins.Secret `json:"tenantID,omitempty"`
	// Stream labels for API request. It can be multiple comma separated of strings specifying  key=value pairs.
	// In addition to fixed parameters, it also allows to add custom record keys (similar to label_keys property).
	Labels []string `json:"labels,omitempty"`
	// Optional list of record keys that will be placed as stream labels.
	// This configuration property is for records key only.
	LabelKeys []string `json:"labelKeys,omitempty"`
	// Specify the label map file path. The file defines how to extract labels from each record.
	LabelMapPath string `json:"labelMapPath,omitempty"`
	// Optional list of keys to remove.
	RemoveKeys []string `json:"removeKeys,omitempty"`
	// If set to true and after extracting labels only a single key remains, the log line sent to Loki will be the value of that key in line_format.
	// +kubebuilder:validation:Enum:=on;off
	DropSingleKey string `json:"dropSingleKey,omitempty"`
	// Format to use when flattening the record to a log line. Valid values are json or key_value.
	// If set to json,  the log line sent to Loki will be the Fluent Bit record dumped as JSON.
	// If set to key_value, the log line will be each item in the record concatenated together (separated by a single space) in the format.
	// +kubebuilder:validation:Enum:=json;key_value
	LineFormat string `json:"lineFormat,omitempty"`
	// If set to true, it will add all Kubernetes labels to the Stream labels.
	// +kubebuilder:validation:Enum:=on;off
	AutoKubernetesLabels string `json:"autoKubernetesLabels,omitempty"`
	// Specify the name of the key from the original record that contains the Tenant ID.
	// The value of the key is set as X-Scope-OrgID of HTTP header. It is useful to set Tenant ID dynamically.
	TenantIDKey  string `json:"tenantIDKey,omitempty"`
	*plugins.TLS `json:"tls,omitempty"`
	// Include fluentbit networking options for this output-plugin
	*plugins.Networking `json:"networking,omitempty"`
	// Limit the maximum number of Chunks in the filesystem for the current output logical destination.
	TotalLimitSize string `json:"totalLimitSize,omitempty"`
	// Enables dedicated thread(s) for this output. Default value is set since version 1.8.13. For previous versions is 0.
	Workers *int32 `json:"workers,omitempty"`
}

// implement Section() method
func (_ *Loki) Name() string {
	return "loki"
}

// implement Section() method
func (l *Loki) Params(sl plugins.SecretLoader) (*params.KVs, error) {
	kvs := params.NewKVs()
	if l.Host != "" {
		kvs.Insert("host", l.Host)
	}
	if l.Port != nil {
		kvs.Insert("port", fmt.Sprint(*l.Port))
	}
	if l.Uri != "" {
		kvs.Insert("uri", l.Uri)
	}
	if l.HTTPUser != nil {
		u, err := sl.LoadSecret(*l.HTTPUser)
		if err != nil {
			return nil, err
		}
		kvs.Insert("http_user", u)
	}
	if l.HTTPPasswd != nil {
		pwd, err := sl.LoadSecret(*l.HTTPPasswd)
		if err != nil {
			return nil, err
		}
		kvs.Insert("http_passwd", pwd)
	}
	if l.BearerToken != nil {
		bearerToken, err := sl.LoadSecret(*l.BearerToken)
		if err != nil {
			return nil, err
		}
		kvs.Insert("bearer_token", bearerToken)
	}
	if l.TenantID != nil {
		id, err := sl.LoadSecret(*l.TenantID)
		if err != nil {
			return nil, err
		}
		kvs.Insert("tenant_id", id)
	}
	if l.Labels != nil && len(l.Labels) > 0 {
		kvs.Insert("labels", strings.Join(l.Labels, ","))
	}
	if l.LabelKeys != nil && len(l.LabelKeys) > 0 {
		kvs.Insert("label_keys", strings.Join(l.LabelKeys, ","))
	}
	if l.LabelMapPath != "" {
		kvs.Insert("label_map_path", l.LabelMapPath)
	}
	if l.RemoveKeys != nil && len(l.RemoveKeys) > 0 {
		kvs.Insert("remove_keys", strings.Join(l.RemoveKeys, ","))
	}
	if l.DropSingleKey != "" {
		kvs.Insert("drop_single_key", l.DropSingleKey)
	}
	if l.LineFormat != "" {
		kvs.Insert("line_format", l.LineFormat)
	}
	if l.AutoKubernetesLabels != "" {
		kvs.Insert("auto_kubernetes_labels", l.AutoKubernetesLabels)
	}
	if l.TenantIDKey != "" {
		kvs.Insert("tenant_id_key", l.TenantIDKey)
	}
	if l.TLS != nil {
		tls, err := l.TLS.Params(sl)
		if err != nil {
			return nil, err
		}
		kvs.Merge(tls)
	}
	if l.Networking != nil {
		net, err := l.Networking.Params(sl)
		if err != nil {
			return nil, err
		}
		kvs.Merge(net)
	}
	if l.TotalLimitSize != "" {
		kvs.Insert("storage.total_limit_size", l.TotalLimitSize)
	}
	if l.Workers != nil {
		kvs.Insert("workers", fmt.Sprint(*l.Workers))
	}
	return kvs, nil
}
