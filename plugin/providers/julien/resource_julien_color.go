package julien

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceJulienColor() *schema.Resource {
	return &schema.Resource{
		Create: resourceJulienColorCreate,
		Read:   resourceJulienColorCreate,
		Update: resourceJulienColorCreate,
		Delete: resourceJulienColorCreate,
		Schema: resourceVLANSchema(),
	}
}

func resourceJulienColorCreate(d *schema.ResourceData, meta interface{}) error {
	return errors.New("ho no")
}
