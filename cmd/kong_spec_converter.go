package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"sigs.k8s.io/yaml"
)

type KongSpecConverterCommand struct {
	Command *cli.Command
}

func NewKongSpecConverterCommand() *KongSpecConverterCommand {
	var (
		toFormat      string
		inputFilePath string
	)
	command := &cli.Command{
		Name:  "convert",
		Usage: "conert to JSON or YAML",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "to-format",
				Usage:       "provide the format using --to-format <yaml or json>",
				Destination: &toFormat,
				Category:    "converter",
			},
			&cli.StringFlag{
				Name:        "input",
				Usage:       "provide your input file using --input <file-path>",
				Destination: &inputFilePath,
				Category:    "converter",
			},
		},
		Action: func(cCtx *cli.Context) error {

			data, err := os.ReadFile(inputFilePath)
			if err != nil {
				fmt.Println(err)
				return err
			}

			if toFormat == "json" {
				jsonConvertedSpec, err := yaml.YAMLToJSON(data)
				if err != nil {
					fmt.Println("Err in converting yaml to json", err)
				}
				fmt.Println(string(jsonConvertedSpec))
				return nil
			}

			if toFormat == "yaml" {
				jsonConvertedSpec, err := yaml.JSONToYAML(data)
				if err != nil {
					fmt.Println("Err in converting json to yaml", err)
				}
				fmt.Println(string(jsonConvertedSpec))
				return nil
			}

			return nil
		},
	}

	return &KongSpecConverterCommand{Command: command}
}
