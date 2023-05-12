package typesense

import "github.com/hashicorp/terraform-plugin-framework/types"

// clusterModel maps Typesense cluster schema data.
type clusterModel struct {
	ID     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Status types.String `tfsdk:"status"`
}

func NewClient(key string) (*typesenseClient, error) {
	return &typesenseClient{}, nil
}

type typesenseClient struct{}

func (c *typesenseClient) GetCluster(id string) (*clusterModel, error) {
	return &clusterModel{
		ID:     types.StringValue(id),
		Name:   types.StringValue("foobar"),
		Status: types.StringValue("unknown"),
	}, nil
}

func (c *typesenseClient) CreateCluster(model clusterModel) (*clusterModel, error) {
	return nil, nil
}
