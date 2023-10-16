package plugins

import "knox/kong_entity"

type CORSConfig struct {
	Origins           []string `yaml:"origins"`
	Headers           []string `yaml:"headers"`
	ExposeHeaders     []string `yaml:"expose_headers"`
	Methods           []string `yaml:"methods"`
	MaxAge            int      `yaml:"max_age"`
	Credential        bool     `yaml:"credentials"`
	PreFlightContinue bool     `yaml:"preflight_continue"`
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
