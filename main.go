package main

import (
	"knox/cmd"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	cliCommands := []*cli.Command{
		cmd.NewKongSpecTemplateCommand().Command,
		cmd.NewKongSpecVisualiseCommand().Command,
		cmd.NewKongSpecConverterCommand().Command}
	app := &cli.App{
		Name:     "CLI for generating/visualising kong-spec",
		Commands: cliCommands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
