package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	TIME_DS = "time"
)

var (
	_ datasource.DataSource              = &timeDataSource{}
	_ datasource.DataSourceWithConfigure = &timeDataSource{}
)

func NewTimeDataSource() datasource.DataSource {
	return &timeDataSource{}
}

type timeDataSource struct{}

func (d *timeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, res *datasource.MetadataResponse) {
	res.TypeName = req.ProviderTypeName + "_" + TIME_DS
}

func (d *timeDataSource) Schema(_ context.Context, req datasource.SchemaRequest, res *datasource.SchemaResponse) {
	res.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"input": schema.StringAttribute{
				Optional:    true,
				Description: "Input time string",
			},
			"input_format": schema.StringAttribute{
				Optional:    true,
				Description: "Input time format. Default is RFC3339(\"2006-01-02T15:04:05Z07:00\").",
			},
			"input_location": schema.StringAttribute{
				Optional:    true,
				Description: "Input timezone location. Default is the system localtime.",
			},
			"output": schema.StringAttribute{
				Computed:    true,
				Description: "Output time string",
			},
			"output_format": schema.StringAttribute{
				Optional:    true,
				Description: "Output time format. Default is RFC3339(\"2006-01-02T15:04:05Z07:00\")",
			},
			"output_location": schema.StringAttribute{
				Optional:    true,
				Description: "Output timezone location. Default is the system localtime.",
			},
			"aws_cron": schema.StringAttribute{
				Computed:    true,
				Description: "AWS cron expression in output location.",
			},
			"cron": schema.StringAttribute{
				Computed:           true,
				DeprecationMessage: "You should use aws_cron instead of this",
				Description:        "AWS cron expression in output location.",
			},
			"unix": schema.Int64Attribute{
				Computed:    true,
				Description: "Unix time in seconds",
			},
		},
	}
}

func (d *timeDataSource) Read(ctx context.Context, req datasource.ReadRequest, res *datasource.ReadResponse) {
	var config timeDataSourceModel

	diags := req.Config.Get(ctx, &config)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}

	var err error
	t := time.Now()
	loc := time.Local

	inputFormat := config.InputFormat.ValueString()
	if inputFormat == "" {
		inputFormat = time.RFC3339
	}

	inputLocation := config.InputLocation.ValueString()
	if inputLocation != "" {
		if loc, err = time.LoadLocation(inputLocation); err != nil {
			res.Diagnostics.AddError(
				"Input location loading error",
				"Cannot load the input_location.\n\n"+
					fmt.Sprintf("Error: %s", err),
			)
			return
		}
	}

	input := config.Input.ValueString()
	if input != "" {
		if t, err = time.ParseInLocation(inputFormat, input, loc); err != nil {
			res.Diagnostics.AddError(
				"Input time string parsing error",
				"Cannot parse the input time string.\n\n"+
					fmt.Sprintf("Error: %s", err),
			)
			return
		}
	}

	outputFormat := config.OutputFormat.ValueString()
	if outputFormat == "" {
		outputFormat = time.RFC3339
	}

	outloc := time.Local
	outputLocation := config.OutputLocation.ValueString()
	if outputLocation != "" {
		if outloc, err = time.LoadLocation(outputLocation); err != nil {
			res.Diagnostics.AddError(
				"Output location loading error",
				"Cannot load the output_location.\n\n"+
					fmt.Sprintf("Error: %s", err),
			)
		}
	}

	out := t.In(outloc)

	state := timeDataSourceModel{
		input:          t,
		output:         out,
		Input:          config.Input,
		InputFormat:    types.StringValue(inputFormat),
		InputLocation:  types.StringValue(inputLocation),
		Output:         types.StringValue(out.Format(outputFormat)),
		OutputFormat:   types.StringValue(outputFormat),
		OutputLocation: types.StringValue(outputLocation),
		AwsCron:        types.StringValue(cron(out)),
		Cron:           types.StringValue(cron(out)),
		Unix:           types.Int64Value(out.Unix()),
	}
	diags = res.State.Set(ctx, state)
	res.Diagnostics.Append(diags...)
}

func (d *timeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, res *datasource.ConfigureResponse) {
}

type timeDataSourceModel struct {
	input          time.Time
	output         time.Time
	Input          types.String `tfsdk:"input"`
	InputFormat    types.String `tfsdk:"input_format"`
	InputLocation  types.String `tfsdk:"input_location"`
	Output         types.String `tfsdk:"output"`
	OutputFormat   types.String `tfsdk:"output_format"`
	OutputLocation types.String `tfsdk:"output_location"`
	AwsCron        types.String `tfsdk:"aws_cron"`
	Cron           types.String `tfsdk:"cron"`
	Unix           types.Int64  `tfsdk:"unix"`
}

func cron(t time.Time) string {
	return fmt.Sprintf("%d %d %d %d ? %d", t.Minute(), t.Hour(), t.Day(), t.Month(), t.Year())
}
