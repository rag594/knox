package plugins

import "knox/kong_entity"

type PreFunctionConfig struct {
	Certificate     []string `yaml:"certificate"`
	Rewrite         []string `yaml:"rewrite"`
	Access          []string `yaml:"access"`
	HeaderFilter    []string `yaml:"header_filter"`
	BodyFilter      []string `yaml:"body_filter"`
	Log             []string `yaml:"log"`
	WSHandShake     []string `yaml:"ws_handshake"`
	WSClientFrame   []string `yaml:"ws_client_frame"`
	WSUpstreamFrame []string `yaml:"ws_upstream_frame"`
	WSClose         []string `yaml:"ws_close"`
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
