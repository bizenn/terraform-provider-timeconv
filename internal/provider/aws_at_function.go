// Copyright (c) Tatsuya BIZENN <bizenn@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/function"
)

type awsAt struct{}

// Definition implements function.Function.
func (a *awsAt) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Convert to AWS at expression string",
		Description: "Convert to AWS at expression string in input time's timezone.  See: https://docs.aws.amazon.com/ja_jp/autoscaling/application/APIReference/API_PutScheduledAction.html#API_PutScheduledAction_RequestSyntax",

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
func (a *awsAt) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "aws_at"
}

// Run implements function.Function.
func (a *awsAt) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input timetypes.RFC3339
	var output string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &input))

	t, _ := input.ValueRFC3339Time()
	output = fmt.Sprintf("at(%s)", t.Format("2006-01-02T15:04:05"))
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, output))
}

var _ function.Function = (*awsAt)(nil)

func NewAwsAtFunction() function.Function {
	return &awsAt{}
}
