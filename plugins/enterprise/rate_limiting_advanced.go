package enterprise

import "knox/kong_entity"

type RateLimitingAdvancedConfig struct {
	Identifier            string      `yaml:"identifier"`
	WindowSize            []int       `yaml:"window_size"`
	WindowType            string      `yaml:"window_type"`
	Limit                 []int       `yaml:"limit"`
	SyncRate              float64     `yaml:"sync_rate"`
	Namespace             string      `yaml:"namespace"`
	Strategy              string      `yaml:"startegy"`
	DictionaryName        string      `yaml:"dictionary_name"`
	HideClientHeaders     bool        `yaml:"hide_client_headers"`
	RetryAfterJitterMax   int         `yaml:"retry_after_jitter_max"`
	HeaderName            string      `yaml:"header_name"`
	Path                  string      `yaml:"path"`
	RedisConfig           RedisConfig `yaml:"redis"`
	EnforceConsumerGroups bool        `yaml:"enforce_consumer_groups"`
	ConsumerGroups        []string    `yaml:"consumer_groups"`
	DisablePenalty        bool        `yaml:"disable_penalty"`
	ErrorCode             int         `yaml:"error_code"`
	ErrorMessage          string      `yaml:"error_message"`
}

type RedisConfig struct {
	Host              string   `yaml:"host"`
	Port              int      `yaml:"port"`
	Timeout           int      `yaml:"timeout"`
	ConnectTimeout    int      `yaml:"connect_timeout"`
	SendTimeout       int      `yaml:"send_timeout"`
	ReadTimeout       int      `yaml:"read_timeout"`
	Username          string   `yaml:"username"`
	Password          string   `yaml:"password"`
	SentinelUsername  string   `yaml:"sentinel_username"`
	SentinelPassword  string   `yaml:"sentinel_password"`
	Database          int      `yaml:"database"`
	KeepAlivePoolSize int      `yaml:"keepalive_pool_size"`
	KeepAliveBacklog  string   `yaml:"keepalive_backlog"`
	SentinelMaster    string   `yaml:"sentinel_master"`
	SentinelRole      string   `yaml:"sentinel_role"`
	ClusterAddresses  []string `yaml:"cluster_addresses"`
	SSL               bool     `yaml:"ssl"`
	SSLVerify         bool     `yaml:"ssl_verify"`
	ServerName        string   `yaml:"server_name"`
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
