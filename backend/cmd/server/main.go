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
	database := db.New("transactions.db")
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

