package bcc

// Subnet deinfe subnet of vpc
type Subnet struct {
	SubnetID    string `json:"subnetId"`
	Name        string `json:"name"`
	ZoneName    string `json:"zoneName"`
	Cidr        string `json:"cidr"`
	VpcID       string `json:"vpcId"`
	SubnetType  string `json:"subnetType"`
	Description string `json:"description"`
}

// CreateSubnetArgs define args create a subnet
type CreateSubnetArgs struct {
	Name        string `json:"name"`
	ZoneName    string `json:"zoneName"`
	Cidr        string `json:"cidr"`
	VpcID       string `json:"vpcId"`
	SubnetType  string `json:"subnetType,omitempty"`
	Description string `json:"description,omitempty"`
}

// CreateSubnetResponse define response of creating a subnet
type CreateSubnetResponse struct {
	SubnetID string `json:"subnetId"`
}

// CreateSubnet create a subnet
// https://cloud.baidu.com/doc/VPC/API.html#.E5.88.9B.E5.BB.BA.E5.AD.90.E7.BD.91
// func (c *Client) CreateSubnet(args *CreateSubnetArgs) (string, error) {

// }
