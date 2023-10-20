package kong_entity

type Service struct {
	Name           string   `yaml:"name"`
	Host           string   `yaml:"host"`
	Port           int      `yaml:"port"`
	Protocol       string   `yaml:"protocol"`
	Path           string   `yaml:"path"`
	Retries        int      `yaml:"retries"`
	Enabled        bool     `yaml:"enabled"`
	ReadTimeout    int      `yaml:"read_timeout"`
	ConnectTimeout int      `yaml:"connect_timeout"`
	WriteTimeout   int      `yaml:"write_timeout"`
	Plugins        []Plugin `yaml:"plugins,omitempty"`
	Routes         []Route  `yaml:"routes,omitempty"`
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
