package kong_entity

type Route struct {
	Name                   string   `yaml:"name"`
	Protocols              []string `yaml:"protocols"`
	Methods                []string `yaml:"methods"`
	Hosts                  []string `yaml:"hosts"`
	Paths                  []string `yaml:"paths"`
	HTTPRedirectStatusCode int      `yaml:"https_redirect_status_code"`
	RegexPriority          int      `yaml:"regex_priority"`
	StripPath              bool     `yaml:"strip_path"`
	PathHandling           string   `yaml:"path_handling"`
	PreserveHost           bool     `yaml:"preserve_host"`
	RequestBuffering       bool     `yaml:"request_buffering"`
	ResponseBuffering      bool     `yaml:"response_buffering"`
	Plugins                []Plugin `yaml:"plugin,omitempty"`
}

func DefaultRoute() Route {
	return Route{
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
}
