package blb

import (
	"testing"

	"fmt"

	"github.com/drinktee/bce-sdk-go/util"
)

func TestCreateLoadBalance(t *testing.T) {
	blbClient.SetDebug(true)
	args := &CreateLoadBalancerArgs{
		Name: "golang-sun123",
	}
	blb, err := blbClient.CreateLoadBalancer(args)

	if err != nil {
		fmt.Println(err)
		t.Error(util.FormatTest("TestCreateLoadBalance", err.Error(), "nil"))

	} else {
		fmt.Println(blb.Address)
	}
}

func TestDescribeLoadBalancers(t *testing.T) {
	// blbClient.Endpoint = "bcc.bce-api.baidu.com"
	blbClient.SetDebug(true)
	args := &DescribeLoadBalancersArgs{
		LoadBalancerName: "bakendtest",
	}
	list, err := blbClient.DescribeLoadBalancers(args)

	if err != nil {
		fmt.Println(err)
		t.Error(util.FormatTest("TestDescribeLoadBalancers", err.Error(), "nil"))
	}
	fmt.Println(len(list))
	for _, blb := range list {
		fmt.Println(blb.PublicIp)
	}
}

func TestUpdateLoadBalancer(t *testing.T) {
	blbClient.SetDebug(true)
	args := &UpdateLoadBalancerArgs{
		LoadBalancerId: "lb-e5b33752",
		Name:           "golang-123",
	}
	err := blbClient.UpdateLoadBalancer(args)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteLoadBalancer(t *testing.T) {
	blbClient.SetDebug(true)
	args := &DeleteLoadBalancerArgs{
		LoadBalancerId: "lb-426fad2b",
	}
	err := blbClient.DeleteLoadBalancer(args)
	if err != nil {
		t.Error(err)
	}
}
