package cmd

import (
	"fmt"
	"knox/kong_entity"
	"os"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

type KongSpecVisualiseCommand struct {
	Command *cli.Command
}

type KongSpecGraph struct {
	Graph graph.Graph[string, string]
}

func NewKongSpecVisualiseCommand() *KongSpecVisualiseCommand {
	var (
		specFilePath   string
		outputFilePath string
	)
	command := &cli.Command{
		Name:  "visualise",
		Usage: "generate the graphviz",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "spec",
				Usage:       "load the file using --spec <file-path>",
				Destination: &specFilePath,
				Category:    "visualise",
			},
			&cli.StringFlag{
				Name:        "output",
				Usage:       "provide your output file using --output <file-path>",
				Destination: &outputFilePath,
				Category:    "visualise",
			},
		},
		Action: func(cCtx *cli.Context) error {

			data, err := os.ReadFile(specFilePath)
			if err != nil {
				fmt.Println(err)
				return err
			}

			var root kong_entity.KongRootNode

			err = yaml.Unmarshal(data, &root)

			if err != nil {
				fmt.Println("Err ", err)
			}

			g := graph.New(graph.StringHash, graph.Directed(), graph.Acyclic())

			specGraph := &KongSpecGraph{Graph: g}

			specGraph.CreateGraph(root)

			file, _ := os.Create(outputFilePath)
			_ = draw.DOT(g, file)

			return nil
		},
	}

	return &KongSpecVisualiseCommand{Command: command}
}

func (g *KongSpecGraph) CreateGraph(root kong_entity.KongRootNode) {
	for _, service := range root.Services {
		g.Graph.AddVertex(service.Name,
			graph.VertexAttribute("colorscheme", "blues3"),
			graph.VertexAttribute("style", "filled"),
			graph.VertexAttribute("color", "2"),
			graph.VertexAttribute("fillcolor", "1"))

		for _, route := range service.Routes {
			g.Graph.AddVertex(route.Name,
				graph.VertexAttribute("colorscheme", "greens3"),
				graph.VertexAttribute("style", "filled"),
				graph.VertexAttribute("color", "2"),
				graph.VertexAttribute("fillcolor", "1"))
			g.Graph.AddEdge(service.Name, route.Name)
			for _, rPlugin := range route.Plugins {
				routePluginNode := fmt.Sprintf("%s-%s", route.Name, rPlugin.Name)
				g.Graph.AddVertex(routePluginNode,
					graph.VertexAttribute("colorscheme", "purples3"),
					graph.VertexAttribute("style", "filled"),
					graph.VertexAttribute("color", "2"),
					graph.VertexAttribute("fillcolor", "1"))
				g.Graph.AddEdge(route.Name, routePluginNode)
			}
		}

		for _, sPlugin := range service.Plugins {
			servicePluginNode := fmt.Sprintf("%s-%s", service.Name, sPlugin.Name)
			g.Graph.AddVertex(servicePluginNode,
				graph.VertexAttribute("colorscheme", "reds3"),
				graph.VertexAttribute("style", "filled"),
				graph.VertexAttribute("color", "2"),
				graph.VertexAttribute("fillcolor", "1"))
			g.Graph.AddEdge(service.Name, servicePluginNode)
		}
	}
}
