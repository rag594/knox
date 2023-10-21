package kong_entity

type Service struct {
	Name           string   `json:"name" yaml:"name"`
	Host           string   `json:"host" yaml:"host"`
	Port           int      `json:"port" yaml:"port"`
	Protocol       string   `json:"protocol" yaml:"protocol"`
	Path           string   `json:"path" yaml:"path"`
	Retries        int      `json:"retries" yaml:"retries"`
	Enabled        bool     `json:"enabled" yaml:"enabled"`
	ReadTimeout    int      `json:"read_timeout" yaml:"read_timeout"`
	ConnectTimeout int      `json:"connect_timeout" yaml:"connect_timeout"`
	WriteTimeout   int      `json:"write_timeout" yaml:"write_timeout"`
	Plugins        []Plugin `json:"plugins,omitempty" yaml:"plugins,omitempty"`
	Routes         []Route  `json:"routes,omitempty" yaml:"routes,omitempty"`
}

func (s *Service) AddPlugin(p Plugin) {
	s.Plugins = append(s.Plugins, p)
}

func (s *Service) AddRoute(r Route) {
	s.Routes = append(s.Routes, r)
}

func WithServicePlugins(p []Plugin) func(*Service) {
	return func(s *Service) {
		for _, plugin := range p {
			s.AddPlugin(plugin)
		}
	}
}

func WithRoutes(r []Route) func(*Service) {
	return func(s *Service) {
		for _, route := range r {
			s.AddRoute(route)
		}
	}
}

func NewService(options ...func(*Service)) *Service {
	s := &Service{
		Name:           "DefaultService",
		Host:           "xyz.com",
		Port:           80,
		Path:           "/",
		Retries:        0,
		Enabled:        true,
		ReadTimeout:    60000,
		ConnectTimeout: 60000,
		WriteTimeout:   60000,
	}

	for _, opt := range options {
		opt(s)
	}

	return s
}
