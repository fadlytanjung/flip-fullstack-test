package config

import (
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

var globalConfig *GlobalConfig

// LoadConfig loads configuration from .env file and environment variables
// Priority: Environment Variables > .env file > Defaults
func LoadConfig() (*GlobalConfig, error) {
	// Load defaults first
	loadDefaults()

	// Load from .env file (if exists)
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	// Enable environment variable binding
	viper.AutomaticEnv()

	// Bind all config fields to environment variables
	emptyConfig := &GlobalConfig{}
	fieldCount := reflect.TypeOf(emptyConfig).Elem().NumField()

	for i := 0; i < fieldCount; i++ {
		// Get the mapstructure tag name (e.g., "ENV", "PORT", etc.)
		field := reflect.TypeOf(emptyConfig).Elem().Field(i)
		tag := field.Tag.Get("mapstructure")
		
		if tag != "" {
			// Bind environment variable
			if err := viper.BindEnv(tag); err != nil {
				fmt.Printf("Error binding env var %s: %v\n", tag, err)
			}
		}
	}

	// Read config from .env file (ignore error if file doesn't exist)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Config file found but has errors
			fmt.Printf("Error reading .env file: %v\n", err)
		}
		// If file not found, continue (will use env vars and defaults)
	}

	// Unmarshal into config struct
	config := &GlobalConfig{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Override database path for production environment
	if config.Environment == "production" && config.DatabasePath == "transactions.db" {
		config.DatabasePath = "/tmp/transactions.db"
	}

	// Store globally
	globalConfig = config

	return config, nil
}

// GetConfig returns the loaded global configuration
func GetConfig() *GlobalConfig {
	if globalConfig == nil {
		panic("config not loaded - call LoadConfig() first")
	}
	return globalConfig
}



