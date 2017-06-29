package blb

import (
	"fmt"
	"testing"

	"github.com/drinktee/bce-sdk-go/util"
)

func TestAddBackendServers(t *testing.T) {
	blbClient.SetDebug(true)
	args := &AddBackendServersArgs{
		LoadBalancerId: "lb-e5b33752",
		BackendServerList: []BackendServer{
			BackendServer{
				InstanceId: "i-YWIy3FQx",
				Weight:     50,
			},
			BackendServer{
				InstanceId: "i-vfBlsqNG",
				Weight:     50,
			},
		},
	}
	err := blbClient.AddBackendServers(args)
	if err != nil {
		t.Error(err)
	}
}

func TestDescribeBackendServers(t *testing.T) {
	blbClient.SetDebug(true)
	args := &DescribeBackendServersArgs{
		LoadBalancerId: "lb-e5b33752",
	}
	list, err := blbClient.DescribeBackendServers(args)

	if err != nil {
		fmt.Println(err)
		t.Error(util.FormatTest("DescribeBackendServers", err.Error(), "nil"))
	}
	fmt.Println(len(list))
	for _, blb := range list {
		fmt.Println(blb)
	}
}

func TestUpdateBackendServers(t *testing.T) {
	blbClient.SetDebug(true)
	args := &UpdateBackendServersArgs{
		LoadBalancerId: "lb-e5b33752",
		BackendServerList: []BackendServer{BackendServer{
			InstanceId: "i-vfBlsqNG",
			Weight:     99,
		}},
	}
	err := blbClient.UpdateBackendServers(args)
	if err != nil {
		t.Error(err)
	}
}

func TestRemoveBackendServers(t *testing.T) {
	blbClient.SetDebug(true)
	args := &RemoveBackendServersArgs{
		LoadBalancerId:    "lb-e5b33752",
		BackendServerList: []string{"i-vfBlsqNG", "i-vfBlsqNG"},
	}

	err := blbClient.RemoveBackendServers(args)
	if err != nil {
		t.Error(err)
	}
}
