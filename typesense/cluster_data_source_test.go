package typesense

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCoffeesDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + `
data "typesense_cluster" "test" {
  id = "05umtgeli2v8b19np"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify against a sandbox cluster to ensure all attributes are set
					resource.TestCheckResourceAttr("data.typesense_cluster.test", "auto_upgrade_capacity", "false"),
					resource.TestCheckResourceAttr("data.typesense_cluster.test", "high_availability", "no"),
					resource.TestCheckResourceAttr("data.typesense_cluster.test", "high_performance_disk", "no"),
					resource.TestCheckResourceAttr("data.typesense_cluster.test", "id", "05umtgeli2v8b19np"),
					resource.TestCheckResourceAttr("data.typesense_cluster.test", "load_balancing", "no"),
					resource.TestCheckResourceAttr("data.typesense_cluster.test", "memory", "0.5_gb"),
					resource.TestCheckResourceAttr("data.typesense_cluster.test", "name", "sandbox"),
					resource.TestCheckResourceAttr("data.typesense_cluster.test", "region", "n_california"),
					resource.TestCheckResourceAttr("data.typesense_cluster.test", "search_delivery_network", "off"),
					resource.TestCheckResourceAttr("data.typesense_cluster.test", "status", "in_service"),
					resource.TestCheckResourceAttr("data.typesense_cluster.test", "typesense_server_version", "0.24.1"),
					resource.TestCheckResourceAttr("data.typesense_cluster.test", "vcpu", "2_vcpus_1_hr_burst_per_day"),
				),
			},
		},
	})
}
