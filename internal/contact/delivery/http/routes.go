package http

import "github.com/gofiber/fiber/v2"

func RegisterContactRoutes( api fiber.Router, handler *ContactHandler,) {

	contacts := api.Group("/contacts")

	contacts.Post("/", handler.CreateContact)
}