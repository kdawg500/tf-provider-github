package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type ghDataSource struct{}

func NewGithubDataSource() datasource.DataSource {
	return &ghDataSource{}
}

func (g *ghDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	panic("unimplemented")
}

func (g *ghDataSource) Read(ctx context.Context, res datasource.ReadRequest, resp *datasource.ReadResponse) {
	panic("unimplemented")
}

func (g *ghDataSource) Schema(ctx context.Context, res datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	panic("unimplemented")
}
