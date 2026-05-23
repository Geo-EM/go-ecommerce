package auth

import (
	"errors"

	"github.com/gofiber/fiber/v3"
)

func (t *TokenService) GetCurrentUser(ctx fiber.Ctx) (*JwtClaims, error) {
	user := ctx.Locals("user")
	if user == nil {
		return nil, errors.New("Unauthorized, consider logging in again")
	}

	claims, ok := user.(*JwtClaims)
	if ok && claims.UserID > 0 {
		return claims, nil
	}

	return nil, errors.New("Unauthorized, consider logging in again")

}
