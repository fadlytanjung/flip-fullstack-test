package logger

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// HTTPLogger creates a Fiber middleware for HTTP request logging
func HTTPLogger(l *Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := c.Context().Time()

		// Process request
		err := c.Next()

		// Calculate duration
		duration := c.Context().Time().Sub(start)

		// Build log fields
		fields := []zap.Field{
			String("method", c.Method()),
			String("path", c.Path()),
			String("ip", c.IP()),
			Int("status", c.Response().StatusCode()),
			zap.Duration("latency", duration),
			String("user_agent", c.Get("User-Agent")),
		}

		// Add query params if present
		if queryString := string(c.Context().QueryArgs().QueryString()); queryString != "" {
			fields = append(fields, String("query", queryString))
		}

		// Add error if present
		if err != nil {
			fields = append(fields, Error(err))
		}

		// Log based on status code
		status := c.Response().StatusCode()
		if err != nil || status >= 500 {
			l.Error("Request failed", fields...)
		} else if status >= 400 {
			l.Warn("Client error", fields...)
		} else {
			l.Info("Request completed", fields...)
		}

		return err
	}
}

// RecoveryHandler creates a Fiber middleware for panic recovery
func RecoveryHandler(l *Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				l.Error("Panic recovered",
					String("method", c.Method()),
					String("path", c.Path()),
					Any("panic", r),
					zap.Stack("stack"),
				)

				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":  500,
					"message": "Internal server error",
					"error":   "An unexpected error occurred",
				})
			}
		}()

		return c.Next()
	}
}

