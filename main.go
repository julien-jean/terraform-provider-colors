package main

import (
	"terraform-provider-colors/colors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: colors.Provider,
	})
}
