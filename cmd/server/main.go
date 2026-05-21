package main

import (
	"log"

	"github.com/Anvarsha-k/Contact-Management-API/config"

	_ "github.com/Anvarsha-k/Contact-Management-API/docs"

	contactHttp "github.com/Anvarsha-k/Contact-Management-API/internal/contact/delivery/http"
	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/repository"
	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/service"
	"github.com/Anvarsha-k/Contact-Management-API/internal/middleware"

	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"
)

// @title Contact Management API
// @version 1.0
// @description REST API for managing contacts.
// @host localhost:8080
// @BasePath /api/v1
func main() {

	// Load application configuration.
	cfg := config.LoadConfig()

	// Initialize database connection.
	db, err := config.ConnectDatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Fiber application.
	app := fiber.New(
		fiber.Config{
			ErrorHandler: func(
				c *fiber.Ctx,
				err error,
			) error {

				code := fiber.StatusInternalServerError

				if e, ok := err.(*fiber.Error); ok {
					code = e.Code
				}

				return c.Status(code).JSON(
					fiber.Map{
						"success": false,
						"message": err.Error(),
					},
				)
			},
		},
	)

	// Register global middlewares.
	middleware.Register(app)

	// Health check endpoint.
	app.Get("/health", func(c *fiber.Ctx) error {

		return c.JSON(
			fiber.Map{
				"success": true,
				"message": "server is running",
			},
		)
	})

	// Swagger documentation route.
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Dependency injection.
	contactRepository := repository.NewContactRepository(db)

	contactService := service.NewContactService(
		contactRepository,
	)

	contactHandler := contactHttp.NewContactHandler(
		contactService,
	)

	// API route group.
	api := app.Group("/api/v1")

	contactHttp.RegisterContactRoutes(
		api,
		contactHandler,
	)

	// Start HTTP server.
	log.Fatal(app.Listen(":" + cfg.AppPort))
}