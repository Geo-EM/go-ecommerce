package auth

import (
	"strings"

	"github.com/gofiber/fiber/v3"
)

func (t TokenService) Authorize(ctx fiber.Ctx) error {

	authHeader := ctx.Get("Authorization")

	if authHeader == "" {
		return fiber.ErrUnauthorized
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := t.ValidateToken(token)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	ctx.Locals("userID", claims.UserID)
	ctx.Locals("role", claims.Role)

	return ctx.Next()
}
