package client

import (
	"time"
)

// Protocol is a dictionary of domains
type Protocol struct {
	Domains map[string]Domain `json:"domains"`
}

// Domain describes type defs
type Domain struct {
	Inputs map[string]InputDef `json:"inputs"`
	// Outputs map[string]OutputDef `json:"outputs"`
	// Types map[string]TypeDef `json:"types"`
}

type InputDef struct {
	SourceOutputKey string `json:"sourceOutputKey"`
	InputMethod     string `json:"inputMethod"`
}

func (c *ApiClient) FetchProtocol() (protocol *Protocol, err error) {
	err = c.requestProtocol("/schema.json", &protocol)
	if err == nil {
		c.protocolLoadedAt = time.Now()
		c.protocol = protocol
	}
	return protocol, err
}

func (c *ApiClient) GetProtocol() (*Protocol, error) {
	if c.protocol != nil {
		return c.protocol, nil
	}

	return c.FetchProtocol()
}
