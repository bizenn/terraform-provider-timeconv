// Copyright (c) Tatsuya BIZENN <bizenn@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestTimeDataSource(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				data "timeconv_time" "example" {
					input = "2023-02-15T16:35:00+09:00"
					output_location = "America/Los_Angeles"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.timeconv_time.example", "output", "2023-02-14T23:35:00-08:00"),
					resource.TestCheckResourceAttr("data.timeconv_time.example", "aws_cron", "35 23 14 2 ? 2023"),
					resource.TestCheckResourceAttr("data.timeconv_time.example", "cron", "35 23 14 2 ? 2023"),
					resource.TestCheckResourceAttr("data.timeconv_time.example", "unix", "1676446500"),
					resource.TestCheckResourceAttr("data.timeconv_time.example", "at", "at(2023-02-15T07:35:00)"),
				),
			},
			{
				Config: `
				data "timeconv_time" "example" {
					input = "2023-02-15T16:35:00+09:00"
					output_location = "UTC"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.timeconv_time.example", "output", "2023-02-15T07:35:00Z"),
					resource.TestCheckResourceAttr("data.timeconv_time.example", "aws_cron", "35 7 15 2 ? 2023"),
					resource.TestCheckResourceAttr("data.timeconv_time.example", "cron", "35 7 15 2 ? 2023"),
					resource.TestCheckResourceAttr("data.timeconv_time.example", "unix", "1676446500"),
					resource.TestCheckResourceAttr("data.timeconv_time.example", "at", "at(2023-02-15T07:35:00)"),
				),
			},
			{
				Config: `
				data "timeconv_time" "example" {
					input = "2023-02-15 07:36:05"
					input_format = "2006-01-02 15:04:05"
					input_location = "Asia/Tokyo"
					output_format = "02 Jan 06 15:04 MST"
					output_location = "UTC"
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.timeconv_time.example", "output", "14 Feb 23 22:36 UTC"),
					resource.TestCheckResourceAttr("data.timeconv_time.example", "aws_cron", "36 22 14 2 ? 2023"),
					resource.TestCheckResourceAttr("data.timeconv_time.example", "cron", "36 22 14 2 ? 2023"),
					resource.TestCheckResourceAttr("data.timeconv_time.example", "unix", "1676414165"),
					resource.TestCheckResourceAttr("data.timeconv_time.example", "at", "at(2023-02-14T22:36:05)"),
				),
			},
		},
	})
}
