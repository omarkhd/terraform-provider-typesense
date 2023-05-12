package typesense

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &clusterDataSource{}
	_ datasource.DataSourceWithConfigure = &clusterDataSource{}

	clusterDataSourceSchema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"memory": schema.StringAttribute{
				Computed: true,
			},
			"vcpu": schema.StringAttribute{
				Computed: true,
			},
			"high_performance_disk": schema.StringAttribute{
				Computed: true,
			},
			"typesense_server_version": schema.StringAttribute{
				Computed: true,
			},
			"high_availability": schema.StringAttribute{
				Computed: true,
			},
			"search_delivery_network": schema.StringAttribute{
				Computed: true,
			},
			"load_balancing": schema.StringAttribute{
				Computed: true,
			},
			"region": schema.StringAttribute{
				Computed: true,
			},
			"auto_upgrade_capacity": schema.BoolAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Computed: true,
			},
		},
	}
)

func NewClusterDataSource() datasource.DataSource {
	return &clusterDataSource{}
}

type clusterDataSource struct {
	client *typesenseClient
}

func (cds *clusterDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cluster"
}

func (cds *clusterDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = clusterDataSourceSchema
}

func (cds *clusterDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// Get current state
	var config typesenseClusterModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	cluster, err := cds.client.GetCluster(config.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read Typesense cluster",
			err.Error(),
		)
		return
	}
	tcm := &typesenseClusterModel{
		ID:                     types.StringValue(cluster.ID),
		Name:                   types.StringValue(cluster.Name),
		Memory:                 types.StringValue(cluster.Memory),
		VCPU:                   types.StringValue(cluster.VCPU),
		HighPerformanceDisk:    types.StringValue(cluster.HighPerformanceDisk),
		TypesenseServerVersion: types.StringValue(cluster.TypesenseServerVersion),
		HighAvailability:       types.StringValue(cluster.HighAvailability),
		SearchDeliveryNetwork:  types.StringValue(cluster.SearchDeliveryNetwork),
		LoadBalancing:          types.StringValue(cluster.LoadBalancing),
		Region:                 types.StringValue(cluster.Regions[0]),
		AutoUpgradeCapacity:    types.BoolValue(cluster.AutoUpgradeCapacity),
		Status:                 types.StringValue(cluster.Status),
	}
	// Set state
	diags = resp.State.Set(ctx, tcm)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Configure adds the provider configured client to the data source.
func (cds *clusterDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	cds.client = req.ProviderData.(*typesenseClient)
}
