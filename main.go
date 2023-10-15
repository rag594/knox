package main

import (
	"kong-config-generator/cmd"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	cliCommands := []*cli.Command{cmd.NewKongSpecTemplateCommand().Command}
	app := &cli.App{
		Name:     "CLI for generating kong-spec template",
		Commands: cliCommands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
