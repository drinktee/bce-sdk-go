package blb

import (
	"fmt"
	"testing"

	"github.com/drinktee/bce-sdk-go/util"
)

func TestCreateTCPListener(t *testing.T) {
	blbClient.SetDebug(true)
	args := &CreateTCPListenerArgs{
		LoadBalancerId: "lb-f5d263e5",
		ListenerPort:   8080,
		BackendPort:    8080,
		Scheduler:      "LeastConnection",
	}
	err := blbClient.CreateTCPListener(args)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateUDPListener(t *testing.T) {
	blbClient.SetDebug(true)
	args := &CreateUDPListenerArgs{
		LoadBalancerId:    "lb-f5d263e5",
		ListenerPort:      8888,
		BackendPort:       8888,
		Scheduler:         "LeastConnection",
		HealthCheckString: "hello",
	}
	err := blbClient.CreateUDPListener(args)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateHTTPListener(t *testing.T) {
	blbClient.SetDebug(true)
	args := &CreateHTTPListenerArgs{
		LoadBalancerId: "lb-f5d263e5",
		ListenerPort:   8899,
		BackendPort:    8899,
		Scheduler:      "LeastConnection",
	}
	err := blbClient.CreateHTTPListener(args)
	if err != nil {
		t.Error(err)
	}
}

func TestDescribeTCPListener(t *testing.T) {
	blbClient.SetDebug(true)
	args := &DescribeTCPListenerArgs{
		LoadBalancerId: "lb-f5d263e5",
		ListenerPort:   80,
	}
	list, err := blbClient.DescribeTCPListener(args)

	if err != nil {
		fmt.Println(err)
		t.Error(util.FormatTest("ListInstances", err.Error(), "nil"))
	}
	fmt.Println(len(list))
	for _, blb := range list {
		fmt.Println(blb.ListenerPort)
	}
}

func TestDescribeUDPListener(t *testing.T) {
	blbClient.SetDebug(true)
	args := &DescribeUDPListenerArgs{
		LoadBalancerId: "lb-f5d263e5",
		// ListenerPort:   80,
	}
	list, err := blbClient.DescribeUDPListener(args)

	if err != nil {
		fmt.Println(err)
		t.Error(util.FormatTest("DescribeUDPListener", err.Error(), "nil"))
	}
	fmt.Println(len(list))
	for _, blb := range list {
		fmt.Println(blb.ListenerPort)
	}
}
