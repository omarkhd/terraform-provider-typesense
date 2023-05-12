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
  memory = "0.5_gb"
  vcpu = "2_vcpus_1_hr_burst_per_day"
  region = "oregon"
  name = "example"
}
