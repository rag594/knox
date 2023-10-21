package plugins

import "knox/kong_entity"

type CORSConfig struct {
	Origins           []string `json:"origins" yaml:"origins"`
	Headers           []string `json:"headers" yaml:"headers"`
	ExposeHeaders     []string `json:"expose_headers" yaml:"expose_headers"`
	Methods           []string `json:"methods" yaml:"methods"`
	MaxAge            int      `json:"max_age" yaml:"max_age"`
	Credential        bool     `json:"credentials" yaml:"credentials"`
	PreFlightContinue bool     `json:"preflight_continue" yaml:"preflight_continue"`
}

func defaultCORSConfig() CORSConfig {
	return CORSConfig{
		Origins:           []string{"https://abc.com"},
		Headers:           []string{"Content-Type", "Origin"},
		PreFlightContinue: true,
	}
}

func DefaultCORSPlugin() kong_entity.Plugin {
	return kong_entity.Plugin{
		Config:   defaultCORSConfig(),
		Enabled:  true,
		Name:     "cors",
		Protcols: []string{"http", "https"},
	}
}
