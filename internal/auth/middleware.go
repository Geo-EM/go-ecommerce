package auth

import (
	"e-commerce/internal/api/rest/response"
	"e-commerce/internal/domain"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func resUnAuth(ctx fiber.Ctx) error {
	return response.Unauthorized(ctx, "Unauthorized, consider logging in again")
}

func (t *TokenService) Authorize(ctx fiber.Ctx) error {
	claims, err := t.mainAuthorize(ctx)
	if err != nil || claims == nil {
		return resUnAuth(ctx)
	}

	if claims.UserID > 0 {
		ctx.Locals("user", claims)
		return ctx.Next()
	}

	return resUnAuth(ctx)
}

func (t *TokenService) AuthorizeSeller(ctx fiber.Ctx) error {

	claims, err := t.mainAuthorize(ctx)
	if err != nil || claims == nil {
		return resUnAuth(ctx)
	}

	if claims.UserID > 0 && claims.Role == domain.UserType.SELLER {
		ctx.Locals("user", claims)
		return ctx.Next()
	}

	if claims.UserID > 0 && claims.Role == domain.UserType.BUYER {
		return response.Unauthorized(ctx, "Unauthorized, please join seller program first")
	}

	return resUnAuth(ctx)
}

func (t *TokenService) mainAuthorize(ctx fiber.Ctx) (*JwtClaims, error) {
	authHeader := ctx.Get("Authorization")

	if authHeader == "" {
		return nil, resUnAuth(ctx)
	}

	parts := strings.SplitN(authHeader, " ", 2)

	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, resUnAuth(ctx)
	}

	tokenString := parts[1]

	claims, err := t.ValidateToken(tokenString)

	if err != nil || claims == nil {
		return nil, resUnAuth(ctx)
	}

	return claims, nil
}
