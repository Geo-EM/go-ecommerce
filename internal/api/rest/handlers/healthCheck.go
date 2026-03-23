package handlers

import (
	"e-commerce/internal/api/rest"

	"github.com/gofiber/fiber/v3"
)

func SetupHealthCheckRoute(restHandler *rest.RestHandler) {
	app := restHandler.App

	app.Get("/health-check", handler)
}

func handler(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Working!",
	})
}
