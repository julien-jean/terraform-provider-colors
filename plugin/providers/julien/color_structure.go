package julien

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

var resourceVLANOptionalFields = linearSearchSlice{
	"l2_domain_id",
	"description",
}

func bareVLANSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"vlan_id": &schema.Schema{
			Type: schema.TypeInt,
		},
		"l2_domain_id": &schema.Schema{
			Type: schema.TypeInt,
		},
		"name": &schema.Schema{
			Type: schema.TypeString,
		},
		"number": &schema.Schema{
			Type: schema.TypeInt,
		},
		"description": &schema.Schema{
			Type: schema.TypeString,
		},
		"edit_date": &schema.Schema{
			Type: schema.TypeString,
		},
		"custom_fields": &schema.Schema{
			Type: schema.TypeMap,
		},
	}
}

func resourceVLANSchema() map[string]*schema.Schema {
	schema := bareVLANSchema()
	for k, v := range schema {
		switch {
		// VLAN name and number are required
		case k == "name" || k == "number":
			v.Required = true
		case k == "custom_fields":
			v.Optional = true
		case resourceVLANOptionalFields.Has(k):
			v.Optional = true
			v.Computed = true
		default:
			v.Computed = true
		}
	}
	return schema
}
