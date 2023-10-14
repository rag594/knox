package kong_entity

type Info struct {
	Defaults   Default  `yaml:"defaults"`
	SelectTags []string `yaml:"select_tags"`
}

type Default struct {
}
