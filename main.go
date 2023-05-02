package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"terraform-provider-typesense/typesense"
)

func main() {
	if err := providerserver.Serve(context.Background(), typesense.New, providerserver.ServeOpts{
		Address: "omarkhd.net/terraform/typesense",
	}); err != nil {
		log.Fatal(err.Error())
	}
}
