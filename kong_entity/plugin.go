package kong_entity

type Plugin struct {
	Config   interface{} `yaml:"config"`
	Enabled  bool        `yaml:"enabled"`
	Name     string      `yaml:"name"`
	Protcols []string    `yaml:"protocols"`
}

func (p Plugin) PluginName() string {
	return p.Name
}
