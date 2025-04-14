# Copyright (c) Tatsuya BIZENN <bizenn@gmail.com>
# SPDX-License-Identifier: Apache-2.0

output "sample" {
  value = provider::timeconv::unix_cron("0 12 * * ? *")
}
