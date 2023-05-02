package typesense

func NewClient(key string) (*typesenseClient, error) {
	return &typesenseClient{}, nil
}

type typesenseClient struct{}
