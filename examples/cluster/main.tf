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

output "example_cluster_id" {
  value = data.typesense_cluster.example
}
