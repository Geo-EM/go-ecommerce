package handlers

import (
	"e-commerce/internal/api/rest"
	"e-commerce/internal/api/rest/response"
	"e-commerce/internal/dto/userDto"
	"e-commerce/internal/repository"
	"e-commerce/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	userService service.UserService
}

func SetupUserRoutes(restHandler *rest.RestHandler) {
	app := restHandler.App

	handler := UserHandler{
		userService: service.UserService{
			UserRepo:     repository.NewUserRepository(restHandler.DB),
			TokenService: *restHandler.TokenService,
		},
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

func (uh UserHandler) register(ctx fiber.Ctx) error {
	input := userDto.RegisterUserDto{}

	if err := ctx.Bind().Body(&input); err != nil {
		return response.BadRequest(ctx, "Invalid input")
	}

	token, err := uh.userService.RegisterUser(input)
	if err != nil {
		return response.BadRequest(ctx, err.Error())
	}

	return response.Created(ctx, "Registered successfully", fiber.Map{"token": token})
}

func (uh UserHandler) login(ctx fiber.Ctx) error {
	input := userDto.LoginUserDto{}

	if err := ctx.Bind().Body(&input); err != nil {
		return response.BadRequest(ctx, "Invalid input")
	}

	token, err := uh.userService.LoginUser(input)
	if err != nil {
		return response.Unauthorized(ctx, "Invalid credentials")
	}

	return response.OK(ctx, "Logged in successfully", fiber.Map{"token": token})
}

func (uh UserHandler) getVerificationCode(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (uh UserHandler) verify(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (uh UserHandler) getProfile(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (uh UserHandler) createProfile(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (uh UserHandler) getCart(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (uh UserHandler) updateCart(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (uh UserHandler) getOrders(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (uh UserHandler) getOrderById(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (uh UserHandler) becomeSeller(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}
