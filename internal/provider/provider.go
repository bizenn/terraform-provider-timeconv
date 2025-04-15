// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure ScaffoldingProvider satisfies various provider interfaces.
var _ provider.ProviderWithFunctions = &TimeconvProvider{}

// TimeconvProvider defines the provider implementation.
type TimeconvProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// TimeconvProviderModel describes the provider data model.
type TimeconvProviderModel struct {
}

func (p *TimeconvProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "timeconv"
	resp.Version = p.version
}

func (p *TimeconvProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{},
	}
}

func (p *TimeconvProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *TimeconvProvider) Resources(ctx context.Context) []func() resource.Resource {
	return nil
}

func (p *TimeconvProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewTimeDataSource,
	}
}

func (p *TimeconvProvider) Functions(_ context.Context) []func() function.Function {
	return []func() function.Function{
		NewTimezoneFunction,
		NewFormatFunction,
		NewAwsAtFunction,
		NewZoneNameFunction,
		NewZoneOffsetFunction,
		NewAwsCronFunction,
		NewUnixCronFunction,
		NewParseFunction,
		NewParseInLocationFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &TimeconvProvider{
			version: version,
		}
	}
}
