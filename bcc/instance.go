package bcc

type Instance struct {
	InstanceId            string `json:"id"`
	InstanceName          string `json:"name"`
	Description           string `json:"desc"`
	Status                string `json:"status"`
	PaymentTiming         string `json:"paymentTiming"`
	CreationTime          string `json:"createTime"`
	ExpireTime            string `json:"expireTime"`
	PublicIP              string `json:"publicIp"`
	InternalIP            string `json:"internalIp"`
	CpuCount              int    `json:"cpuCount"`
	MemoryCapacityInGB    int    `json:"memoryCapacityInGB"`
	localDiskSizeInGB     int    `json:"localDiskSizeInGB"`
	ImageId               string `json:"imageId"`
	NetworkCapacityInMbps int    `json:"networkCapacityInMbps"`
	PlacementPolicy       string `json:"placementPolicy"`
	ZoneName              string `json:"zoneName"`
	SubnetId              string `json:"subnetId"`
	VpcId                 string `json:"vpcId"`
}

type ListInstancesResponse struct {
	Marker      string     `json:"marker"`
	IsTruncated bool       `json:"isTruncated"`
	NextMarker  string     `json:"nextMarker"`
	MaxKeys     int        `json:"maxKeys"`
	Instances   []Instance `json:"instances"`
}

type GetInstanceResponse struct {
	Ins Instance `json:"instance"`
}
