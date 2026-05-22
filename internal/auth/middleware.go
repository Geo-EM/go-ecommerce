package auth

import (
	"e-commerce/internal/api/rest/response"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func (t *TokenService) Authorize(ctx fiber.Ctx) error {

	authHeader := ctx.Get("Authorization")

	if authHeader == "" {
		return response.Unauthorized(ctx, "Unauthorized, consider logging in again")
	}

	parts := strings.SplitN(authHeader, " ", 2)

	if len(parts) != 2 || parts[0] != "Bearer" {
		return response.Unauthorized(ctx, "Unauthorized, consider logging in again")
	}

	tokenString := parts[1]

	claims, err := t.ValidateToken(tokenString)

	if err == nil && claims.UserID > 0 {
		ctx.Locals("user", claims)
		return ctx.Next()
	}

	return response.Unauthorized(ctx, "Unauthorized, consider logging in again")
}
