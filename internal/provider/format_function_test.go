package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestFormatFonction(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "formatted_time" {
					value = provider::timeconv::format("2024-08-31T01:23:45Z", "2006-01-02 15:04:05")
				}
				output "formatted_time_with_ANSIC" {
					value = provider::timeconv::format("2024-08-07T01:23:45Z", "Mon Jan _2 15:04:05 2006")
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("formatted_time", knownvalue.StringExact("2024-08-31 01:23:45")),
					statecheck.ExpectKnownOutputValue("formatted_time_with_ANSIC", knownvalue.StringExact("Wed Aug  7 01:23:45 2024")),
				},
			},
			{
				Config: `output "invalid_input" {
					value = provider::timeconv::format("2024-08-31T00:00:00", "UTC")
				}`,
				ExpectError: regexp.MustCompile(`Invalid function argument`),
			},
		},
	})
}
