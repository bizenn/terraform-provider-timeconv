// Copyright (c) Tatsuya BIZENN <bizenn@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/function"
)

type zoneName struct{}

// Definition implements function.Function.
func (z *zoneName) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Return timezone name",
		Description: "Return timezone name in input time's timezone.  See: https://pkg.go.dev/time#Time.Zone",

		Parameters: []function.Parameter{
			function.StringParameter{
				Name:           "input",
				Description:    "Input time string in RFC3339 format",
				CustomType:     timetypes.RFC3339Type{},
				AllowNullValue: false,
			},
		},
		Return: function.StringReturn{},
	}
}

// Metadata implements function.Function.
func (z *zoneName) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "zone_name"
}

// Run implements function.Function.
func (z *zoneName) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input timetypes.RFC3339
	var output string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &input))

	t, _ := input.ValueRFC3339Time()
	output, _ = t.Zone()
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, output))
}

var _ function.Function = (*zoneName)(nil)

func NewZoneNameFunction() function.Function {
	return &zoneName{}
}

type zoneOffset struct{}

// Definition implements function.Function.
func (z *zoneOffset) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Return timezone offset in seconds from east of UTC",
		Description: "Return timezone offset in seconds from east of UTC.  See: https://pkg.go.dev/time#Time.Zone",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:           "input",
				Description:    "Input time string in RFC3339 format",
				CustomType:     timetypes.RFC3339Type{},
				AllowNullValue: false,
			},
		},
		Return: function.NumberReturn{},
	}
}

// Metadata implements function.Function.
func (z *zoneOffset) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "zone_offset"
}

// Run implements function.Function.
func (z *zoneOffset) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input timetypes.RFC3339
	var output int

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &input))

	t, _ := input.ValueRFC3339Time()
	_, output = t.Zone()
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, output))
}

var _ function.Function = (*zoneOffset)(nil)

func NewZoneOffsetFunction() function.Function {
	return &zoneOffset{}
}
