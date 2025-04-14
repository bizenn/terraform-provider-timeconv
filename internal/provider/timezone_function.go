// Copyright (c) Tatsuya BIZENN <bizenn@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/function"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &timezone{}

func NewTimezoneFunction() function.Function {
	return &timezone{}
}

type timezone struct{}

func (f *timezone) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "timezone"
}

func (f *timezone) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Convert timezone",
		Description: "Converts a given date time string from one timezone to another, returning the result in the specified format",

		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "Input RFC3339 time string",
				CustomType:  timetypes.RFC3339Type{},
			},
			function.StringParameter{
				Name:        "output_location",
				Description: "output timezone location. like `UTC`, `America/New_York`, `Asia/Tokyo`",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *timezone) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input timetypes.RFC3339
	var outputLocation string

	// Read Terraform argument data into the variable
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &input, &outputLocation))

	// Convert
	var err error
	var oLoc *time.Location
	if oLoc, err = time.LoadLocation(outputLocation); err != nil {
		resp.Error = function.NewFuncError(
			fmt.Sprintf("Output location loading error: %s", err),
		)
		return
	}

	t, _ := input.ValueRFC3339Time()
	output := t.In(oLoc).Format(time.RFC3339)

	// Set the result
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, output))
}
