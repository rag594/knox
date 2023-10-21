package kong_entity

type Info struct {
	Defaults   Default  `json:"defaults" yaml:"defaults"`
	SelectTags []string `json:"select_tags" yaml:"select_tags"`
}

type Default struct {
}
