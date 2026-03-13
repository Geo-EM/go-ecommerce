package auth

import (
	"strings"

	"github.com/gofiber/fiber/v3"
)

func (t *TokenService) Authorize(ctx fiber.Ctx) error {

	authHeader := ctx.Get("Authorization")

	if authHeader == "" {
		return fiber.ErrUnauthorized
	}

	parts := strings.SplitN(authHeader, " ", 2)

	if len(parts) != 2 || parts[0] != "Bearer" {
		return fiber.ErrUnauthorized
	}

	tokenString := parts[1]

	claims, err := t.ValidateToken(tokenString)

	if err == nil && claims.UserID > 0 {
		ctx.Locals("user", claims)
		return ctx.Next()
	}

	return fiber.ErrUnauthorized
}
