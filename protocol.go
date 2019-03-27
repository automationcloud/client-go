package client

import (
	"time"
)

// Protocol is a dictionary of domains.
type Protocol struct {
	Domains map[string]Domain `json:"domains"`
}

// Domain describes type defs.
type Domain struct {
	Inputs map[string]InputDef `json:"inputs"`
	// Outputs map[string]OutputDef `json:"outputs"`
	// Types map[string]TypeDef `json:"types"`
}

// InputDef defines types of inputs allowed by automation job.
type InputDef struct {
	SourceOutputKey string `json:"sourceOutputKey"`
	InputMethod     string `json:"inputMethod"`
}

// FetchProtocol requests protocol containing type definitions for all automation domains.
func (c *ApiClient) FetchProtocol() (protocol *Protocol, err error) {
	err = c.requestProtocol("/schema.json", &protocol)
	if err == nil {
		c.protocolLoadedAt = time.Now()
		c.protocol = protocol
	}
	return protocol, err
}

// GetProtocol returns cached version of a protocol.
func (c *ApiClient) GetProtocol() (*Protocol, error) {
	if c.protocol != nil {
		return c.protocol, nil
	}

	return c.FetchProtocol()
}
