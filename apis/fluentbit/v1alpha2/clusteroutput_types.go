/*
Copyright 2021.

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

package v1alpha2

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"

	"github.com/SeBBBe/fluent-operator/v3/apis/fluentbit/v1alpha2/plugins"
	"github.com/SeBBBe/fluent-operator/v3/apis/fluentbit/v1alpha2/plugins/custom"
	"github.com/SeBBBe/fluent-operator/v3/apis/fluentbit/v1alpha2/plugins/output"
	"github.com/SeBBBe/fluent-operator/v3/pkg/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// OutputSpec defines the desired state of ClusterOutput
type OutputSpec struct {
	// A pattern to match against the tags of incoming records.
	// It's case sensitive and support the star (*) character as a wildcard.
	Match string `json:"match,omitempty"`
	// A regular expression to match against the tags of incoming records.
	// Use this option if you want to use the full regex syntax.
	MatchRegex string `json:"matchRegex,omitempty"`
	// A user friendly alias name for this output plugin.
	// Used in metrics for distinction of each configured output.
	Alias string `json:"alias,omitempty"`
	// Set the plugin's logging verbosity level. Allowed values are: off, error, warn, info, debug and trace, Defaults to the SERVICE section's Log_Level
	// +kubebuilder:validation:Enum:=off;error;warning;info;debug;trace
	LogLevel string `json:"logLevel,omitempty"`

	// AzureBlob defines AzureBlob Output Configuration
	AzureBlob *output.AzureBlob `json:"azureBlob,omitempty"`
	// AzureLogAnalytics defines AzureLogAnalytics Output Configuration
	AzureLogAnalytics *output.AzureLogAnalytics `json:"azureLogAnalytics,omitempty"`
	// CloudWatch defines CloudWatch Output Configuration
	CloudWatch *output.CloudWatch `json:"cloudWatch,omitempty"`
	// RetryLimit represents configuration for the scheduler which can be set independently on each output section.
	// This option allows to disable retries or impose a limit to try N times and then discard the data after reaching that limit.
	RetryLimit string `json:"retry_limit,omitempty"`
	// Elasticsearch defines Elasticsearch Output configuration.
	Elasticsearch *output.Elasticsearch `json:"es,omitempty"`
	// File defines File Output configuration.
	File *output.File `json:"file,omitempty"`
	// Forward defines Forward Output configuration.
	Forward *output.Forward `json:"forward,omitempty"`
	// HTTP defines HTTP Output configuration.
	HTTP *output.HTTP `json:"http,omitempty"`
	// Kafka defines Kafka Output configuration.
	Kafka *output.Kafka `json:"kafka,omitempty"`
	// Null defines Null Output configuration.
	Null *output.Null `json:"null,omitempty"`
	// Stdout defines Stdout Output configuration.
	Stdout *output.Stdout `json:"stdout,omitempty"`
	// TCP defines TCP Output configuration.
	TCP *output.TCP `json:"tcp,omitempty"`
	// Loki defines Loki Output configuration.
	Loki *output.Loki `json:"loki,omitempty"`
	// Syslog defines Syslog Output configuration.
	Syslog *output.Syslog `json:"syslog,omitempty"`
	// InfluxDB defines InfluxDB Output configuration.
	InfluxDB *output.InfluxDB `json:"influxDB,omitempty"`
	// DataDog defines DataDog Output configuration.
	DataDog *output.DataDog `json:"datadog,omitempty"`
	// Firehose defines Firehose Output configuration.
	Firehose *output.Firehose `json:"firehose,omitempty"`
	// Kinesis defines Kinesis Output configuration.
	Kinesis *output.Kinesis `json:"kinesis,omitempty"`
	// Stackdriver defines Stackdriver Output Configuration
	Stackdriver *output.Stackdriver `json:"stackdriver,omitempty"`
	// Splunk defines Splunk Output Configuration
	Splunk *output.Splunk `json:"splunk,omitempty"`
	// OpenSearch defines OpenSearch Output configuration.
	OpenSearch *output.OpenSearch `json:"opensearch,omitempty"`
	// OpenTelemetry defines OpenTelemetry Output configuration.
	OpenTelemetry *output.OpenTelemetry `json:"opentelemetry,omitempty"`
	// PrometheusExporter_types defines Prometheus exporter configuration to expose metrics from Fluent Bit.
	PrometheusExporter *output.PrometheusExporter `json:"prometheusExporter,omitempty"`
	// PrometheusRemoteWrite_types defines Prometheus Remote Write configuration.
	PrometheusRemoteWrite *output.PrometheusRemoteWrite `json:"prometheusRemoteWrite,omitempty"`
	// S3 defines S3 Output configuration.
	S3 *output.S3 `json:"s3,omitempty"`
	// Gelf defines GELF Output configuration.
	Gelf *output.Gelf `json:"gelf,omitempty"`

	// CustomPlugin defines Custom Output configuration.
	CustomPlugin *custom.CustomPlugin `json:"customPlugin,omitempty"`

	// Processors defines the processors configuration
	// +kubebuilder:pruning:PreserveUnknownFields
	Processors *plugins.Config `json:"processors,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=cfbo,scope=Cluster
// +genclient
// +genclient:nonNamespaced

// ClusterOutput is the Schema for the cluster-level outputs API
type ClusterOutput struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec OutputSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterOutputList contains a list of ClusterOutput
type ClusterOutputList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterOutput `json:"items"`
}

// +kubebuilder:object:generate:=false

// OutputByName implements sort.Interface for []ClusterOutput based on the Name field.
type OutputByName []ClusterOutput

func (a OutputByName) Len() int           { return len(a) }
func (a OutputByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a OutputByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func (list ClusterOutputList) Load(sl plugins.SecretLoader) (string, error) {
	var buf bytes.Buffer

	sort.Sort(OutputByName(list.Items))

	for _, item := range list.Items {
		merge := func(p plugins.Plugin) error {
			if p == nil || reflect.ValueOf(p).IsNil() {
				return nil
			}

			buf.WriteString("[Output]\n")
			if p.Name() != "" {
				buf.WriteString(fmt.Sprintf("    Name    %s\n", p.Name()))
			}
			if item.Spec.Match != "" {
				buf.WriteString(fmt.Sprintf("    Match    %s\n", item.Spec.Match))
			}
			if item.Spec.LogLevel != "" {
				buf.WriteString(fmt.Sprintf("    Log_Level    %s\n", item.Spec.LogLevel))
			}
			if item.Spec.MatchRegex != "" {
				buf.WriteString(fmt.Sprintf("    Match_Regex    %s\n", item.Spec.MatchRegex))
			}
			if item.Spec.Alias != "" {
				buf.WriteString(fmt.Sprintf("    Alias    %s\n", item.Spec.Alias))
			}
			if item.Spec.RetryLimit != "" {
				buf.WriteString(fmt.Sprintf("    Retry_Limit    %s\n", item.Spec.RetryLimit))
			}
			kvs, err := p.Params(sl)
			if err != nil {
				return err
			}
			buf.WriteString(kvs.String())
			return nil
		}

		for i := 2; i < reflect.ValueOf(item.Spec).NumField(); i++ {
			p, _ := reflect.ValueOf(item.Spec).Field(i).Interface().(plugins.Plugin)
			if err := merge(p); err != nil {
				return "", err
			}
		}
	}

	return buf.String(), nil
}

func (list ClusterOutputList) LoadAsYaml(sl plugins.SecretLoader, depth int) (string, error) {
	var buf bytes.Buffer

	sort.Sort(OutputByName(list.Items))
	buf.WriteString(fmt.Sprintf("%soutputs:\n", utils.YamlIndent(depth)))
	padding := utils.YamlIndent(depth + 2)
	for _, item := range list.Items {
		merge := func(p plugins.Plugin) error {
			if p == nil || reflect.ValueOf(p).IsNil() {
				return nil
			}

			if p.Name() != "" {
				buf.WriteString(fmt.Sprintf("%s- name: %s\n", utils.YamlIndent(depth+1), p.Name()))
			}
			if item.Spec.Match != "" {
				buf.WriteString(fmt.Sprintf("%smatch: \"%s\"\n", padding, item.Spec.Match))
			}
			if item.Spec.LogLevel != "" {
				buf.WriteString(fmt.Sprintf("%slog_level: %s\n", padding, item.Spec.LogLevel))
			}
			if item.Spec.MatchRegex != "" {
				buf.WriteString(fmt.Sprintf("%smatch_regex: %s\n", padding, item.Spec.MatchRegex))
			}
			if item.Spec.Alias != "" {
				buf.WriteString(fmt.Sprintf("%salias: %s\n", padding, item.Spec.Alias))
			}
			if item.Spec.RetryLimit != "" {
				buf.WriteString(fmt.Sprintf("%sretry_limit: %s\n", padding, item.Spec.RetryLimit))
			}
			if item.Spec.Processors != nil {
				buf.WriteString(fmt.Sprintf("%sprocessors:\n", padding))
				processorYaml, err := yaml.Marshal(item.Spec.Processors)
				if err != nil {
					return fmt.Errorf("error marshalling processor: %w", err)
				}
				buf.WriteString(utils.AdjustYamlIndent(string(processorYaml), depth+4))
			}
			kvs, err := p.Params(sl)
			if err != nil {
				return err
			}
			buf.WriteString(kvs.YamlString(depth + 2))
			return nil
		}

		for i := 2; i < reflect.ValueOf(item.Spec).NumField(); i++ {
			p, _ := reflect.ValueOf(item.Spec).Field(i).Interface().(plugins.Plugin)
			if err := merge(p); err != nil {
				return "", err
			}
		}
	}

	return buf.String(), nil
}
func init() {
	SchemeBuilder.Register(&ClusterOutput{}, &ClusterOutputList{})
}
