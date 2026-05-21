package main

import (
	"log"

	contactHttp "github.com/Anvarsha-k/Contact-Management-API/internal/contact/delivery/http"

	"github.com/Anvarsha-k/Contact-Management-API/config"
	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/repository"
	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/service"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// Load config
	cfg := config.LoadConfig()

	// Initialize db connection
	db, err := config.ConnectDatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}

	
	app := fiber.New()

	
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "server is running",
		})
	})

	// Dependency injection
	contactRepository := repository.NewContactRepository(db)

	contactService := service.NewContactService(
		contactRepository,
	)

	contactHandler := contactHttp.NewContactHandler(
		contactService,
	)

	// API routes
	api := app.Group("/api/v1")

	contactHttp.RegisterContactRoutes(
		api,
		contactHandler,
	)

	
	log.Fatal(app.Listen(":" + cfg.AppPort))
}