package plugins

import (
	"knox/kong_entity"
	"knox/plugins/enterprise"
)

var pluginCache = map[string]kong_entity.Plugin{
	"key-auth":                     DefaultKeyAuthPlugin(),
	"cors":                         DefaultCORSPlugin(),
	"request-transformer-advanced": enterprise.DefaultRequestTransformerAdvancedPlugin(),
	"exit-transformer":             enterprise.DefaultExitTransformerPlugin(),
	"pre-function":                 DefaultPreFunctionPlugin(),
	"rate-limiting-advanced":       enterprise.DefaultRateLimitingAdvancedPlugin(),
	"acl":                          DefaultACLPlugin(),
}

func GetPluginFromCache(key string) kong_entity.Plugin {
	return pluginCache[key]
}
