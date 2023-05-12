package typesense

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const clusterEndpoint = "https://cloud.typesense.org/api/v1/clusters"

type typesenseCluster struct {
	ID                     string   `json:"id"`
	Name                   string   `json:"name"`
	Memory                 string   `json:"memory"`
	VCPU                   string   `json:"vcpu"`
	HighPerformanceDisk    string   `json:"high_performance_disk"`
	TypesenseServerVersion string   `json:"typesense_server_version"`
	HighAvailability       string   `json:"high_availability"`
	SearchDeliveryNetwork  string   `json:"search_delivery_network"`
	LoadBalancing          string   `json:"load_balancing"`
	Regions                []string `json:"regions"`
	AutoUpgradeCapacity    bool     `json:"auto_upgrade_capacity"`
	Status                 string   `json:"status"`
}

type typesenseClusterCreateResponse struct {
	Success bool             `json:"success"`
	Cluster typesenseCluster `json:"cluster"`
}

func NewClient(key string) (*typesenseClient, error) {
	return &typesenseClient{key: key}, nil
}

type typesenseClient struct {
	key string
}

func (c *typesenseClient) GetCluster(id string) (*typesenseCluster, error) {
	req, err := http.NewRequest("GET", clusterEndpoint+"/"+id, nil)
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	hc := http.Client{}
	resp, err := hc.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	var cluster typesenseCluster
	if err = json.Unmarshal(body, &cluster); err != nil {
		return nil, err
	}
	return &cluster, nil

}

func (c *typesenseClient) CreateCluster(model typesenseCluster) (*typesenseCluster, error) {
	params := map[string]interface{}{
		"memory":                  model.Memory,
		"vcpu":                    model.VCPU,
		"regions":                 model.Regions,
		"high_availability":       model.HighAvailability,
		"search_delivery_network": "off",
		"high_performance_disk":   model.HighPerformanceDisk,
		"name":                    model.Name,
		"auto_upgrade_capacity":   model.AutoUpgradeCapacity,
	}
	payload, _ := json.Marshal(params)
	req, err := http.NewRequest("POST", clusterEndpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	c.setHeaders(req)
	hc := http.Client{}
	resp, err := hc.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	var response typesenseClusterCreateResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	if !response.Success {
		return nil, errors.New(string(body))
	}
	return &response.Cluster, nil
}

func (c *typesenseClient) setHeaders(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-TYPESENSE-CLOUD-MANAGEMENT-API-KEY", c.key)
}
