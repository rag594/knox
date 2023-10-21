package kong_entity

type Plugin struct {
	Config   interface{} `json:"config" yaml:"config"`
	Enabled  bool        `json:"enabled" yaml:"enabled"`
	Name     string      `json:"name" yaml:"name"`
	Protcols []string    `json:"protocols" yaml:"protocols"`
}

func (p Plugin) PluginName() string {
	return p.Name
}
