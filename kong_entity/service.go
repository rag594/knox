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
	Plugins        []Plugin `yaml:"plugin,omitempty"`
	Routes         []Route  `yaml:"routes,omitempty"`
}

func DefaultService() Service {
	return Service{
		Name:           "DefaultService",
		Host:           "xyz.com",
		Port:           80,
		Path:           "/",
		Retries:        0,
		Enabled:        true,
		ReadTimeout:    60000,
		ConnectTimeout: 60000,
		WriteTimeout:   60000,
		Routes:         []Route{DefaultRoute()},
	}
}
