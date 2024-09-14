package main

import (
	"terraform-provider-example/tf_provider"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: tf_provider.Provider,
	})
}
