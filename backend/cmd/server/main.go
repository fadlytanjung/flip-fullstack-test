package main

import (
	"log"
	"os"

	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	// Initialize database
	// Use /tmp for Cloud Run (writable temporary storage)
	// Use local transactions.db for local development
	dbPath := "transactions.db"
	if env := os.Getenv("ENV"); env == "production" {
		dbPath = "/tmp/transactions.db"
	}
	
	database := db.New(dbPath)
	defer database.Close()

	// Create Fiber app
	app := fiber.New()

	// Bootstrap all routes and dependencies
	Bootstrap(app, database)

	// Start server
	log.Printf("Server starting on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}

