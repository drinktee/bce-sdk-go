package eip

import (
	"fmt"
	"testing"
)

var testIP string

func TestCreateEip(t *testing.T) {
	bill := &Billing{
		PaymentTiming: "Postpaid",
		BillingMethod: "ByTraffic",
	}
	args := &CreateEipArgs{
		BandwidthInMbps: 998,
		Billing:         bill,
		Name:            "golangtest",
	}
	ip, err := eipClient.CreateEip(args)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ip)
	testIP = ip
}

func TestResizeEip(t *testing.T) {
	args := &ResizeEipArgs{
		BandwidthInMbps: 111,
		Ip:              "180.76.242.209",
	}
	err := eipClient.ResizeEip(args)
	if err != nil {
		t.Error(err)
	}
}

func TestBindEip(t *testing.T) {
	args := &BindEipArgs{
		Ip:           "180.76.242.209",
		InstanceType: "BLB",
		InstanceId:   "lb-f5d263e5",
	}
	err := eipClient.BindEip(args)
	if err != nil {
		t.Error(err)
	}
}

func TestUnbindEip(t *testing.T) {
	args := &EipArgs{
		Ip: "180.76.242.209",
	}
	err := eipClient.UnbindEip(args)
	if err != nil {
		t.Error(err)
	}
}
func TestDeleteEip(t *testing.T) {
	args := &EipArgs{
		Ip: testIP,
	}
	err := eipClient.DeleteEip(args)
	if err != nil {
		t.Error(err)
	}
}
func TestGetEips(t *testing.T) {
	eips, err := eipClient.GetEips(nil)
	if err != nil {
		t.Error(err)
	}
	for _, eip := range eips {
		fmt.Println(eip.Eip)
	}
}
func TestPurchaseReservedEips(t *testing.T) {

}
