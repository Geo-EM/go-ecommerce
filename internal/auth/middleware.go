package auth

import (
	"e-commerce/internal/api/rest/response"
	"e-commerce/internal/domain"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func (t *TokenService) Authorize(ctx fiber.Ctx) error {
	claims, err := t.mainAuthorize(ctx)
	if err != nil {
		return err
	}

	if claims.UserID > 0 {
		ctx.Locals("user", claims)
		return ctx.Next()
	}

	return response.Unauthorized(ctx, "Unauthorized, consider logging in again")
}

func (t *TokenService) AuthorizeSeller(ctx fiber.Ctx) error {

	claims, err := t.mainAuthorize(ctx)
	if err != nil {
		return err
	}

	if claims.UserID > 0 && claims.Role == domain.UserType.SELLER {
		ctx.Locals("user", claims)
		return ctx.Next()
	}

	if claims.UserID > 0 && claims.Role == domain.UserType.BUYER {
		return response.Unauthorized(ctx, "Unauthorized, please join seller program first")
	}

	return response.Unauthorized(ctx, "Unauthorized, consider logging in again")
}

func (t *TokenService) mainAuthorize(ctx fiber.Ctx) (*JwtClaims, error) {
	authHeader := ctx.Get("Authorization")

	if authHeader == "" {
		return nil, response.Unauthorized(ctx, "Unauthorized, consider logging in again")
	}

	parts := strings.SplitN(authHeader, " ", 2)

	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, response.Unauthorized(ctx, "Unauthorized, consider logging in again")
	}

	tokenString := parts[1]

	claims, err := t.ValidateToken(tokenString)

	if err != nil {
		return nil, response.Unauthorized(ctx, "Unauthorized, consider logging in again")
	}

	return claims, nil
}
