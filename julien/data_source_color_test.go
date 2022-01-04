package julien

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const test3Colors = `
data "julien_color" "three_colors" {
	number = 3
}`

func TestAccDataSourceJulienColor(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: test3Colors,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.julien_color", "three_colors", "3"),
				),
			},
		},
	})
}
