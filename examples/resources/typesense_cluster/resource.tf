# Cluster with minimun required attributes.
resource "typesense_cluster" "example" {
  memory = "0.5_gb"
  vcpu = "2_vcpus_1_hr_burst_per_day"
  region = "oregon"
}
