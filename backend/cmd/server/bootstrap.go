package main

import (
	"github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/handler"
	transactionRepo "github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/repository"
	"github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/schemas"
	transactionUseCase "github.com/fadlytanjung/flip-fullstack-test/backend/domain/transaction/use_case"
	uploadHandler "github.com/fadlytanjung/flip-fullstack-test/backend/domain/upload/handler"
	uploadRepo "github.com/fadlytanjung/flip-fullstack-test/backend/domain/upload/repository"
	uploadUseCase "github.com/fadlytanjung/flip-fullstack-test/backend/domain/upload/use_case"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Bootstrap initializes all dependencies and routes
func Bootstrap(app *fiber.App, database *db.Database) {
	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Accept,Authorization,Content-Type,X-CSRF-Token",
	}))

	// Auto-migrate database schema
	database.GetDB().AutoMigrate(&schemas.Transaction{})

	// Initialize dependencies for Transaction domain
	transactionRepository := transactionRepo.NewRepository(database.GetDB())
	transactionUC := transactionUseCase.NewUseCase(transactionRepository)
	transactionHdlr := handler.NewHandler(transactionUC)

	// Initialize dependencies for Upload domain
	uploadRepository := uploadRepo.NewRepository()
	uploadUC := uploadUseCase.NewUseCase(uploadRepository, transactionRepository)
	uploadHdlr := uploadHandler.NewHandler(uploadUC)

	// API routes
	api := app.Group("/api")

	// Health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "healthy",
			"service": "flip-fullstack-test-backend",
		})
	})

	// Upload routes
	api.Post("/upload", uploadHdlr.Upload)
	api.Delete("/clear", uploadHdlr.Clear)

	// Transaction routes
	api.Get("/balance", transactionHdlr.GetBalance)
	api.Get("/issues", transactionHdlr.GetIssues)
	api.Get("/transactions", transactionHdlr.GetTransactions)
}