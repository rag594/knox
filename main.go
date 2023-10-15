package main

import (
	"bytes"
	"fmt"
	"kong-config-generator/kong_entity"
	"kong-config-generator/plugins"

	"gopkg.in/yaml.v3"
)

func main() {

	routes := []kong_entity.Route{*kong_entity.NewRoute()}
	plugins := []kong_entity.Plugin{plugins.DefaultKeyAuthPlugin()}
	service := kong_entity.NewService(kong_entity.WithRoutes(routes), kong_entity.WithServicePlugins(plugins))
	rootNode := kong_entity.NewKongRootNode(kong_entity.WithServices([]kong_entity.Service{*service}))

	var confBytes bytes.Buffer

	yamlEncoder := yaml.NewEncoder(&confBytes)
	yamlEncoder.SetIndent(2)

	err := yamlEncoder.Encode(&rootNode)

	if err != nil {
		fmt.Println("Error ", err)
	}

	fmt.Println(string(confBytes.Bytes())) // yamlData will be in bytes. So converting it to string.
}
