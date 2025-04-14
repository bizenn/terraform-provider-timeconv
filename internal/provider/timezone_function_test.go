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

func TestTimezoneFunction(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "jst_to_utc" {
                    value = provider::timeconv::timezone("2024-08-31T01:23:45+09:00", "UTC")
                }
				output "utc_to_jst" {
                    value = provider::timeconv::timezone("2024-08-31T01:23:45Z", "Asia/Tokyo")
                }`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("jst_to_utc", knownvalue.StringExact("2024-08-30T16:23:45Z")),
					statecheck.ExpectKnownOutputValue("utc_to_jst", knownvalue.StringExact("2024-08-31T10:23:45+09:00")),
				},
			},
			{
				Config: `
                output "invalid_input" {
                    value = provider::timeconv::timezone("2024-08-31T00:00:00", "UTC")
                }`,
				ExpectError: regexp.MustCompile(`Invalid function argument`),
			},
		},
	})
}
