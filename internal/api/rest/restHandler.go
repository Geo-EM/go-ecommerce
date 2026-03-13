package rest

import (
	"e-commerce/internal/auth"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type RestHandler struct {
	App          *fiber.App
	DB           *gorm.DB
	TokenService *auth.TokenService
}
