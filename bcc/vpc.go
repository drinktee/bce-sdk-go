package bcc

// Vpc type define
type Vpc struct {
	VpcID       string `json:"vpcId"`
	Name        string `json:"name"`
	CIDR        string `json:"cidr"`
	Description string `json:"description"`
	IsDefault   bool   `json:"isDefault"`
}

// ShowVpc define ShowVpcModel
type ShowVpc struct {
	VpcID       string   `json:"vpcId"`
	Name        string   `json:"name"`
	CIDR        string   `json:"cidr"`
	Description string   `json:"description"`
	IsDefault   bool     `json:"isDefault"`
	Subnets     []Subnet `json:"subnets"`
}

// CreateVpcArgs define args for creating vpc
type CreateVpcArgs struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Cidr        string `json:"cidr"`
}

// CreateVpcResponse define response
type CreateVpcResponse struct {
	VpcID string `json:"vpcId"`
}

// CreateVpc create a vpc
// https://cloud.baidu.com/doc/VPC/API.html#.E5.88.9B.E5.BB.BAVPC
// func (c *Client) CreateVpc(args *CreateVpcArgs) (string, error) {

// }

// ListVpcArgs args
type ListVpcArgs struct {
	IsDefault bool `json:"isDefault"`
}

// ListVpc list all vpcs
// https://cloud.baidu.com/doc/VPC/API.html#.E6.9F.A5.E8.AF.A2VPC.E5.88.97.E8.A1.A8
// func (c *Client) ListVpc(args *ListVpcArgs) ([]Vpc, error) {

// }

// DeleteVpc delete a vpc by id
// https://cloud.baidu.com/doc/VPC/API.html#.E5.88.A0.E9.99.A4VPC
// func (c *Client) DeleteVpc(vpcid string) error {

// }

// TODO: UpdateVpc
