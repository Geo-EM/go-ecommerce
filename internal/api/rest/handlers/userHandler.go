package handlers

import (
	"e-commerce/internal/api/rest"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	// Service UserService
}

func SetupUserRoutes(restHandler *rest.RestHandler) {
	app := restHandler.App

	handler := UserHandler{}

	// Public endpoints
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	// Private endpoints
	app.Get("/verify", handler.GetVerificationCode)
	app.Post("/verify", handler.verify)

	app.Get("/profile", handler.GetProfile)
	app.Post("/profile", handler.CreateProfile)

	app.Get("/cart", handler.GetCart)
	app.Post("/cart", handler.UpdateCart)

	app.Get("/order", handler.GetOrders)
	app.Get("/order/:id", handler.GetOrderById)

	app.Post("/become-seller", handler.BecomeSeller)

}

func (handler UserHandler) Register(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Registered",
	})
}

func (handler UserHandler) Login(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Logged in",
	})
}

func (handler UserHandler) GetVerificationCode(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (handler UserHandler) verify(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (handler UserHandler) GetProfile(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (handler UserHandler) CreateProfile(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (handler UserHandler) GetCart(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (handler UserHandler) UpdateCart(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (handler UserHandler) GetOrders(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (handler UserHandler) GetOrderById(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (handler UserHandler) BecomeSeller(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}
