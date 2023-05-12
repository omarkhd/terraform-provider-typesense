terraform {
  required_providers {
    typesense = {
      source = "omarkhd.net/terraform/typesense"
    }
  }
}

provider "typesense" {
}

resource "typesense_cluster" "example" {
  name = "example"
}
