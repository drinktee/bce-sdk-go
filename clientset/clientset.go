package clientset

import (
	"fmt"

	"github.com/drinktee/bce-sdk-go/bcc"
	"github.com/drinktee/bce-sdk-go/bce"
	"github.com/drinktee/bce-sdk-go/blb"
	"github.com/drinktee/bce-sdk-go/cce"
	"github.com/drinktee/bce-sdk-go/eip"
)

// Interface contains all methods of clients
type Interface interface {
	Bcc() *bcc.Client
	Blb() *blb.Client
	Eip() *eip.Client
	Cce() *cce.Client
}

// Clientset contains the clients for groups.
type Clientset struct {
	BccClient *bcc.Client
	BlbClient *blb.Client
	EipClient *eip.Client
	CceClient *cce.Client
}

// Bcc retrieves the BccClient
func (c *Clientset) Bcc() *bcc.Client {
	if c == nil {
		return nil
	}
	return c.BccClient
}

// Blb retrieves the BccClient
func (c *Clientset) Blb() *blb.Client {
	if c == nil {
		return nil
	}
	return c.BlbClient
}

// Eip retrieves the BccClient
func (c *Clientset) Eip() *eip.Client {
	if c == nil {
		return nil
	}
	return c.EipClient
}

// Cce retrieves the CceClient
func (c *Clientset) Cce() *cce.Client {
	if c == nil {
		return nil
	}
	return c.CceClient
}

// NewFromConfig create a new Clientset for the given config.
func NewFromConfig(cfg *bce.Config) (*Clientset, error) {
	if cfg == nil {
		return nil, fmt.Errorf("Config cannot be nil")
	}
	var cs Clientset
	var cceCfg = *cfg
	bccConfig := bcc.NewConfig(cfg)
	blbConfig := blb.NewConfig(cfg)
	eipConfig := eip.NewConfig(cfg)
	// cce endpoint is different
	cceConfig := cce.NewConfig(&cceCfg)
	cs.BccClient = bcc.NewClient(bccConfig)
	cs.BlbClient = blb.NewBLBClient(blbConfig)
	cs.EipClient = eip.NewEIPClient(eipConfig)
	cs.CceClient = cce.NewClient(cceConfig)
	return &cs, nil

}
