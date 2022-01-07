package colors

import (
	"context"
	"math/rand"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRandomColors() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRandomColorsRead,
		Schema: map[string]*schema.Schema{
			"number": {
				Type:     schema.TypeInt,
				Default:  1,
				Optional: true,
			},
			"colors": bareColorSchema(),
		},
	}
}

func dataSourceRandomColorsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	nb := d.Get("number").(int)
	colors := []Color{}
	rand.Seed(time.Now().Unix())
	for i := 1; i <= nb; i++ {
		color := RandomColor()
		colors = append(colors, color)
	}

	flattenColors(&colors, d)

	return nil
}
