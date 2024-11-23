package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccServiceDataSource(t *testing.T) {
	resName := "data.coolify_service.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `data "coolify_service" "test" {
					uuid = "` + testAccServiceUUID + `"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "uuid", testAccServiceUUID),
					resource.TestCheckResourceAttr(resName, "created_at", "2024-11-18T06:03:18.000000Z"),
					resource.TestCheckResourceAttrSet(resName, "docker_compose"),
					resource.TestCheckResourceAttr(resName, "name", "service-"+testAccServiceUUID),
					resource.TestCheckNoResourceAttr(resName, "description"),
				),
			},
		},
	})
}
