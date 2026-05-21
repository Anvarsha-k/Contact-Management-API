package response

import "github.com/gofiber/fiber/v2"

// It sends standardized success response
func SuccessResponse(c *fiber.Ctx, statusCode int, message string, data interface{},) error {

	return c.Status(statusCode).JSON(
		APIResponse{
			Success: true,
			Message: message,
			Data:    data,
		},
	)
}