package auth

import (
	"github.com/gofiber/fiber/v3"
)

func (t *TokenService) GetCurrentUser(ctx fiber.Ctx) (*JwtClaims, error) {
	user := ctx.Locals("user")
	if user == nil {
		return nil, fiber.ErrUnauthorized
	}

	claims, ok := user.(*JwtClaims)
	if !ok || claims.UserID <= 0 {
		return nil, fiber.ErrUnauthorized
	}

	return claims, nil
}
