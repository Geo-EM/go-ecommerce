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
func RespondSuccess(ctx fiber.Ctx, status int, data fiber.Map, message string) error {
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
func BadRequest(ctx fiber.Ctx, message ...string) error {
	msg := "Bad Request"
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	return RespondError(ctx, fiber.StatusBadRequest, msg)
}

func Unauthorized(ctx fiber.Ctx, message ...string) error {
	msg := "Unauthorized"
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	return RespondError(ctx, fiber.StatusUnauthorized, msg)
}

func NotFound(ctx fiber.Ctx, message ...string) error {
	msg := "Not Found"
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	return RespondError(ctx, fiber.StatusNotFound, msg)
}

// Success responses
func OK(ctx fiber.Ctx, data fiber.Map, message ...string) error {
	msg := "OK"
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	return RespondSuccess(ctx, fiber.StatusOK, data, msg)
}

func Created(ctx fiber.Ctx, data fiber.Map, message ...string) error {
	msg := "Created"
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	return RespondSuccess(ctx, fiber.StatusCreated, data, msg)
}
