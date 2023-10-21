package plugins

import "knox/kong_entity"

type KeyAuthConfig struct {
	KeyNames        []string `json:"key_names" yaml:"key_names"`
	HideCredentials bool     `json:"hide_credentials" yaml:"hide_credentials"`
	Anonymous       string   `json:"anonymous,omitempty" yaml:"anonymous,omitempty"`
	KeyInHeader     bool     `json:"key_in_header" yaml:"key_in_header"`
	KeyInQuery      bool     `json:"key_in_query" yaml:"key_in_query"`
	KeyInBody       bool     `json:"key_in_body" yaml:"key_in_body"`
	RunOnPreFlight  bool     `json:"run_on_preflight" yaml:"run_on_preflight"`
}

func defaultKeyAuthConfig() KeyAuthConfig {
	return KeyAuthConfig{
		KeyNames:        []string{"apiKey"},
		HideCredentials: true,
		KeyInHeader:     true,
		RunOnPreFlight:  false,
	}
}

func DefaultKeyAuthPlugin() kong_entity.Plugin {
	return kong_entity.Plugin{
		Config:   defaultKeyAuthConfig(),
		Enabled:  true,
		Name:     "key-auth",
		Protcols: []string{"http", "https"},
	}
}
