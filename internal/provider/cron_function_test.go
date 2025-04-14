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

func TestAwsCronFunction(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `output "aws_cron" {
					value = provider::timeconv::aws_cron("0 12 * * ? *")
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("aws_cron", knownvalue.StringExact("cron(0 12 * * ? *)")),
				},
			},
			{
				Config: `output "aws_cron" {
					value = provider::timeconv::aws_cron("0 * ? * FRI *")
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("aws_cron", knownvalue.StringExact("cron(0 * ? * FRI *)")),
				},
			},
			{
				Config: `output "invalid_input" {
					value = provider::timeconv::aws_cron("0 12 * * * *")
				}`,
				ExpectError: regexp.MustCompile(`failed:`),
			},
		},
	})
}

func TestUnixCronFunction(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `output "unix_cron" {
					value = provider::timeconv::unix_cron("0 12 * * ? *")
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("unix_cron", knownvalue.StringExact("cron(0 12 * * *)")),
				},
			},
			{
				Config: `output "unix_cron" {
					value = provider::timeconv::unix_cron("0 * ? * FRI *")
				}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue("unix_cron", knownvalue.StringExact("cron(0 * * * FRI)")),
				},
			},
			{
				Config: `output "unix_cron" {
					value = provider::timeconv::unix_cron("0 12 * * * *")
				}`,
				ExpectError: regexp.MustCompile(`failed:`),
			},
		},
	})
}
