package plugins

import "knox/kong_entity"

type KeyAuthConfig struct {
	KeyNames        []string `yaml:"key_names"`
	HideCredentials bool     `yaml:"hide_credentials"`
	Anonymous       string   `yaml:"anonymous,omitempty"`
	KeyInHeader     bool     `yaml:"key_in_header"`
	KeyInQuery      bool     `yaml:"key_in_query"`
	KeyInBody       bool     `yaml:"key_in_body"`
	RunOnPreFlight  bool     `yaml:"run_on_preflight"`
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
