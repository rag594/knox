package enterprise

import "knox/kong_entity"

type ExitTransformerConfig struct {
	Functions        []string `yaml:"functions"`
	HandleUnknown    bool     `yaml:"handle_unknown"`
	HandleUnexpected bool     `yaml:"handle_unexpected"`
}

func defaultExitTransformerConfig() ExitTransformerConfig {
	return ExitTransformerConfig{}
}

func DefaultExitTransformerPlugin() kong_entity.Plugin {
	return kong_entity.Plugin{
		Config:   defaultExitTransformerConfig(),
		Enabled:  true,
		Name:     "exit-transformer",
		Protcols: []string{"http", "https"},
	}
}
