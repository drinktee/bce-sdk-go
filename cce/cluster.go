package cce

import (
	"encoding/json"

	"fmt"

	"github.com/drinktee/bce-sdk-go/bcc"
	"github.com/drinktee/bce-sdk-go/bce"
)

// ListInstances gets all Instances of a cluster.
func (c *Client) ListInstances(clusterID string) ([]bcc.Instance, error) {
	if clusterID == "" {
		return nil, fmt.Errorf("clusterID should not be nil")
	}
	params := map[string]string{
		"clusterid": clusterID,
	}
	req, err := bce.NewRequest("GET", c.GetURL("/api/service/cce/instance", params), nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.SendRequest(req, nil)

	if err != nil {
		return nil, err
	}

	bodyContent, err := resp.GetBodyContent()

	if err != nil {
		return nil, err
	}

	var insList *bcc.ListInstancesResponse
	err = json.Unmarshal(bodyContent, &insList)

	if err != nil {
		return nil, err
	}

	return insList.Instances, nil
}
