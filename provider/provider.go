package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v70/github"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	typeName    = "github"
	tokenEnvKey = "GITHUB_TOKEN"
)

type ghProvider struct {
	version string
}

type providerModel struct {
	Token types.String `tfsdk:"token"`
}

func (gp *ghProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config providerModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {

		return
	}

	if config.Token.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("token"),
			"Unknown github token",
			fmt.Sprintf("The provider cannot create the github API client as there is an unknown configuration value for the github access token. Either target apply the source of the value first, set the value statically in the configuration, or use the %v environment variable.", tokenEnvKey),
		)
	}

	token := os.Getenv(tokenEnvKey)
	// override env var with config
	if !config.Token.IsNull() {
		token = config.Token.ValueString()
	}

	if token == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("token"),
			"Missing github access token",
			"The provider cannot create the github client as there is a missing or empty value for the github token. "+
				fmt.Sprintf("Set the host value in the configuration or use the %v environment variable. ", tokenEnvKey)+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	client := github.NewClient(nil).WithAuthToken(token)
	resp.DataSourceData = client
	resp.ResourceData = client

}

func (gp *ghProvider) DataSources(context.Context) []func() datasource.DataSource {
	panic("unimplemented")
}

func (gp *ghProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = typeName
	resp.Version = gp.version
}

func (gp *ghProvider) Resources(context.Context) []func() resource.Resource {
	panic("unimplemented")
}

func (g *ghProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"token": schema.StringAttribute{
				Required:    true,
				Description: "Organization or personal github access token",
			},
		},
	}
}

func New() provider.Provider {
	return &ghProvider{}
}
