package blb

import (
	"encoding/json"

	"bytes"

	"github.com/drinktee/bce-sdk-go/bce"
)

type LoadBalancer struct {
	BlbId    string `json:"blbId"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Address  string `json:"address"`
	Status   string `json:"status"`
	PublicIp string `json:"publicIp"`
}

type DescribeLoadBalancersArgs struct {
	LoadBalancerId   string
	LoadBalancerName string
	BCCId            string
	Address          string
}

type DescribeLoadBalancersResponse struct {
	Marker      string         `json:"marker"`
	IsTruncated bool           `json:"isTruncated"`
	NextMarker  string         `json:"nextMarker"`
	MaxKeys     int            `json:"maxKeys"`
	BLBList     []LoadBalancer `json:"blbList"`
}

type CreateLoadBalancerArgs struct {
	Desc string `json:"desc"`
	Name string `json:"name"`
	// ClientToken string
}

type CreateLoadBalancerResponse struct {
	LoadBalancerId string `json:"blbId"`
	Address        string `json:"address"`
	Desc           string `json:"desc"`
	Name           string `json:"name"`
}

// DescribeLoadBalancers Describe loadbalancers
func (c *Client) DescribeLoadBalancers(args *DescribeLoadBalancersArgs) ([]LoadBalancer, error) {
	var params map[string]string
	if args != nil {
		params = map[string]string{
			"blbId":   args.LoadBalancerId,
			"name":    args.LoadBalancerName,
			"bccId":   args.BCCId,
			"address": args.Address,
		}
	}
	req, err := bce.NewRequest("GET", c.GetURL("v1/blb", params), nil)

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
	var blbsResp *DescribeLoadBalancersResponse
	err = json.Unmarshal(bodyContent, &blbsResp)

	if err != nil {
		return nil, err
	}
	return blbsResp.BLBList, nil
}

func (c *Client) CreateLoadBalancer(args *CreateLoadBalancerArgs) (*CreateLoadBalancerResponse, error) {
	var params map[string]string
	if args != nil {
		params = map[string]string{
			"clientToken": c.GenerateClientToken(),
		}
	}
	postContent, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	req, err := bce.NewRequest("POST", c.GetURL("v1/blb", params), bytes.NewReader(postContent))
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
	var blbsResp *CreateLoadBalancerResponse
	err = json.Unmarshal(bodyContent, &blbsResp)

	if err != nil {
		return nil, err
	}
	return blbsResp, nil
}
