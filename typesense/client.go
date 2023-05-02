package typesense

import "github.com/hashicorp/terraform-plugin-framework/types"

func NewClient(key string) (*typesenseClient, error) {
	return &typesenseClient{}, nil
}

type typesenseClient struct{}

func (c *typesenseClient) GetCluster() (*clusterDataSourceModel, error) {
	return &clusterDataSourceModel{
		ID:     types.StringValue("fb"),
		Name:   types.StringValue("foobar"),
		Status: types.StringValue("unknown"),
	}, nil
}
