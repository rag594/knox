package enterprise

import "knox/kong_entity"

type ExitTransformerConfig struct {
	Functions        []string `json:"functions" yaml:"functions"`
	HandleUnknown    bool     `json:"handle_unknown" yaml:"handle_unknown"`
	HandleUnexpected bool     `json:"handle_unexpected" yaml:"handle_unexpected"`
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
