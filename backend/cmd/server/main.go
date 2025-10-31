package main

import (
	"fmt"

	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/config"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/db"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/deps"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// setupMiddleware configures all Fiber middleware
func setupMiddleware(app *fiber.App, l *logger.Logger, cfg *config.GlobalConfig) {
	// Panic recovery with structured logging
	app.Use(logger.RecoveryHandler(l))

	// HTTP request logging
	app.Use(logger.HTTPLogger(l))

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.CorsAllowOrigins,
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Accept,Authorization,Content-Type,X-CSRF-Token",
	}))
}

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	// Initialize logger
	var appLogger *logger.Logger
	
	// Configure logger options
	loggerOpts := []func(*logger.LoggerBuilderOption){}
	
	// Add UDP syncer if configured
	if cfg.LogHostIP != "" && cfg.LogHostPort > 0 {
		loggerOpts = append(loggerOpts, logger.WithUdpSyncer(cfg.LogHostIP, cfg.LogHostPort))
	}
	
	// Add pretty print for development
	if cfg.Environment != "production" {
		loggerOpts = append(loggerOpts, logger.WithPrettyPrint())
	}
	
	appLogger = logger.NewLogger(cfg.ServiceName, cfg.LogLevel, loggerOpts...)
	defer appLogger.Sync()

	appLogger.Info("Starting application",
		logger.String("service", cfg.ServiceName),
		logger.String("version", cfg.ServiceVersion),
		logger.String("environment", cfg.Environment),
		logger.Int("port", cfg.Port),
	)

	// Initialize database
	database := db.New(cfg.DatabasePath)
	defer database.Close()

	// Create Fiber app with middleware
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	// Setup middleware
	setupMiddleware(app, appLogger, cfg)

	// Create app dependencies
	instance := &deps.App{
		Logger: appLogger,
		DB:     database,
		Fiber:  app,
	}

	// Bootstrap application (register routes only)
	BootstrapApp(instance)

	// Start server
	appLogger.Info("Server starting", logger.Int("port", cfg.Port))
	if err := app.Listen(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		appLogger.Fatal("Failed to start server", logger.Error(err))
	}
}
