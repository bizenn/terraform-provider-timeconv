// Copyright (c) Tatsuya BIZENN <bizenn@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/winebarrel/cronplan"
)

type awsCron struct{}

// Definition implements function.Function.
func (a *awsCron) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Convert to AWS cron expression string(w/ \"cron(...)\")",
		Description: "Convert to AWS cron expression string(w/ \"cron(...)\"). See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-scheduled-rule-pattern.html#eb-cron-expressions",

		Parameters: []function.Parameter{
			function.StringParameter{
				Name:           "input",
				Description:    "Input time string in Amazon EventBridge cron format (w/o \"cron(...)\"). See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-scheduled-rule-pattern.html#eb-cron-expressions",
				AllowNullValue: false,
			},
		},
		Return: function.StringReturn{},
	}
}

// Metadata implements function.Function.
func (a *awsCron) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "aws_cron"
}

// Run implements function.Function.
func (a *awsCron) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	var output string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &input))

	expr, err := cronplan.Parse(input)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		return
	}
	output = fmt.Sprintf("cron(%s)", expr.String())
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, output))
}

var _ function.Function = (*awsCron)(nil)

func NewAwsCronFunction() function.Function {
	return &awsCron{}
}

type unixCron struct{}

// Definition implements function.Function.
func (u *unixCron) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Convert to Unix cron expression string(w/ \"cron(...)\")",
		Description: "Convert to Unix cron expression string(w/ \"cron(...)\").  See: https://crontab.guru/",

		Parameters: []function.Parameter{
			function.StringParameter{
				Name:           "input",
				Description:    "Input time string in Amazon EventBridge cron format (w/o \"cron(...)\"). See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-scheduled-rule-pattern.html#eb-cron-expressions",
				AllowNullValue: false,
			},
		},
		Return: function.StringReturn{},
	}
}

// Metadata implements function.Function.
func (u *unixCron) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "unix_cron"
}

// Run implements function.Function.
func (u *unixCron) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	var output string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &input))

	expr, err := cronplan.Parse(input)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		return
	}
	fs := []string{"*", "*", "*", "*", "*"}
	fs[0] = expr.Minute.String()
	fs[1] = expr.Hour.String()
	if expr.DayOfMonth.String() != "?" {
		fs[2] = expr.DayOfMonth.String()
	}
	fs[3] = expr.Month.String()
	if expr.DayOfWeek.String() != "?" {
		fs[4] = expr.DayOfWeek.String()
	}
	output = fmt.Sprintf("cron(%s)", strings.Join(fs, " "))
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, output))
}

var _ function.Function = (*unixCron)(nil)

func NewUnixCronFunction() function.Function {
	return &unixCron{}
}
