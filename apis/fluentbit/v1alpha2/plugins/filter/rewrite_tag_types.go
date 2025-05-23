package filter

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/SeBBBe/fluent-operator/v3/apis/fluentbit/v1alpha2/plugins"
	"github.com/SeBBBe/fluent-operator/v3/apis/fluentbit/v1alpha2/plugins/params"
)

// +kubebuilder:object:generate:=true

// RewriteTag define a `rewrite_tag` filter, allows to re-emit a record under a new Tag. <br />
// Once a record has been re-emitted, the original record can be preserved or discarded. <br />
// **For full documentation, refer to https://docs.fluentbit.io/manual/pipeline/filters/rewrite-tag**
type RewriteTag struct {
	plugins.CommonParams `json:",inline"`
	// Defines the matching criteria and the format of the Tag for the matching record.
	// The Rule format have four components: KEY REGEX NEW_TAG KEEP.
	Rules []string `json:"rules,omitempty"`
	// When the filter emits a record under the new Tag, there is an internal emitter
	// plugin that takes care of the job. Since this emitter expose metrics as any other
	// component of the pipeline, you can use this property to configure an optional name for it.
	EmitterName        string `json:"emitterName,omitempty"`
	EmitterMemBufLimit string `json:"emitterMemBufLimit,omitempty"`
	EmitterStorageType string `json:"emitterStorageType,omitempty"`
}

func (_ *RewriteTag) Name() string {
	return "rewrite_tag"
}

func (r *RewriteTag) Params(_ plugins.SecretLoader) (*params.KVs, error) {
	kvs := params.NewKVs()
	err := r.AddCommonParams(kvs)
	if err != nil {
		return kvs, err
	}
	for _, rule := range r.Rules {
		kvs.Insert("Rule", rule)
	}
	if r.EmitterName != "" {
		kvs.Insert("Emitter_Name", r.EmitterName)
	}
	if r.EmitterMemBufLimit != "" {
		kvs.Insert("Emitter_Mem_Buf_Limit", r.EmitterMemBufLimit)
	}
	if r.EmitterStorageType != "" {
		kvs.Insert("Emitter_Storage.type", r.EmitterStorageType)
	}
	return kvs, nil
}

func (r *RewriteTag) MakeNamespaced(ns string) {
	for idx, rule := range r.Rules {
		parts := strings.Fields(rule)
		parts[2] = fmt.Sprintf("%x.%s", md5.Sum([]byte(ns)), parts[2])
		r.Rules[idx] = strings.Join(parts, " ")
	}
}
