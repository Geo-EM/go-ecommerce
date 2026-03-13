package response

import (
	"maps"

	"github.com/gofiber/fiber/v3"
)

// RespondError sends a standardized error JSON response
func RespondError(ctx fiber.Ctx, status int, message string) error {
	return ctx.Status(status).JSON(fiber.Map{
		"status":  "error",
		"message": message,
	})
}

// RespondSuccess sends a standardized success JSON response
// `message` is required, `data` is optional extra fields
func RespondSuccess(ctx fiber.Ctx, status int, message string, data fiber.Map) error {
	res := fiber.Map{
		"status":  "success",
		"message": message,
	}
	if data != nil {
		maps.Copy(res, data)
	}
	return ctx.Status(status).JSON(res)
}

// Shortcut helpers for common HTTP status codes //
// Error responses
func BadRequest(ctx fiber.Ctx, message string) error {
	return RespondError(ctx, fiber.StatusBadRequest, message)
}

func Unauthorized(ctx fiber.Ctx, message string) error {
	return RespondError(ctx, fiber.StatusUnauthorized, message)
}

func NotFound(ctx fiber.Ctx, message string) error {
	return RespondError(ctx, fiber.StatusNotFound, message)
}

// Success responses
func OK(ctx fiber.Ctx, message string, data fiber.Map) error {
	return RespondSuccess(ctx, fiber.StatusOK, message, data)
}

func Created(ctx fiber.Ctx, message string, data fiber.Map) error {
	return RespondSuccess(ctx, fiber.StatusCreated, message, data)
}
