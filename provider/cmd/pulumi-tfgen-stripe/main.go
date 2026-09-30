package main

import (
	pftfgen "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfgen"

	stripe_provider "github.com/nellisauction/pulumi-stripe/provider"
)

func main() {
	pftfgen.Main("stripe", stripe_provider.Provider())
}
