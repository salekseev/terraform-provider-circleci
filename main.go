package main

import (
	"github.com/samanthaq/terraform-provider-circleci/circleci"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: circleci.Provider})
}
