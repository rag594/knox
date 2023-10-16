package cmd

import (
	"fmt"
	"knox/kong_entity"
	"knox/plugins"
	"strings"

	"github.com/urfave/cli/v2"
)

type KongSpecTemplateCommand struct {
	Command *cli.Command
}

func NewKongSpecTemplateCommand() *KongSpecTemplateCommand {
	var (
		servicePluginsFlag string
		routePluginsFlag   string
		globalPluginsFlag  string
	)
	command := &cli.Command{
		Name:  "generate",
		Usage: "generate kong-spec template",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "service-plugins",
				Usage:       "add service-plugins in comma-separated format. Example --service-plugins key-auth,cors ",
				Destination: &servicePluginsFlag,
			},
			&cli.StringFlag{
				Name:        "route-plugins",
				Usage:       "add route-plugins in comma-separated format. Example --route-plugins key-auth,cors ",
				Destination: &routePluginsFlag,
			},
			&cli.StringFlag{
				Name:        "global-plugins",
				Usage:       "add global-plugins in comma-separated format. Example --global-plugins key-auth,cors ",
				Destination: &globalPluginsFlag,
			},
		},
		Action: func(cCtx *cli.Context) error {

			routes := getRoutes(routePluginsFlag)
			services := getServices(servicePluginsFlag, routes)
			rootNode := kong_entity.NewKongRootNode(kong_entity.WithServices(services))

			fmt.Println(rootNode.ToYaml())

			return nil
		},
	}

	return &KongSpecTemplateCommand{Command: command}
}

func getServices(servicePluginsFlag string, routes []kong_entity.Route) []kong_entity.Service {
	servicePlugins := []kong_entity.Plugin{}
	if len(routes) == 0 {
		route := kong_entity.NewRoute()
		routes = []kong_entity.Route{*route}
	}

	if servicePluginsFlag != "" {
		servicePluginsSplit := strings.Split(servicePluginsFlag, ",")
		for _, sPlugins := range servicePluginsSplit {
			servicePlugins = append(servicePlugins, plugins.GetPluginFromCache(sPlugins))
		}
		if servicePlugins != nil && len(servicePlugins) != 0 {
			service := kong_entity.NewService(
				kong_entity.WithRoutes(routes),
				kong_entity.WithServicePlugins(servicePlugins))
			services := []kong_entity.Service{*service}

			return services
		}
	}

	service := kong_entity.NewService(kong_entity.WithRoutes(routes))
	services := []kong_entity.Service{*service}

	return services
}

func getRoutes(routePluginsFlag string) []kong_entity.Route {
	var routePlugins []kong_entity.Plugin

	if routePluginsFlag != "" {

		routePluginsSplit := strings.Split(routePluginsFlag, ",")
		for _, rPlugins := range routePluginsSplit {
			routePlugins = append(routePlugins, plugins.GetPluginFromCache(rPlugins))
		}

		if routePlugins != nil && len(routePlugins) != 0 {
			route := kong_entity.NewRoute(kong_entity.WithRoutePlugins(routePlugins))
			routes := []kong_entity.Route{*route}

			return routes
		}
	}

	route := kong_entity.NewRoute()
	routes := []kong_entity.Route{*route}

	return routes
}
