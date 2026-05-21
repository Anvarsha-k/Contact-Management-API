package http

import "github.com/gofiber/fiber/v2"

func RegisterContactRoutes( api fiber.Router, handler *ContactHandler,) {

	contacts := api.Group("/contacts")

	contacts.Post("/", handler.CreateContact)
	contacts.Get("/", handler.ListContacts)
	contacts.Get("/:id", handler.GetContactByID)
	contacts.Put("/:id", handler.UpdateContact)
	contacts.Delete("/:id", handler.DeleteContact)
}