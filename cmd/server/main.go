package main

import (
	"log"
	"github.com/Anvarsha-k/Contact-Management-API/config"

	"github.com/gofiber/fiber/v2"
)

func main() {

	cfg := config.LoadConfig()

	_, err := config.ConnectDatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	// Health check endpoint.
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "server is running",
		})
	})

	log.Fatal(app.Listen(":" + cfg.AppPort))
}