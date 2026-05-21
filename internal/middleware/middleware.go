package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Register global middlewares
func Register(app *fiber.App) {

	// Recover from panics
	app.Use(recover.New())

	// Request logging middleware
	app.Use(fiberLogger.New())

	// Rate limiting middleware
	app.Use(
		limiter.New(
			limiter.Config{
				Max:        100,
				Expiration: time.Minute,
			},
		),
	)
}