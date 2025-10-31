package deps

import (
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/db"
	"github.com/fadlytanjung/flip-fullstack-test/backend/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

// App holds all application dependencies
type App struct {
	Logger *logger.Logger
	DB     *db.Database
	Fiber  *fiber.App
}

