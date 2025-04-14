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

func TestAwsAtFunction(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `output "aws_at" {
					value = provider::timeconv::aws_at("2024-08-31T01:23:45Z")
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("aws_at", knownvalue.StringExact("at(2024-08-31T01:23:45)")),
				},
			},
			{
				Config: `output "invalid_input" {
					value = provider::timeconv::aws_at("2024-08-31T00:00:00")
				}`,
				ExpectError: regexp.MustCompile(`Invalid function argument`),
			},
		},
	})
}
