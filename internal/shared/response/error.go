package response

import "github.com/gofiber/fiber/v2"

// ITs sends standardized error response
func ErrorResponse(c *fiber.Ctx, statusCode int, message string, errors interface{}) error {

	return c.Status(statusCode).JSON(
		APIResponse{
			Success: false,
			Message: message,
			Errors:  errors,
		},
	)
}