package colors

import (
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Color struct {
	Name string
	Code int
}

func bareColorSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
				"code": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}
}

func flattenColors(c *[]Color, d *schema.ResourceData) {
	cs := make([]interface{}, len(*c))

	for i, color := range *c {
		flattenedColor := make(map[string]interface{})
		flattenedColor["name"] = color.Name
		flattenedColor["code"] = strconv.FormatInt(int64(color.Code), 16)

		cs[i] = flattenedColor
	}

	d.Set("colors", &cs)

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
}
