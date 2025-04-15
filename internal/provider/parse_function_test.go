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

func TestParseFunction(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "parsed_without_layout" {
					value = provider::timeconv::parse(null, "2024-08-31T01:23:45+09:00")
				}
				output "parsed_time" {
					value = provider::timeconv::parse("2006-01-02 15:04:05 MST", "2024-08-31 01:23:45 JST")
				}
				output "parsed_time_with_ANSIC" {
					value = provider::timeconv::parse("Mon Jan _2 15:04:05 2006", "Wed Aug  7 01:23:45 2024")
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("parsed_without_layout", knownvalue.StringExact("2024-08-31T01:23:45+09:00")),
					statecheck.ExpectKnownOutputValue("parsed_time", knownvalue.StringExact("2024-08-31T01:23:45+09:00")),
					statecheck.ExpectKnownOutputValue("parsed_time_with_ANSIC", knownvalue.StringExact("2024-08-07T01:23:45Z")),
				},
			},
			{
				Config: `output "invalid_input" {
					value = provider::timeconv::parse("2006-01-02 15:04:05", "2024-08-31T00:00:00")
				}`,
				ExpectError: regexp.MustCompile(`failed: parsing time`),
			},
		},
	})
}

func TestParseInLocationFunction(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "parsed_without_layout" {
					value = provider::timeconv::parse_in_location(null, "2024-08-31T01:23:45Z", "UTC")
				}
				output "parsed_time" {
					value = provider::timeconv::parse_in_location("2006-01-02 15:04:05 MST", "2024-08-31 01:23:45 JST", "Asia/Tokyo")
				}
				output "parsed_time_with_ANSIC" {
					value = provider::timeconv::parse_in_location("Mon Jan _2 15:04:05 2006", "Wed Aug  7 01:23:45 2024", "Asia/Tokyo")
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("parsed_without_layout", knownvalue.StringExact("2024-08-31T01:23:45Z")),
					statecheck.ExpectKnownOutputValue("parsed_time", knownvalue.StringExact("2024-08-31T01:23:45+09:00")),
					statecheck.ExpectKnownOutputValue("parsed_time_with_ANSIC", knownvalue.StringExact("2024-08-07T01:23:45+09:00")),
				},
			},
			{
				Config: `output "invalid_input" {
					value = provider::timeconv::parse_in_location("2006-01-02 15:04:05", "2024-08-31T00:00:00", "Asia/Tokyo")
				}`,
				ExpectError: regexp.MustCompile(`failed: parsing time`),
			},
			{
				Config: `output "invalid_location" {
					value = provider::timeconv::parse_in_location(null, "2024-08-31T01:23:45+09:00", "invalid/location")
				}`,
				ExpectError: regexp.MustCompile(`failed: unknown time`),
			},
		},
	})
}
