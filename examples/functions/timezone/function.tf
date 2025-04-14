# Copyright (c) Tatsuya BIZENN <bizenn@gmail.com>
# SPDX-License-Identifier: Apache-2.0

output "test" {
  value = provider::timeconv::timezone("2024-08-31T01:23:45+0900", "UTC")
}
