package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"terraform-provider-typesense/typesense"
)

// Provider documentation generation.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name typesense

func main() {
	if err := providerserver.Serve(context.Background(), typesense.New, providerserver.ServeOpts{
		Address: "omarkhd.net/terraform/typesense",
	}); err != nil {
		log.Fatal(err.Error())
	}
}
