package julien

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"julien_color": dataSourceColor(),
		},
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"color": "The color to use",
	}
}

var diags diag.Diagnostics

func providerConfigure(d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return nil, diags
}
