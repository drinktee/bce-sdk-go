package bcc

import (
	"encoding/json"

	"github.com/drinktee/bce-sdk-go/bce"
)

// Endpoint contains all endpoints of Baidu Cloud BCC.
var Endpoint = map[string]string{
	"bj": "bcc.bj.baidubce.com",
	"gz": "bcc.gz.baidubce.com",
}

// Config contains all options for bos.Client.
type Config struct {
	*bce.Config
}

func NewConfig(config *bce.Config) *Config {
	return &Config{config}
}

// Client is the bos client implemention for Baidu Cloud BOS API.
type Client struct {
	*bce.Client
}

func NewClient(config *Config) *Client {
	bceClient := bce.NewClient(config.Config)
	return &Client{bceClient}
}

// GetURL generates the full URL of http request for Baidu Cloud BOS API.
func (c *Client) GetURL(objectKey string, params map[string]string) string {
	host := c.Endpoint

	if host == "" {
		host = Endpoint[c.GetRegion()]
	}

	uriPath := objectKey

	return c.Client.GetURL(host, uriPath, params)
}

// ListInstances gets all Instances.
func (c *Client) ListInstances(option *bce.SignOption) ([]Instance, error) {

	req, err := bce.NewRequest("GET", c.GetURL("v2/instance", nil), nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.SendRequest(req, option)

	if err != nil {
		return nil, err
	}

	bodyContent, err := resp.GetBodyContent()

	if err != nil {
		return nil, err
	}

	var insList *ListInstancesResponse
	err = json.Unmarshal(bodyContent, &insList)

	if err != nil {
		return nil, err
	}

	return insList.Instances, nil
}

func (c *Client) GetInstance(instanceId string, option *bce.SignOption) (*Instance, error) {

	req, err := bce.NewRequest("GET", c.GetURL("v2/instance"+"/"+instanceId, nil), nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.SendRequest(req, option)

	if err != nil {
		return nil, err
	}

	bodyContent, err := resp.GetBodyContent()

	if err != nil {
		return nil, err
	}

	var ins *GetInstanceResponse
	err = json.Unmarshal(bodyContent, &ins)

	if err != nil {
		return nil, err
	}

	return &ins.Ins, nil
}
