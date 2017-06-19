package bcc

import (
	"fmt"
	"testing"
)

func TestListVpc(t *testing.T) {
	// ts := httptest.NewServer(EipHandler())
	// defer ts.Close()
	// eipClient.Endpoint = ts.URL
	// eips, err := eipClient.GetEips(nil)
	bccClient.Endpoint = "bcc.bce-api.baidu.com"
	args := ListVpcArgs{
		IsDefault: false,
	}
	vpcs, err := bccClient.ListVpc(&args)
	if err != nil {
		t.Error(err)
	}
	for _, vpc := range vpcs {
		fmt.Println(vpc.VpcID)
	}
}
