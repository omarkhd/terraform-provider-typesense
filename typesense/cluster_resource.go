package typesense

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &clusterResource{}
	_ resource.ResourceWithConfigure = &clusterResource{}

	clusterResourceSchema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"memory": schema.StringAttribute{
				Required: true,
			},
			"vcpu": schema.StringAttribute{
				Required: true,
			},
			"high_performance_disk": schema.StringAttribute{
				Computed: true,
				Default:  stringdefault.StaticString("no"),
				Optional: true,
			},
			"typesense_server_version": schema.StringAttribute{
				Computed: true,
			},
			"high_availability": schema.StringAttribute{
				Computed: true,
				Default:  stringdefault.StaticString("no"),
				Optional: true,
			},
			"search_delivery_network": schema.StringAttribute{
				Computed: true,
			},
			"load_balancing": schema.StringAttribute{
				Computed: true,
			},
			"region": schema.StringAttribute{
				Required: true,
			},
			"auto_upgrade_capacity": schema.BoolAttribute{
				Computed: true,
				Default:  booldefault.StaticBool(false),
				Optional: true,
			},
			"status": schema.StringAttribute{
				Computed: true,
			},
		},
	}
)

// NewClusterResource is a helper function to simplify the provider implementation.
func NewClusterResource() resource.Resource {
	return &clusterResource{}
}

// clusterResource is the resource implementation.
type clusterResource struct {
	client *typesenseClient
}

// Configure adds the provider configured client to the resource.
func (cr *clusterResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	cr.client = req.ProviderData.(*typesenseClient)
}

// Metadata returns the resource type name.
func (cr *clusterResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cluster"
}

// Schema defines the schema for the resource.
func (cr *clusterResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = clusterResourceSchema
}

// Create creates the resource and sets the initial Terraform state.
func (cr *clusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan typesenseClusterModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create new cluster
	cluster, err := cr.client.CreateCluster(typesenseCluster{
		Memory:                plan.Memory.ValueString(),
		VCPU:                  plan.VCPU.ValueString(),
		Regions:               []string{plan.Region.ValueString()},
		HighAvailability:      plan.HighAvailability.ValueString(),
		SearchDeliveryNetwork: plan.SearchDeliveryNetwork.ValueString(),
		HighPerformanceDisk:   plan.HighPerformanceDisk.ValueString(),
		Name:                  plan.Name.ValueString(),
		AutoUpgradeCapacity:   plan.AutoUpgradeCapacity.ValueBool(),
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating cluster",
			"Could not create cluster, unexpected error: "+err.Error(),
		)
		return
	}
	plan.ID = types.StringValue(cluster.ID)
	plan.Name = types.StringValue(cluster.Name)
	plan.Memory = types.StringValue(cluster.Memory)
	plan.VCPU = types.StringValue(cluster.VCPU)
	plan.HighPerformanceDisk = types.StringValue(cluster.HighPerformanceDisk)
	plan.TypesenseServerVersion = types.StringValue(cluster.TypesenseServerVersion)
	plan.HighAvailability = types.StringValue(cluster.HighAvailability)
	plan.SearchDeliveryNetwork = types.StringValue(cluster.SearchDeliveryNetwork)
	plan.LoadBalancing = types.StringValue(cluster.LoadBalancing)
	plan.Region = types.StringValue(cluster.Regions[0])
	plan.AutoUpgradeCapacity = types.BoolValue(cluster.AutoUpgradeCapacity)
	plan.Status = types.StringValue(cluster.Status)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (cr *clusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state typesenseClusterModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed cluster value from Typesense
	cluster, err := cr.client.GetCluster(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Typesense Cluster",
			"Could not read Typesense Cluster ID "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}
	state.ID = types.StringValue(cluster.ID)
	state.Name = types.StringValue(cluster.Name)
	state.Memory = types.StringValue(cluster.Memory)
	state.VCPU = types.StringValue(cluster.VCPU)
	state.HighPerformanceDisk = types.StringValue(cluster.HighPerformanceDisk)
	state.TypesenseServerVersion = types.StringValue(cluster.TypesenseServerVersion)
	state.HighAvailability = types.StringValue(cluster.HighAvailability)
	state.SearchDeliveryNetwork = types.StringValue(cluster.SearchDeliveryNetwork)
	state.LoadBalancing = types.StringValue(cluster.LoadBalancing)
	state.Region = types.StringValue(cluster.Regions[0])
	state.AutoUpgradeCapacity = types.BoolValue(cluster.AutoUpgradeCapacity)
	state.Status = types.StringValue(cluster.Status)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (cr *clusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (cr *clusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
