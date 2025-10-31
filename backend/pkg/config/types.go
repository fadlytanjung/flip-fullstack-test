package config

type (
	// GlobalConfig holds all application configuration
	GlobalConfig struct {
		// Environment config
		Environment string `mapstructure:"ENV"`
		Debug       bool   `mapstructure:"DEBUG"`
		Port        int    `mapstructure:"PORT"`
		LogLevel    string `mapstructure:"LOG_LEVEL"`

		// Service metadata
		ServiceName    string `mapstructure:"SERVICE_NAME"`
		ServiceVersion string `mapstructure:"SERVICE_VERSION"`

		// Database config
		DatabasePath string `mapstructure:"DATABASE_PATH"`

		// Logging config
		LogHostIP   string `mapstructure:"LOG_HOST_IP"`
		LogHostPort int    `mapstructure:"LOG_HOST_PORT"`

		// CORS config
		CorsAllowOrigins string `mapstructure:"CORS_ALLOW_ORIGINS"`
	}
)



