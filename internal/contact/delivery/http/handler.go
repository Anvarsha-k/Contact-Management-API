package http

import (
	"strconv"

	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/dto"
	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/service"
	customvalidator "github.com/Anvarsha-k/Contact-Management-API/internal/contact/validator"
	"github.com/Anvarsha-k/Contact-Management-API/internal/shared/response"
	"github.com/gofiber/fiber/v2"
)

// ContactHandler handles contact HTTP requests
type ContactHandler struct {
	service service.ContactService
}

func NewContactHandler(service service.ContactService) *ContactHandler {

	return &ContactHandler{
		service: service,
	}
}

// CreateContact handles contact creation

// CreateContact godoc
// @Summary Create contact
// @Description Create a new contact
// @Tags Contacts
// @Accept json
// @Produce json
// @Param request body dto.CreateContactRequest true "Contact payload"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Router /contacts [post]
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

	// Validate request fields
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

// ListContacts handles contact listing

// ListContacts godoc
// @Summary List contacts
// @Description Fetch contacts with pagination, filtering and search
// @Tags Contacts
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Page size"
// @Param search query string false "Search keyword"
// @Param status query string false "Contact status"
// @Param sort_by query string false "Sort field"
// @Param order query string false "Sort order"
// @Success 200 {object} response.APIResponse
// @Router /contacts [get]
func (h *ContactHandler) ListContacts(c *fiber.Ctx) error {

	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	if page < 1 || limit < 1 || limit > 100 {

		return c.Status(fiber.StatusBadRequest).JSON(
			response.APIResponse{
				Success: false,
				Message: "invalid pagination values",
			},
		)
	}

	query := dto.ContactListQuery{
		Page:   page,
		Limit:  limit,
		Search: c.Query("search"),
		Status: c.Query("status"),
		SortBy: c.Query("sort_by"),
		Order:  c.Query("order"),
	}

	contacts, err := h.service.ListContacts(c.Context(), query)

	if err != nil {

		return c.Status(
			fiber.StatusInternalServerError,
		).JSON(
			response.APIResponse{
				Success: false,
				Message: "failed to fetch contacts",
			},
		)
	}

	return c.JSON(
		response.APIResponse{
			Success: true,
			Message: "contacts fetched successfully",
			Data:    contacts,
		},
	)
}

// GetContactByID handles fetching single contact.

// GetContactByID godoc
// @Summary Get contact by ID
// @Description Fetch single contact details
// @Tags Contacts
// @Produce json
// @Param id path int true "Contact ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /contacts/{id} [get]
func (h *ContactHandler) GetContactByID(c *fiber.Ctx) error {

	idParam := c.Params("id")

	id, err := strconv.ParseUint(idParam, 10, 32)

	if err != nil {

		return response.ErrorResponse(c, fiber.StatusBadRequest, "invalid contact ID", nil)
	}

	contact, err := h.service.GetContactByID(c.Context(), uint(id))

	if err != nil {

		return response.ErrorResponse(c, fiber.StatusNotFound, err.Error(), nil)
	}

	return response.SuccessResponse(c, fiber.StatusOK, "contact fetched successfully", contact)
}

// handles contact update endpoint

// UpdateContact godoc
// @Summary Update contact
// @Description Update existing contact
// @Tags Contacts
// @Accept json
// @Produce json
// @Param id path int true "Contact ID"
// @Param request body dto.UpdateContactRequest true "Update payload"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /contacts/{id} [put]
func (h *ContactHandler) UpdateContact(
	c *fiber.Ctx,
) error {

	idParam := c.Params("id")

	id, err := strconv.ParseUint(idParam, 10, 32)

	if err != nil {

		return response.ErrorResponse(
			c,
			fiber.StatusBadRequest,
			"invalid contact ID",
			nil,
		)
	}

	var request dto.UpdateContactRequest

	if err := c.BodyParser(&request); err != nil {

		return response.ErrorResponse(
			c,
			fiber.StatusBadRequest,
			"invalid request payload",
			nil,
		)
	}

	// Validate request
	if err := customvalidator.ValidateStruct(request); err != nil {

		return response.ErrorResponse(c, fiber.StatusBadRequest, "validation failed", err.Error())
	}

	contact, err := h.service.UpdateContact(c.Context(), uint(id), request)

	if err != nil {

		statusCode := fiber.StatusBadRequest

		if err.Error() == "contact not found" {
			statusCode = fiber.StatusNotFound
		}

		return response.ErrorResponse(c, statusCode, err.Error(), nil)
	}

	return response.SuccessResponse(c, fiber.StatusOK, "contact updated successfully", contact)
}

//handles contact deletion
// DeleteContact godoc
// @Summary Delete contact
// @Description Soft delete contact
// @Tags Contacts
// @Produce json
// @Param id path int true "Contact ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /contacts/{id} [delete]
func (h *ContactHandler) DeleteContact(c *fiber.Ctx) error {

	idParam := c.Params("id")

	id, err := strconv.ParseUint(idParam, 10, 32)

	if err != nil {

		return response.ErrorResponse(c, fiber.StatusBadRequest, "invalid contact ID", nil)
	}

	err = h.service.DeleteContact(
		c.Context(),
		uint(id),
	)

	if err != nil {

		return response.ErrorResponse(c, fiber.StatusNotFound, err.Error(), nil)
	}

	return response.SuccessResponse(c, fiber.StatusOK, "contact deleted successfully", nil)
}
