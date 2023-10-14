package main

import (
	"bytes"
	"fmt"
	"kong-config-generator/kong_entity"

	"gopkg.in/yaml.v3"
)

func main() {

	rootNode := kong_entity.DefaultRootNode("3.0")

	var confBytes bytes.Buffer

	yamlEncoder := yaml.NewEncoder(&confBytes)
	yamlEncoder.SetIndent(2)

	err := yamlEncoder.Encode(&rootNode)

	if err != nil {
		fmt.Println("Error ", err)
	}

	fmt.Println(" --- YAML ---")
	fmt.Println(string(confBytes.Bytes())) // yamlData will be in bytes. So converting it to string.
}
