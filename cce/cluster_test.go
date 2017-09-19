package cce

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/drinktee/bce-sdk-go/util"
)

func TestListInstances(t *testing.T) {
	ts := httptest.NewServer(ClusterHandler())
	defer ts.Close()
	cceClient.Endpoint = ts.URL
	cceClient.SetDebug(true)
	list, err := cceClient.ListInstances("a")

	if err != nil {
		t.Error(util.FormatTest("ListInstances", err.Error(), "nil"))
	}
	for _, ins := range list {
		fmt.Println(ins.VpcId)
	}
}

func TestScaleUp(t *testing.T) {
	ts := httptest.NewServer(ClusterHandler())
	defer ts.Close()
	cceClient.Endpoint = ts.URL
	cceClient.SetDebug(true)
	args := &ScaleUpClusterArgs{
		ClusterID: "c-NqYwWEhu",
		OrderContent: OrderContent{
			Items: []OrderItem{
				OrderItem{
					Config: BccOrderConfig{
						CPU: 100,
					},
				},
			},
		},
	}
	res, err := cceClient.ScaleUpCluster(args)

	if err != nil {
		t.Fatalf("ScaleUpCluster fail: %v", err)
	}
	if res.ClusterID != "c-NqYwWEhu" {
		t.Fatalf("ScaleUpCluster ClusterID fail")
	}
}
