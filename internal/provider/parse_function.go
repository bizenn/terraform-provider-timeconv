// Copyright (c) Tatsuya BIZENN <bizenn@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type parse struct{}

// Definition implements function.Function.
func (p *parse) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Parse a time string",
		Description: "Parse a time string using the specified format(golang time package style). If layout is null, RFC3339 format is used. See: https://pkg.go.dev/time#Parse",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:           "layout",
				Description:    "Layout string represents input time format(golang time package style)",
				AllowNullValue: true,
			},
			function.StringParameter{
				Name:           "input",
				Description:    "Input time string",
				AllowNullValue: false,
			},
		},
		Return: function.StringReturn{},
	}
}

// Metadata implements function.Function.
func (p *parse) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "parse"
}

// Run implements function.Function.
func (p *parse) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var layout types.String
	var input string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &layout, &input))
	layoutString := time.RFC3339
	if !layout.IsNull() {
		layoutString = layout.ValueString()
	}
	t, err := time.Parse(layoutString, input)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		return
	}
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, timetypes.NewRFC3339TimeValue(t)))
}

var _ function.Function = (*parse)(nil)

func NewParseFunction() function.Function {
	return &parse{}
}

type parseInLocation struct{}

// Definition implements function.Function.
func (p *parseInLocation) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Parse a time string",
		Description: "Parse a time string using the specified format(golang time package style). If layout is null, RFC3339 format is used. See: https://pkg.go.dev/time#ParseInLocation",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:           "layout",
				Description:    "Layout string represents input time format(golang time package style)",
				AllowNullValue: true,
			},
			function.StringParameter{
				Name:           "input",
				Description:    "Input time string",
				AllowNullValue: false,
			},
			function.StringParameter{
				Name:           "location",
				Description:    "Location string represents input time zone",
				AllowNullValue: false,
			},
		},
		Return: function.StringReturn{},
	}
}

// Metadata implements function.Function.
func (p *parseInLocation) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "parse_in_location"
}

// Run implements function.Function.
func (p *parseInLocation) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var layout types.String
	var input string
	var location string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &layout, &input, &location))
	layoutString := time.RFC3339
	if !layout.IsNull() {
		layoutString = layout.ValueString()
	}
	loc, err := time.LoadLocation(location)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		return
	}
	t, err := time.ParseInLocation(layoutString, input, loc)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		return
	}
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, timetypes.NewRFC3339TimeValue(t)))
}

var _ function.Function = (*parseInLocation)(nil)

func NewParseInLocationFunction() function.Function {
	return &parseInLocation{}
}
