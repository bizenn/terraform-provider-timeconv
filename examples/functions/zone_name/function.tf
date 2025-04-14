# Copyright (c) Tatsuya BIZENN <bizenn@gmail.com>
# SPDX-License-Identifier: Apache-2.0

output "testUTC" {
  value = provider::zone_name::zone_name("2024-08-31T01:23:45Z")
}

output "testJST" {
  value = provider::zone_name::zone_name("2024-08-31T01:23:45+09:00")
}
