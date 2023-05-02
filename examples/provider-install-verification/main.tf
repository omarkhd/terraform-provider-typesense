terraform {
  required_providers {
    typesense = {
      source = "omarkhd.net/terraform/typesense"
    }
  }
}

provider "typesense" {
  key = "foobar"
}

data "typesense_cluster" "example" {}
