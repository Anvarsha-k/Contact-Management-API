package http

import (

	customvalidator "github.com/Anvarsha-k/Contact-Management-API/internal/contact/validator"
	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/dto"
	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/service"
	"github.com/Anvarsha-k/Contact-Management-API/internal/shared/response"
	"github.com/gofiber/fiber/v2"
)

// ContactHandler handles contact HTTP requests.
type ContactHandler struct {
	service service.ContactService
}

func NewContactHandler( service service.ContactService,) *ContactHandler {

	return &ContactHandler{
		service: service,
	}
}

// CreateContact handles contact creation endpoint.
func (h *ContactHandler) CreateContact(c *fiber.Ctx) error {

	var request dto.CreateContactRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.APIResponse{
				Success: false,
				Message: "invalid request payload",
			},
		)
	}

	// Validate request fields.
	if err := customvalidator.ValidateStruct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.APIResponse{
				Success: false,
				Message: "validation failed",
				Errors:  err.Error(),
			},
		)
	}

	contact, err := h.service.CreateContact(
		c.Context(),
		request,
	)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.APIResponse{
				Success: false,
				Message: err.Error(),
			},
		)
	}

	return c.Status(fiber.StatusCreated).JSON(
		response.APIResponse{
			Success: true,
			Message: "contact created successfully",
			Data:    contact,
		},
	)
}