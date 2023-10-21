package kong_entity

import (
	"bytes"

	"gopkg.in/yaml.v3"
)

type KongRootNode struct {
	FormatVersion string    `json:"_format_version" yaml:"_format_version"`
	Info          Info      `json:"_info" yaml:"_info"`
	Services      []Service `json:"services" yaml:"services"`
	Plugins       []Plugin  `json:"plugins,omitempty" yaml:"plugins,omitempty"`
}

func (k *KongRootNode) AddService(s Service) {
	k.Services = append(k.Services, s)
}

func (r *KongRootNode) AddPlugin(p Plugin) {
	r.Plugins = append(r.Plugins, p)
}

func WithServices(s []Service) func(*KongRootNode) {
	return func(k *KongRootNode) {
		for _, service := range s {
			k.AddService(service)
		}
	}
}

func WithRootNodePlugins(p []Plugin) func(*KongRootNode) {
	return func(k *KongRootNode) {
		for _, plugin := range p {
			k.AddPlugin(plugin)
		}
	}
}

func NewKongRootNode(options ...func(*KongRootNode)) *KongRootNode {
	k := &KongRootNode{
		FormatVersion: "3.0",
		Info:          Info{},
	}

	for _, opt := range options {
		opt(k)
	}

	return k
}

func (k *KongRootNode) ToYaml() string {
	var confBytes bytes.Buffer

	yamlEncoder := yaml.NewEncoder(&confBytes)
	yamlEncoder.SetIndent(2)

	err := yamlEncoder.Encode(k)

	if err != nil {
		panic("error in encoding to yaml")
	}

	return string(confBytes.Bytes())
}
