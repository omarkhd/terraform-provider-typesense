package typesense

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

var testClusterName = fmt.Sprintf("test-%d", time.Now().Unix())

func TestAccOrderResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + fmt.Sprintf(`
resource "typesense_cluster" "test" {
	name = "%s"
	memory = "0.5_gb"
	vcpu = "2_vcpus_1_hr_burst_per_day"
	region = "oregon"
	auto_upgrade_capacity = true
}
`, testClusterName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("typesense_cluster.test", "auto_upgrade_capacity", "true"),
					resource.TestCheckResourceAttr("typesense_cluster.test", "high_availability", "no"),
					resource.TestCheckResourceAttr("typesense_cluster.test", "high_performance_disk", "no"),
					//resource.TestCheckResourceAttr("typesense_cluster.test", "id", "05umtgeli2v8b19np"),
					resource.TestCheckResourceAttr("typesense_cluster.test", "load_balancing", "no"),
					resource.TestCheckResourceAttr("typesense_cluster.test", "memory", "0.5_gb"),
					resource.TestCheckResourceAttr("typesense_cluster.test", "name", testClusterName),
					resource.TestCheckResourceAttr("typesense_cluster.test", "region", "oregon"),
					resource.TestCheckResourceAttr("typesense_cluster.test", "search_delivery_network", "off"),
					resource.TestCheckResourceAttr("typesense_cluster.test", "status", "initializing"),
					resource.TestCheckResourceAttr("typesense_cluster.test", "typesense_server_version", "0.24.1"),
					resource.TestCheckResourceAttr("typesense_cluster.test", "vcpu", "2_vcpus_1_hr_burst_per_day"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "typesense_cluster.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: providerConfig + fmt.Sprintf(`
resource "typesense_cluster" "test" {
  name = "%s_tmp"
  memory = "0.5_gb"
  vcpu = "2_vcpus_1_hr_burst_per_day"
  region = "oregon"
  auto_upgrade_capacity = false
}
`, testClusterName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("typesense_cluster.test", "auto_upgrade_capacity", "false"),
					resource.TestCheckResourceAttr("typesense_cluster.test", "name", testClusterName+"_tmp"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
