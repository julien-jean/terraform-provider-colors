package julien

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"color": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "#FF0000",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"julien_color": resourceJulienColor(),
		},
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"color": "The color to use",
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	return nil, errors.New("oops")
}
