terraform {
  required_providers {
    typesense = {
      source = "omarkhd.net/terraform/typesense"
    }
  }
}

provider "typesense" {}

data "typesense_cluster" "example" {}
