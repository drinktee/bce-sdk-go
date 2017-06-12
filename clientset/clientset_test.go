package clientset

import (
	"testing"

	"fmt"

	"github.com/drinktee/bce-sdk-go/bce"
)

var credentials, _ = bce.NewCredentialsFromFile("../aksk-test.json")

func TestNewFromConfig(t *testing.T) {
	cfg, err := bce.NewConfigFromFile("../aksk-test.json")
	if err != nil {
		t.Error(err)
	} else {
		cs, err := NewFromConfig(cfg)
		if err != nil {
			t.Error(err)
		} else {
			is, err := cs.Bcc().ListInstances(nil)
			if err != nil {
				t.Error(err)
			} else {
				for _, i := range is {
					fmt.Println(i.InstanceName)
				}
			}
		}

	}

}
