package plugins

import "kong-config-generator/kong_entity"

var pluginCache = map[string]kong_entity.Plugin{
	"key-auth": DefaultKeyAuthPlugin(),
}

func GetPluginFromCache(key string) kong_entity.Plugin {
	return pluginCache[key]
}
