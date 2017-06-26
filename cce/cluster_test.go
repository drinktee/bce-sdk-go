package cce

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/drinktee/bce-sdk-go/util"
)

func TestListInstances(t *testing.T) {
	ts := httptest.NewServer(InstancesHandler())
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
