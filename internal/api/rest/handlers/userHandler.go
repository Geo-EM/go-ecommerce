package handlers

import (
	"e-commerce/internal/api/rest"
	"e-commerce/internal/dto/userDto"
	"e-commerce/internal/repository"
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

	handler := UserHandler{
		userService: service.UserService{UserRepo: repository.NewUserRepository(restHandler.DB)},
	}

	// Public endpoints
	app.Post("/register", handler.register)
	app.Post("/login", handler.login)

	// Private endpoints
	app.Post("/verify", handler.verify)
	app.Get("/verify", handler.getVerificationCode)

	app.Post("/profile", handler.createProfile)
	app.Get("/profile", handler.getProfile)

	app.Post("/cart", handler.updateCart)
	app.Get("/cart", handler.getCart)

	app.Get("/order", handler.getOrders)
	app.Get("/order/:id", handler.getOrderById)

	app.Post("/become-seller", handler.becomeSeller)

}

func (handler UserHandler) register(ctx fiber.Ctx) error {
	userInput := userDto.RegisterUserDto{}
	if err := ctx.Bind().Body(&userInput); err != nil {
		log.Println(userInput)
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid inputs",
		})
	}

	token, err := handler.userService.RegisterUser(userInput)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"token": token,
	})
}

func (handler UserHandler) login(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Logged in",
	})
}

func (handler UserHandler) getVerificationCode(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (handler UserHandler) verify(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (handler UserHandler) getProfile(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (handler UserHandler) createProfile(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (handler UserHandler) getCart(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (handler UserHandler) updateCart(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (handler UserHandler) getOrders(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (handler UserHandler) getOrderById(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (handler UserHandler) becomeSeller(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}
