# Copyright (c) Tatsuya BIZENN <bizenn@gmail.com>
# SPDX-License-Identifier: Apache-2.0

terraform {
  required_providers {
    timeconv = {
      source = "github.com/bizenn/timeconv"
    }
  }
}

provider "timeconv" {}

locals {
  configurations = {
    "01" = {
      input           = "2023-02-15T07:36:00+09:00"
      input_format    = ""
      input_location  = ""
      output_format   = ""
      output_location = ""
    }
    "02" = {
      input           = "2023-02-15T07:36:00+09:00"
      input_format    = ""
      input_location  = ""
      output_format   = ""
      output_location = "UTC"
    }
    "03" = {
      input           = "2023-02-15 07:36:05"
      input_format    = "2006-01-02 15:04:05"
      input_location  = "Asia/Tokyo"
      output_format   = "02 Jan 06 15:04 MST"
      output_location = "UTC"
    }
  }
}

data "timeconv_time" "example" {
  for_each = local.configurations

  input           = each.value.input
  input_format    = each.value.input_format
  input_location  = each.value.input_location
  output_format   = each.value.output_format
  output_location = each.value.output_location
}

output "timeconv_time" {
  value = data.timeconv_time.example
}