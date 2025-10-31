package main

import (
	transactionHandler "github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/handler"
	"github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/schemas"
	uploadHandler "github.com/fadlytanjung/flip-fullstack-test/backend/domain/upload/handler"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/config"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/deps"
	"github.com/gofiber/fiber/v2"
)

// BootstrapApp registers all routes and domain handlers
func BootstrapApp(d *deps.App) *deps.App {
	cfg := config.GetConfig()

	// Auto-migrate database schema
	d.DB.GetDB().AutoMigrate(&schemas.Transaction{})

	// Health check
	d.Fiber.Get("/api/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":      "healthy",
			"service":     cfg.ServiceName,
			"version":     cfg.ServiceVersion,
			"environment": cfg.Environment,
		})
	})

	// Register domain APIs
	transactionHandler.RegisterApi(d)
	uploadHandler.RegisterApi(d)

	return d
}
