package config

import (
	"os"

	"github.com/spf13/viper"
)

// loadDefaults sets CONCRETE default values for local development
// In production, these MUST be overridden by environment variables
func loadDefaults() {
	// Environment config - CONCRETE defaults for local development
	viper.SetDefault("ENV", "development")           // Production MUST set ENV=production
	viper.SetDefault("DEBUG", false)
	viper.SetDefault("PORT", 9000)                   // Cloud Run provides PORT env var
	viper.SetDefault("LOG_LEVEL", "info")            // Production can override to "debug" if needed

	// Service metadata
	viper.SetDefault("SERVICE_NAME", "flip-fullstack-test-backend")
	
	// Read version from file if exists
	versionBytes, err := os.ReadFile("VERSION")
	if err == nil {
		// Remove trailing newline
		version := string(versionBytes)
		if len(version) > 0 && version[len(version)-1] == '\n' {
			version = version[:len(version)-1]
		}
		viper.SetDefault("SERVICE_VERSION", version)
	} else {
		viper.SetDefault("SERVICE_VERSION", "1.0.0")
	}

	// Database config
	viper.SetDefault("DATABASE_PATH", "transactions.db")  // Production overrides to /tmp/transactions.db

	// Logging config (optional for UDP shipping)
	viper.SetDefault("LOG_HOST_IP", "")
	viper.SetDefault("LOG_HOST_PORT", 0)

	// CORS config
	viper.SetDefault("CORS_ALLOW_ORIGINS", "*")           // Production can set specific origins
}



