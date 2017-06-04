package bcc

import (
	"os"
	"testing"

	"fmt"

	"github.com/drinktee/bce-sdk-go/bce"
	"github.com/drinktee/bce-sdk-go/util"
)

var credentials = bce.NewCredentials("0b565874be53440ba579d8a70cac8f98", "d8ad18e630c64bf4b4661ce9b49164c8")

//var bceConfig = bce.NewConfig(credentials)
var bceConfig = &bce.Config{
	Credentials: credentials,
	Checksum:    true,
	Region:      os.Getenv("BOS_REGION"),
}
var bccConfig = NewConfig(bceConfig)
var bccClient = NewClient(bccConfig)

func TestListInstances(t *testing.T) {
	bccClient.SetDebug(true)
	list, err := bccClient.ListInstances(nil)

	if err != nil {
		fmt.Println(err)
		t.Error(util.FormatTest("ListInstances", err.Error(), "nil"))
	}

	for _, ins := range list {
		fmt.Println(ins.InstanceId)
	}
}

func TestGetInstance(t *testing.T) {
	// bccClient.SetDebug(true)
	ins, err := bccClient.GetInstance("i-5fOFPL5J", nil)

	if err != nil {
		fmt.Println(err)
		t.Error(util.FormatTest("ListInstances", err.Error(), "nil"))
	}
	fmt.Printf("%+v\n", ins)
}
