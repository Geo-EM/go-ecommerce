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

	pubRoutes := app.Group("/auth")
	// Public endpoints
	pubRoutes.Post("/register", handler.register)
	pubRoutes.Post("/login", handler.login)

	privRoutes := app.Group("/users", restHandler.TokenService.Authorize)
	// Private endpoints
	privRoutes.Post("/verify", handler.verify)
	privRoutes.Get("/verify", handler.getVerificationCode)

	privRoutes.Post("/profile", handler.createProfile)
	privRoutes.Get("/profile", handler.getProfile)

	privRoutes.Post("/cart", handler.updateCart)
	privRoutes.Get("/cart", handler.getCart)

	privRoutes.Get("/order", handler.getOrders)
	privRoutes.Get("/order/:id", handler.getOrderById)

	privRoutes.Post("/become-seller", handler.becomeSeller)

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

	return response.Created(ctx, fiber.Map{"token": token}, "Registered successfully")
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

	return response.OK(ctx, fiber.Map{"token": token}, "Logged in successfully")
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
	userClaims, err := uh.userService.TokenService.GetCurrentUser(ctx)
	if err != nil {
		return response.Unauthorized(ctx, "Unauthorized, consider logging in again")
	}

	user, err := uh.userService.GetUserProfile(userClaims.UserID)
	if err != nil {
		return response.Unauthorized(ctx, "Unauthorized, consider logging in again")
	}

	return response.RespondSuccess(ctx, http.StatusOK, fiber.Map{"user": user})
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
