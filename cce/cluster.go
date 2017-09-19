package cce

import (
	"bytes"
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
	req, err := bce.NewRequest("GET", c.GetURL("/v1/instance", params), nil)

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

// ScaleUpClusterArgs define  args
type ScaleUpClusterArgs struct {
	ClusterID       string              `json:"clusterUuid"`
	CdsPreMountInfo bcc.CdsPreMountInfo `json:"cdsPreMountInfo"`
	OrderContent    OrderContent        `json:"orderContent"`
}

// ScaleUpClusterResponse define  args
type ScaleUpClusterResponse struct {
	ClusterID string   `json:"clusterUuid"`
	OrderID   []string `json:"orderId"`
}

// OrderContent define  bcc order content
type OrderContent struct {
	PaymentMethod []string    `json:"paymentMethod"`
	Items         []OrderItem `json:"items"`
}

// OrderItem define  bcc order content item
type OrderItem struct {
	Config        interface{} `json:"config"`
	PaymentMethod []string    `json:"paymentMethod"`
}

// BccOrderConfig define BCC order config
type BccOrderConfig struct {
	// 付费类型，一期只支持postpay
	ProductType string `json:"productType"`
	Region      string `json:"region"`
	LogicalZone string `json:"logicalZone"`
	// 普通BCC
	InstanceType string `json:"instanceType"`
	// 这些参数默认就行 容器产品用不到
	FpgaCard string `json:"fpgaCard"`
	GpuCard  int    `json:"gpuCard"`
	GpuCount int    `json:"gpuCount"`

	CPU    int `json:"cpu"`
	Memory int `json:"memory"`
	// 就一个镜像 ubuntu1604
	ImageType string `json:"imageType"`
	// 系统类型
	OsType string `json:"osType"`
	// 系统版本
	OsVersion string `json:"osVersion"`
	// 系统盘大小
	DiskSize int `json:"diskSize"`
	// 暂时为空
	EbsSize []int `json:"ebsSize"`
	// 是否需要购买EIP
	IfBuyEip bool `json:"ifBuyEip"`
	// eip名称
	EipName        string `json:"eipName"`
	SubProductType string `json:"subProductType"`
	// eip带宽
	BandwidthInMbps int `json:"bandwidthInMbps"`

	SubnetUuiD      string `json:"subnetUuid"`      // 子网uuid
	SecurityGroupID string `json:"securityGroupId"` // 安全组id

	AdminPass        string `json:"adminPass"`
	AdminPassConfirm string `json:"adminPassConfirm"`
	PurchaseLength   int    `json:"purchaseLength"`
	// 购买的虚机个数
	PurchaseNum int `json:"purchaseNum"`

	AutoRenewTimeUnit   string                `json:"autoRenewTimeUnit"`
	AutoRenewTime       int64                 `json:"autoRenewTime"`
	CreateEphemeralList []CreateEphemeralList `json:"createEphemeralList"`
	// 是否自动续费 默认即可 后付费不存在这个问题
	AutoRenew bool `json:"autoRenew"`
	// 镜像id 用默认即可 固定是ubuntu1604
	ImageID           string `json:"imageId"`
	OsName            string `json:"osName"`
	SecurityGroupName string `json:"securityGroupName"`
	// BCC
	ServiceType string `json:"serviceType"`
}

// CreateEphemeralList define storage
type CreateEphemeralList struct {
	// 磁盘存储类型 从页面创建虚机时 看到请求 默认是ssd
	StorageType string `json:"storageType"`
	// 磁盘大小
	SizeInGB int `json:"sizeInGB"`
}

// CdsOrderConfig define CDS order config
type CdsOrderConfig struct {
	// 付费类型，一期只支持postpay
	productType string `json:"productType"`
	// "zoneA"
	logicalZone    string `json:"logicalZone"`
	region         string `json:"sizeInGB"` // "bj"
	purchaseNum    int    `json:"sizeInGB"` // 1
	purchaseLength int    `json:"sizeInGB"` // 1
	autoRenewTime  int    `json:"sizeInGB"` // 0
	// "month"
	autoRenewTimeUnit string               `json:"sizeInGB"`
	cdsDiskSize       []bcc.DiskSizeConfig `json:"sizeInGB"`
	// "CDS"
	serviceType string `json:"sizeInGB"`
}

// EipOrderConfig define CDS order config
type EipOrderConfig struct {
	// 付费类型，一期只支持postpay
	ProductType     string `json:"productType"`
	BandwidthInMbps int    `json:"bandwidthInMbps"` // 1000
	Region          string `json:"region"`          // "bj"
	SubProductType  string `json:"subProductType"`  // "netraffic",
	// EIP购买数量应该是购买BCC数量的总和
	PurchaseNum       int    `json:"purchaseNum"`
	PurchaseLength    int    `json:"purchaseLength"`    // 1
	AutoRenewTime     int    `json:"autoRenewTime"`     // 0
	AutoRenewTimeUnit string `json:"autoRenewTimeUnit"` // "month",
	Name              string `json:"name"`              // "kkk"
	ServiceType       string `json:"serviceType"`       // "EIP"
}

// ScaleUpCluster scaleup a  cluster
func (c *Client) ScaleUpCluster(args *ScaleUpClusterArgs) (*ScaleUpClusterResponse, error) {
	var params map[string]string
	if args != nil {
		params = map[string]string{
			"clientToken": c.GenerateClientToken(),
			"scalingUp":   "",
		}
	}
	postContent, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	req, err := bce.NewRequest("POST", c.GetURL("v1/cluster", params), bytes.NewBuffer(postContent))
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
	var scResp *ScaleUpClusterResponse
	err = json.Unmarshal(bodyContent, &scResp)

	if err != nil {
		return nil, err
	}
	return scResp, nil
}
