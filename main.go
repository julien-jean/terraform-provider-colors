package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/julien-jean/terraform-provider/plugin/providers/julien"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: julien.Provider,
	})
}
