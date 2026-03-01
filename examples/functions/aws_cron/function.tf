# Copyright Tatsuya BIZENN <bizenn@gmail.com> 2024, 0
# SPDX-License-Identifier: Apache-2.0

output "sample" {
  value = provider::timeconv::aws_cron("0 12 * * ? *")
}
