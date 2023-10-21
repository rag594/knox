package enterprise

import "knox/kong_entity"

type RateLimitingAdvancedConfig struct {
	Identifier            string      `json:"identifier" yaml:"identifier"`
	WindowSize            []int       `json:"window_size" yaml:"window_size"`
	WindowType            string      `json:"window_type" yaml:"window_type"`
	Limit                 []int       `json:"limit" yaml:"limit"`
	SyncRate              float64     `json:"sync_rate" yaml:"sync_rate"`
	Namespace             string      `json:"namespace" yaml:"namespace"`
	Strategy              string      `json:"startegy" yaml:"startegy"`
	DictionaryName        string      `json:"dictionary_name" yaml:"dictionary_name"`
	HideClientHeaders     bool        `json:"hide_client_headers" yaml:"hide_client_headers"`
	RetryAfterJitterMax   int         `json:"retry_after_jitter_max" yaml:"retry_after_jitter_max"`
	HeaderName            string      `json:"header_name" yaml:"header_name"`
	Path                  string      `json:"path" yaml:"path"`
	RedisConfig           RedisConfig `json:"redis" yaml:"redis"`
	EnforceConsumerGroups bool        `json:"enforce_consumer_groups" yaml:"enforce_consumer_groups"`
	ConsumerGroups        []string    `json:"consumer_groups" yaml:"consumer_groups"`
	DisablePenalty        bool        `json:"disable_penalty" yaml:"disable_penalty"`
	ErrorCode             int         `json:"error_code" yaml:"error_code"`
	ErrorMessage          string      `json:"error_message" yaml:"error_message"`
}

type RedisConfig struct {
	Host              string   `json:"host" yaml:"host"`
	Port              int      `json:"port" yaml:"port"`
	Timeout           int      `json:"timeout" yaml:"timeout"`
	ConnectTimeout    int      `json:"connect_timeout" yaml:"connect_timeout"`
	SendTimeout       int      `json:"send_timeout" yaml:"send_timeout"`
	ReadTimeout       int      `json:"read_timeout" yaml:"read_timeout"`
	Username          string   `json:"username" yaml:"username"`
	Password          string   `json:"password" yaml:"password"`
	SentinelUsername  string   `json:"sentinel_username" yaml:"sentinel_username"`
	SentinelPassword  string   `json:"sentinel_password" yaml:"sentinel_password"`
	Database          int      `json:"database" yaml:"database"`
	KeepAlivePoolSize int      `json:"keepalive_pool_size" yaml:"keepalive_pool_size"`
	KeepAliveBacklog  string   `json:"keepalive_backlog" yaml:"keepalive_backlog"`
	SentinelMaster    string   `json:"sentinel_master" yaml:"sentinel_master"`
	SentinelRole      string   `json:"sentinel_role" yaml:"sentinel_role"`
	ClusterAddresses  []string `json:"cluster_addresses" yaml:"cluster_addresses"`
	SSL               bool     `json:"ssl" yaml:"ssl"`
	SSLVerify         bool     `json:"ssl_verify" yaml:"ssl_verify"`
	ServerName        string   `json:"server_name" yaml:"server_name"`
}

func defaulltRateLimitingAdvancedConfig() RateLimitingAdvancedConfig {
	return RateLimitingAdvancedConfig{
		Identifier: "consumer",
		WindowSize: []int{60},
		WindowType: "fixed",
		SyncRate:   0.02,
		Limit:      []int{100},
		RedisConfig: RedisConfig{
			Host:     "redis://abc",
			Port:     6379,
			Database: 0,
		},
	}
}

func DefaultRateLimitingAdvancedPlugin() kong_entity.Plugin {
	return kong_entity.Plugin{
		Config:   defaulltRateLimitingAdvancedConfig(),
		Enabled:  true,
		Name:     "rate-limiting-advanced",
		Protcols: []string{"http", "https"},
	}
}
