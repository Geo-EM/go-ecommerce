package handlers

import (
	"e-commerce/internal/api/rest"
	"e-commerce/internal/dto/userDto"
	"e-commerce/internal/service"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	userService service.UserService
}

func SetupUserRoutes(restHandler *rest.RestHandler) {
	app := restHandler.App

	userService := service.UserService{}
	handler := UserHandler{
		userService: userService,
	}

	// Public endpoints
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	// Private endpoints
	app.Post("/verify", handler.verify)
	app.Get("/verify", handler.GetVerificationCode)

	app.Post("/profile", handler.CreateProfile)
	app.Get("/profile", handler.GetProfile)

	app.Post("/cart", handler.UpdateCart)
	app.Get("/cart", handler.GetCart)

	app.Get("/order", handler.GetOrders)
	app.Get("/order/:id", handler.GetOrderById)

	app.Post("/become-seller", handler.BecomeSeller)

}

func (handler UserHandler) Register(ctx fiber.Ctx) error {
	userInput := userDto.RegisterUserDto{}
	if err := ctx.Bind().Body(&userInput); err != nil {
		log.Println(userInput)
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid inputs",
		})
	}

	token, err := handler.userService.RegisterUser(userInput)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Failed to register",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"token": token,
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
