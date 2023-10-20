package plugins

import "knox/kong_entity"

type ACLConfig struct {
	Allow            []string `yaml:"allow"`
	Deny             []string `yaml:"deny"`
	HideGroupHeaders bool     `yaml:"hide_group_headers"`
}

func defaultACLConfig() ACLConfig {
	return ACLConfig{}
}

func DefaultACLPlugin() kong_entity.Plugin {
	return kong_entity.Plugin{
		Config:   defaultACLConfig(),
		Enabled:  true,
		Name:     "acl",
		Protcols: []string{"http", "https"},
	}
}
