package kong_entity

type KongRootNode struct {
	FormatVersion string    `yaml:"_format_version"`
	Info          Info      `yaml:"_info"`
	Services      []Service `yaml:"services"`
	Plugins       []Plugin  `yaml:"plugin,omitempty"`
}

func DefaultRootNode(version string) KongRootNode {
	return KongRootNode{
		FormatVersion: version,
		Info:          Info{},
		Services:      []Service{DefaultService()},
	}
}
