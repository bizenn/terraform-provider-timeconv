# Copyright (c) Tatsuya BIZENN <bizenn@gmail.com>
# SPDX-License-Identifier: Apache-2.0

output "sample" {
  value = provider::timeconv::format("2024-08-31T01:23:45+0900", "YYYY-MM-DD HH:mm:ss")
}
