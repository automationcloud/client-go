package client

type Protocol struct {
	Domains map[string]Domain `json:"domains"`
}

type Domain struct {
	Inputs map[string]InputDef `json:"inputs"`
	// Outputs map[string]OutputDef `json:"outputs"`
	// Types map[string]TypeDef `json:"types"`
}

type InputDef struct {
	SourceOutputKey string `json:"sourceOutputKey"`
	InputMethod     string `json:"inputMethod"`
}

func (c *ApiClient) FetchProtocol() (protocol Protocol, err error) {
	err = c.Get("https://protocol.automationcloud.net/schema.json", &protocol)
	return
}

var protocol Protocol
var protocolLoaded = false

func (c *ApiClient) GetProtocol() (Protocol, error) {
	var err error
	if !protocolLoaded {
		protocolLoaded = true
		protocol, err = c.FetchProtocol()
		return protocol, err
	}

	return protocol, nil
}
