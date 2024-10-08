package subgraph_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/wundergraph/cosmo/terraform-provider-cosmo/internal/acceptance"
)

func TestAccSubgraphResource(t *testing.T) {
	rName := acctest.RandomWithPrefix("test-subgraph")
	rNamespace := "default"
	rBaseSubgraphName := "base-subgraph"
	routingURL := "https://example.com"
	updatedRoutingURL := "https://updated-example.com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acceptance.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acceptance.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSubgraphResourceConfig(rName, rNamespace, rBaseSubgraphName, routingURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cosmo_subgraph.test", "name", rName),
					resource.TestCheckResourceAttr("cosmo_subgraph.test", "namespace", rNamespace),
					resource.TestCheckResourceAttr("cosmo_subgraph.test", "base_subgraph_name", rBaseSubgraphName),
					resource.TestCheckResourceAttr("cosmo_subgraph.test", "routing_url", routingURL),
					resource.TestCheckResourceAttr("cosmo_subgraph.test", "labels.#", "2"),
				),
			},
			{
				Config: testAccSubgraphResourceConfig(rName, rNamespace, rBaseSubgraphName, updatedRoutingURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cosmo_subgraph.test", "routing_url", updatedRoutingURL),
				),
			},
			{
				ResourceName: "cosmo_subgraph.test",
				RefreshState: true,
			},
		},
	})
}

func testAccSubgraphResourceConfig(name, namespace, baseSubgraphName, routingURL string) string {
	return fmt.Sprintf(`
resource "cosmo_subgraph" "test" {
  name                = "%s"
  namespace           = "%s"
  base_subgraph_name  = "%s"
  routing_url         = "%s"
  labels              = ["team=backend", "stage=dev"]
}
`, name, namespace, baseSubgraphName, routingURL)
}
