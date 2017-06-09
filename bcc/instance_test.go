package bcc

import (
	"testing"

	"net/http/httptest"

	"github.com/drinktee/bce-sdk-go/util"
)

func TestDescribeInstance(t *testing.T) {
	ts := httptest.NewServer(InstancesHandler())
	defer ts.Close()
	// bccClient.SetDebug(true)
	bccClient.Endpoint = ts.URL
	ins, err := bccClient.DescribeInstance("i-YufwpQAe", nil)
	if err != nil {
		t.Error(util.FormatTest("ListInstances", err.Error(), "nil"))
	}
	if ins.InstanceName != "instance-luz2ef4l-1" {
		t.Error("name error!")
	}
}

func TestListInstances(t *testing.T) {
	ts := httptest.NewServer(InstancesHandler())
	defer ts.Close()
	bccClient.Endpoint = ts.URL
	list, err := bccClient.ListInstances(nil)
	if err != nil {
		t.Error(util.FormatTest("ListInstances", err.Error(), "nil"))
	}
	for _, ins := range list {
		if ins.InstanceId != "i-IyWRtII7" {
			t.Error("instanceId error")
		}
	}
}
