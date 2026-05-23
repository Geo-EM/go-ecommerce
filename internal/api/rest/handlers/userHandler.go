package handlers

import (
	"e-commerce/internal/api/rest"
	"e-commerce/internal/api/rest/response"
	"e-commerce/internal/dto/userDto"
	"e-commerce/internal/repository"
	"e-commerce/internal/service"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	userService service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	handler := UserHandler{
		userService: service.UserService{
			UserRepo:     repository.NewUserRepository(rh.DB),
			TokenService: *rh.TokenService,
			AppConfig:    *rh.AppConfig,
		},
	}

	pubRoutes := app.Group("/auth")
	// Public endpoints
	pubRoutes.Post("/register", handler.register)
	pubRoutes.Post("/login", handler.login)

	privRoutes := app.Group("/users", rh.TokenService.Authorize)
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
		return response.BadRequest(ctx, "Invalid credentials")
	}

	token, err := uh.userService.LoginUser(input)
	if err != nil {
		return response.Unauthorized(ctx, "Invalid credentials")
	}

	return response.OK(ctx, fiber.Map{"token": token}, "Logged in successfully")
}

func (uh UserHandler) getVerificationCode(ctx fiber.Ctx) error {
	userClaims, err := uh.userService.TokenService.GetCurrentUser(ctx)
	if err != nil {
		return response.Unauthorized(ctx, err.Error())
	}

	// Create verification code adn send to user
	if err := uh.userService.GetVerificationCode(userClaims.UserID); err != nil {
		return response.BadRequest(ctx, err.Error())
	}

	return response.OK(ctx, nil, "Verification code sent successfully")
}

func (uh UserHandler) verify(ctx fiber.Ctx) error {
	userClaims, err := uh.userService.TokenService.GetCurrentUser(ctx)
	if err != nil {
		return response.Unauthorized(ctx, err.Error())
	}

	input := userDto.VerifyCodeUserDto{}
	if err := ctx.Bind().Body(&input); err != nil {
		return response.BadRequest(ctx, err.Error())
	}

	if err := uh.userService.VerifyUserVerificationCode(userClaims.UserID, input.VerificationCode); err != nil {
		return response.BadRequest(ctx, err.Error())
	}

	return response.OK(ctx, nil, "Verified Successfully")
}

func (uh UserHandler) getProfile(ctx fiber.Ctx) error {
	userClaims, err := uh.userService.TokenService.GetCurrentUser(ctx)
	if err != nil {
		return response.Unauthorized(ctx, err.Error())
	}

	user, err := uh.userService.GetUserProfile(userClaims.UserID)
	if err != nil {
		return response.Unauthorized(ctx, err.Error())
	}

	return response.OK(ctx, fiber.Map{"user": user})
}

func (uh UserHandler) becomeSeller(ctx fiber.Ctx) error {
	userClaims, err := uh.userService.TokenService.GetCurrentUser(ctx)
	if err != nil {
		return response.Unauthorized(ctx, err.Error())
	}

	input := userDto.SellerDto{}
	if err := ctx.Bind().Body(&input); err != nil {
		return response.BadRequest(ctx, "Invalid inputs")
	}

	token, err := uh.userService.BecomeSeller(userClaims.UserID, input)
	if err != nil {
		return response.BadRequest(ctx, err.Error())
	}

	return response.OK(ctx, fiber.Map{"token": token}, "Become seller successfully")
}

func (uh UserHandler) createProfile(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (uh UserHandler) getCart(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (uh UserHandler) updateCart(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (uh UserHandler) getOrders(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}

func (uh UserHandler) getOrderById(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Succeed!",
	})
}
