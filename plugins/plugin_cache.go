package plugins

import "knox/kong_entity"

var pluginCache = map[string]kong_entity.Plugin{
	"key-auth": DefaultKeyAuthPlugin(),
	"cors":     DefaultCORSPlugin(),
}

func GetPluginFromCache(key string) kong_entity.Plugin {
	return pluginCache[key]
}
