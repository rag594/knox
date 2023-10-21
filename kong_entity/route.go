package kong_entity

type Route struct {
	Name                   string   `json:"name" yaml:"name"`
	Protocols              []string `json:"protocols" yaml:"protocols"`
	Methods                []string `json:"methods" yaml:"methods"`
	Hosts                  []string `json:"hosts" yaml:"hosts"`
	Paths                  []string `json:"paths" yaml:"paths"`
	HTTPRedirectStatusCode int      `json:"https_redirect_status_code" yaml:"https_redirect_status_code"`
	RegexPriority          int      `json:"regex_priority" yaml:"regex_priority"`
	StripPath              bool     `json:"strip_path" yaml:"strip_path"`
	PathHandling           string   `json:"path_handling" yaml:"path_handling"`
	PreserveHost           bool     `json:"preserve_host" yaml:"preserve_host"`
	RequestBuffering       bool     `json:"request_buffering" yaml:"request_buffering"`
	ResponseBuffering      bool     `json:"response_buffering" yaml:"response_buffering"`
	Plugins                []Plugin `json:"plugins,omitempty" yaml:"plugins,omitempty"`
}

func (r *Route) AddPlugin(p Plugin) {
	r.Plugins = append(r.Plugins, p)
}

func WithRoutePlugins(p []Plugin) func(*Route) {
	return func(r *Route) {
		for _, plugin := range p {
			r.AddPlugin(plugin)
		}
	}
}

func NewRoute(options ...func(*Route)) *Route {
	r := &Route{
		Name:                   "route1",
		Protocols:              []string{"http", "https"},
		Methods:                []string{"POST"},
		Hosts:                  []string{"abc.com"},
		Paths:                  []string{"/abc"},
		HTTPRedirectStatusCode: 426,
		RegexPriority:          0,
		StripPath:              false,
		PathHandling:           "v0",
		PreserveHost:           true,
		RequestBuffering:       true,
		ResponseBuffering:      true,
	}

	for _, opt := range options {
		opt(r)
	}

	return r
}
