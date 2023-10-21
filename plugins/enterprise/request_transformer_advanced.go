package enterprise

import "knox/kong_entity"

type RequestTransformerAdvancedConfig struct {
	HTTPMethod string    `json:"http_method" yaml:"http_method"`
	RemoveOp   RemoveOp  `json:"remove" yaml:"remove"`
	RenameOp   RenameOp  `json:"rename" yaml:"rename"`
	ReplaceOp  ReplaceOp `json:"replace" yaml:"replace"`
	AddOp      AddOp     `json:"add" yaml:"add"`
	AppendOp   AppendOp  `json:"append" yaml:"append"`
	DotInKeys  bool      `json:"dot_in_keys" yaml:"dot_in_keys"`
}

type RemoveOp struct {
	Body        []string `json:"body" yaml:"body"`
	Headers     []string `json:"headers" yaml:"headers"`
	QueryString []string `json:"querystring" yaml:"querystring"`
}

type RenameOp struct {
	Body        []string `json:"body" yaml:"body"`
	Headers     []string `json:"headers" yaml:"headers"`
	QueryString []string `json:"querystring" yaml:"querystring"`
}

type ReplaceOp struct {
	Body        []string `json:"body" yaml:"body"`
	Headers     []string `json:"headers" yaml:"headers"`
	QueryString []string `json:"querystring" yaml:"querystring"`
	JsonTypes   []string `json:"json_types" yaml:"json_types"`
	URI         string   `json:"uri" yaml:"uri"`
}

type AddOp struct {
	Body        []string `json:"body" yaml:"body"`
	Headers     []string `json:"headers" yaml:"headers"`
	QueryString []string `json:"querystring" yaml:"querystring"`
	JsonTypes   []string `json:"json_types" yaml:"json_types"`
}

type AppendOp struct {
	Body        []string `json:"body" yaml:"body"`
	Headers     []string `json:"headers" yaml:"headers"`
	QueryString []string `json:"querystring" yaml:"querystring"`
	JsonTypes   []string `json:"json_types" yaml:"json_types"`
}

type AllowOp struct {
	Body []string `json:"body" yaml:"body"`
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
