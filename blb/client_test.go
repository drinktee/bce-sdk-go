package blb

import (
	"os"

	"github.com/drinktee/bce-sdk-go/bce"
)

var blbClient *Client

func init() {
	var credentials, _ = bce.NewCredentialsFromFile("../aksk-test.json")

	//var bceConfig = bce.NewConfig(credentials)
	var bceConfig = &bce.Config{
		Credentials: credentials,
		Checksum:    true,
		Region:      os.Getenv("BOS_REGION"),
	}
	var bccConfig = NewConfig(bceConfig)
	blbClient = NewBLBClient(bccConfig)
}
