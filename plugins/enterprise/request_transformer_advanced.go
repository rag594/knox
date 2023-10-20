package enterprise

import "knox/kong_entity"

type RequestTransformerAdvancedConfig struct {
	HTTPMethod string    `yaml:"http_method"`
	RemoveOp   RemoveOp  `yaml:"remove"`
	RenameOp   RenameOp  `yaml:"rename"`
	ReplaceOp  ReplaceOp `yaml:"replace"`
	AddOp      AddOp     `yaml:"add"`
	AppendOp   AppendOp  `yaml:"append"`
	DotInKeys  bool      `yaml:"dot_in_keys"`
}

type RemoveOp struct {
	Body        []string `yaml:"body"`
	Headers     []string `yaml:"headers"`
	QueryString []string `yaml:"querystring"`
}

type RenameOp struct {
	Body        []string `yaml:"body"`
	Headers     []string `yaml:"headers"`
	QueryString []string `yaml:"querystring"`
}

type ReplaceOp struct {
	Body        []string `yaml:"body"`
	Headers     []string `yaml:"headers"`
	QueryString []string `yaml:"querystring"`
	JsonTypes   []string `yaml:"json_types"`
	URI         string   `yaml:"uri"`
}

type AddOp struct {
	Body        []string `yaml:"body"`
	Headers     []string `yaml:"headers"`
	QueryString []string `yaml:"querystring"`
	JsonTypes   []string `yaml:"json_types"`
}

type AppendOp struct {
	Body        []string `yaml:"body"`
	Headers     []string `yaml:"headers"`
	QueryString []string `yaml:"querystring"`
	JsonTypes   []string `yaml:"json_types"`
}

type AllowOp struct {
	Body []string `yaml:"body"`
}

func defaultRequestTransformerAdvanced() RequestTransformerAdvancedConfig {
	return RequestTransformerAdvancedConfig{
		ReplaceOp: ReplaceOp{
			URI: "/route/b",
		},
	}
}

func DefaultRequestTransformerAdvancedPlugin() kong_entity.Plugin {
	return kong_entity.Plugin{
		Config:   defaultRequestTransformerAdvanced(),
		Enabled:  true,
		Name:     "request-transformer-advanced",
		Protcols: []string{"http", "https"},
	}
}
