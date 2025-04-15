package custom

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/SeBBBe/fluent-operator/v3/apis/fluentd/v1alpha1/plugins"
	"github.com/SeBBBe/fluent-operator/v3/apis/fluentd/v1alpha1/plugins/params"
)

type CustomPlugin struct {
	Config string `json:"config"`
}

func (c *CustomPlugin) Name() string {
	return ""
}

func (c *CustomPlugin) Params(_ plugins.SecretLoader) (*params.PluginStore, error) {
	ps := params.NewPluginStore("")
	ps.Content = indentation(c.Config)
	return ps, nil
}

func indentation(config string) string {
	splits := strings.Split(config, "\n")
	var buf bytes.Buffer
	for _, split := range splits {
		if split != "" {
			buf.WriteString(fmt.Sprintf("  %s\n", split))
		}
	}
	return buf.String()
}
