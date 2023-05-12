package typesense

import "github.com/hashicorp/terraform-plugin-framework/types"

func NewClient(key string) (*typesenseClient, error) {
	return &typesenseClient{}, nil
}

type typesenseClient struct{}

func (c *typesenseClient) GetCluster(id string) (*clusterDataSourceModel, error) {
	return &clusterDataSourceModel{
		ID:     types.StringValue(id),
		Name:   types.StringValue("foobar"),
		Status: types.StringValue("unknown"),
	}, nil
}

func (c *typesenseClient) CreateCluster(model clusterDataSourceModel) (*clusterDataSourceModel, error) {
	return nil, nil
}
