// Generate copyright headers
//go:generate go get github.com/hashicorp/copywrite
//go:generate go run github.com/hashicorp/copywrite headers -d .. --config ../.copywrite.hcl

// Format Terraform code for use in documentation.
// If you do not have Terraform installed, you can remove the formatting command, but it is suggested
// to ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ../examples/

// Generate documentation.
//
//go:generate go get github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-dir .. -provider-name timeconv
//go:generate go mod tidy
package tools
