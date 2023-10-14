package kong_entity

type KongRootNode struct {
	FormatVersion string    `yaml:"_format_version"`
	Services      []Service `yaml:"services"`
	Routes        []Route   `yaml:"routes"`
	Plugins       []Plugin  `yaml:"plugin,omitempty"`
}
