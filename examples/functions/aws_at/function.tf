# Copyright Tatsuya BIZENN <bizenn@gmail.com> 2024, 0
# SPDX-License-Identifier: Apache-2.0

output "sample" {
  value = provider::aws_at::function("2024-08-31T01:23:45Z")
}
