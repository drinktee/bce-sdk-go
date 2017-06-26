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
	// bccClient.Endpoint = "bcc.bce-api.baidu.com"
	bccClient.SetDebug(true)
	args := ListRouteArgs{
		VpcID: "45d38bda-00b1-4146-a40a-27885d473353",
	}
	rs, err := bccClient.ListRouteTable(&args)
	if err != nil {
		t.Error(err)
	}
	for _, r := range rs {
		// fmt.Printf("%+v", r)
		fmt.Println(r.RouteRuleID)
		// fmt.Println(r.NexthopID)
	}
}

func TestCreateRouteRule(t *testing.T) {
	// bccClient.Endpoint = "bcc.bce-api.baidu.com"
	args := CreateRouteRuleArgs{
		RouteTableID:       "rt-wc5rd05e8fzs",
		SourceAddress:      "0.0.0.0/0",
		DestinationAddress: "172.17.112.0/24",
		NexthopID:          "i-ddUE7vVn",
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
	// bccClient.Endpoint = "bcc.bce-api.baidu.com"
	err := bccClient.DeleteRoute("rr-p9dbxrxdcsrh")
	if err != nil {
		t.Error(err)
	}
}
