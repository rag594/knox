package plugins

import "knox/kong_entity"

type PreFunctionConfig struct {
	Certificate     []string `json:"certificate" yaml:"certificate"`
	Rewrite         []string `json:"rewrite" yaml:"rewrite"`
	Access          []string `json:"access" yaml:"access"`
	HeaderFilter    []string `json:"header_filter" yaml:"header_filter"`
	BodyFilter      []string `json:"body_filter" yaml:"body_filter"`
	Log             []string `json:"log" yaml:"log"`
	WSHandShake     []string `json:"ws_handshake" yaml:"ws_handshake"`
	WSClientFrame   []string `json:"ws_client_frame" yaml:"ws_client_frame"`
	WSUpstreamFrame []string `json:"ws_upstream_frame" yaml:"ws_upstream_frame"`
	WSClose         []string `json:"ws_close" yaml:"ws_close"`
}

func defaultPreFunctionConfig() PreFunctionConfig {
	return PreFunctionConfig{}
}

func DefaultPreFunctionPlugin() kong_entity.Plugin {
	return kong_entity.Plugin{
		Config:   defaultPreFunctionConfig(),
		Enabled:  true,
		Name:     "pre-function",
		Protcols: []string{"http", "https"},
	}
}
