package bcc

import (
	"fmt"
	"testing"
)

func TestListRouteTable(t *testing.T) {
	// ts := httptest.NewServer(EipHandler())
	// defer ts.Close()
	// eipClient.Endpoint = ts.URL
	// eips, err := eipClient.GetEips(nil)
	bccClient.Endpoint = "bcc.bce-api.baidu.com"
	bccClient.SetDebug(true)
	args := ListRouteArgs{
		VpcID: "vpc-y9f84g5peuzi",
	}
	rs, err := bccClient.ListRouteTable(&args)
	if err != nil {
		t.Error(err)
	}
	for _, r := range rs {
		// fmt.Printf("%+v", r)
		fmt.Println(r.NexthopID)
	}
}

func TestCreateRouteRule(t *testing.T) {
	bccClient.Endpoint = "bcc.bce-api.baidu.com"
	args := CreateRouteRuleArgs{
		RouteTableID:       "rt-kq2wgkeshqqy",
		SourceAddress:      "0.0.0.0/0",
		DestinationAddress: "172.17.112.0/24",
		NexthopID:          "i-zcimqrvw",
		NexthopType:        "custom",
		Description:        "a",
	}
	id, err := bccClient.CreateRouteRule(&args)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(id)
}

func TestDeleteRoute(t *testing.T) {
	bccClient.Endpoint = "bcc.bce-api.baidu.com"
	err := bccClient.DeleteRoute("rr-zrsr6sacnytq")
	if err != nil {
		t.Error(err)
	}
}
