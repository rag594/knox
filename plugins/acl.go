package plugins

import "knox/kong_entity"

type ACLConfig struct {
	Allow            []string `json:"allow" yaml:"allow"`
	Deny             []string `json:"deny" yaml:"deny"`
	HideGroupHeaders bool     `json:"hide_group_headers" yaml:"hide_group_headers"`
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
