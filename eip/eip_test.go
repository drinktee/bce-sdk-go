package eip

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestCreateEip(t *testing.T) {
	ts := httptest.NewServer(EipHandler())
	defer ts.Close()
	eipClient.Endpoint = ts.URL
	bill := &Billing{
		PaymentTiming: "Postpaid",
		BillingMethod: "ByTraffic",
	}
	args := &CreateEipArgs{
		BandwidthInMbps: 998,
		Billing:         bill,
		Name:            "k8sblbtest",
	}
	ip, err := eipClient.CreateEip(args)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ip)
}

var expectResizeEip = &ResizeEipArgs{
	BandwidthInMbps: 111,
	Ip:              "180.76.242.209",
}

func TestResizeEip(t *testing.T) {
	ts := httptest.NewServer(EipHandler())
	defer ts.Close()
	eipClient.Endpoint = ts.URL
	args := &ResizeEipArgs{
		BandwidthInMbps: 111,
		Ip:              "180.76.242.209",
	}
	err := eipClient.ResizeEip(args)
	if err != nil {
		t.Error(err)
	}
}

var expectBindEip = &BindEipArgs{
	Ip:           "180.76.242.209",
	InstanceType: "BLB",
	InstanceId:   "lb-f5d263e5",
}

func TestBindEip(t *testing.T) {
	// ts := httptest.NewServer(EipHandler())
	// defer ts.Close()
	// eipClient.Endpoint = ts.URL
	eipClient.SetDebug(true)
	args := &BindEipArgs{
		Ip:           "180.76.247.163",
		InstanceType: "BLB",
		InstanceId:   "lb-ccb6c6bf",
	}
	err := eipClient.BindEip(args)
	if err != nil {
		t.Error(err)
	}
}

var expectUnbindEip = &EipArgs{
	Ip: "180.76.247.182",
}

func TestUnbindEip(t *testing.T) {
	ts := httptest.NewServer(EipHandler())
	defer ts.Close()
	eipClient.Endpoint = ts.URL
	err := eipClient.UnbindEip(expectUnbindEip)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteEip(t *testing.T) {
	// ts := httptest.NewServer(EipHandler())
	// defer ts.Close()
	// eipClient.Endpoint = ts.URL
	err := eipClient.DeleteEip(expectUnbindEip)
	if err != nil {
		t.Error(err)
	}
}
func TestGetEips(t *testing.T) {
	ts := httptest.NewServer(EipHandler())
	defer ts.Close()
	eipClient.Endpoint = ts.URL
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
