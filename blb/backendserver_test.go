package blb

import (
	"fmt"
	"testing"

	"github.com/drinktee/bce-sdk-go/util"
)

func TestAddBackendServers(t *testing.T) {
	blbClient.SetDebug(true)
	args := &AddBackendServersArgs{
		LoadBalancerId: "lb-f5d263e5",
		BackendServerList: []BackendServer{
			BackendServer{
				InstanceId: "i-dSo6P8oU",
				Weight:     50,
			},
			BackendServer{
				InstanceId: "i-3VgKJmSh",
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
		LoadBalancerId: "lb-f5d263e5",
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
		LoadBalancerId:    "lb-f5d263e5",
		BackendServerList: []BackendServer{},
	}
	err := blbClient.UpdateBackendServers(args)
	if err != nil {
		t.Error(err)
	}
}

func TestRemoveBackendServers(t *testing.T) {
	blbClient.SetDebug(true)
	args := &RemoveBackendServersArgs{
		LoadBalancerId:    "lb-f5d263e5",
		BackendServerList: []string{"i-dSo6P8oU"},
	}

	err := blbClient.RemoveBackendServers(args)
	if err != nil {
		t.Error(err)
	}
}
