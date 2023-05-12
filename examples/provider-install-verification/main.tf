terraform {
  required_providers {
    typesense = {
      source = "omarkhd.net/terraform/typesense"
    }
  }
}

provider "typesense" {
}

data "typesense_cluster" "example" {
  id = "05umtgeli2v8b19np"
}

output "example_cluster" {
  value = data.typesense_cluster.example
}
