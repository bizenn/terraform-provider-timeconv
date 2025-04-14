// Copyright (c) Tatsuya BIZENN <bizenn@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/function"
)

type format struct{}

// Definition implements function.Function.
func (f *format) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Format a time string",
		Description: "Format a time string using the specified format(golang time package style).",

		Parameters: []function.Parameter{
			function.StringParameter{
				Name:           "input",
				Description:    "Input time string in RFC3339 format",
				CustomType:     timetypes.RFC3339Type{},
				AllowNullValue: false,
			},
			function.StringParameter{
				Name:           "output_format",
				Description:    "Output time format(golang time package style)",
				AllowNullValue: false,
			},
		},
		Return: function.StringReturn{},
	}
}

// Metadata implements function.Function.
func (f *format) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "format"
}

// Run implements function.Function.
func (f *format) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input timetypes.RFC3339
	var outputFormat string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &input, &outputFormat))

	t, _ := input.ValueRFC3339Time()
	output := t.Format(outputFormat)
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, output))
}

var _ function.Function = (*format)(nil)

func NewFormatFunction() function.Function {
	return &format{}
}
