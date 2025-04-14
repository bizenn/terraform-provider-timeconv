// Copyright (c) Tatsuya BIZENN <bizenn@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestZoneNameFunction(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `output "zone_name" {
					value = provider::timeconv::zone_name("2024-08-31T01:23:45Z")
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("zone_name", knownvalue.StringExact("UTC")),
				},
			},
			{
				Config: `output "invalid_input" {
					value = provider::timeconv::zone_name("2024-08-31T00:00:00")
				}`,
				ExpectError: regexp.MustCompile(`Invalid function argument`),
			},
		},
	})
}

func TestZoneOffsetFunction(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `output "zone_offset" {
					value = provider::timeconv::zone_offset("2024-08-31T01:23:45Z")
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("zone_offset", knownvalue.Int64Exact(0)),
				},
			},
			{
				Config: `output "zone_offset" {
					value = provider::timeconv::zone_offset("2024-08-31T01:23:45+09:00")
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("zone_offset", knownvalue.Int64Exact(32400)),
				},
			},
			{
				Config: `output "invalid_input" {
					value = provider::timeconv::zone_offset("2024-08-31T00:00:00")
				}`,
				ExpectError: regexp.MustCompile(`Invalid function argument`),
			},
		},
	})
}
