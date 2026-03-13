package api

import (
	"e-commerce/config"
	"e-commerce/internal/api/rest"
	"e-commerce/internal/api/rest/handlers"
	"e-commerce/internal/auth"
	"e-commerce/internal/domain"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	// Initialize database connection and auto-migrate models
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connecting to database:", err)
	}
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatalln("Error migrating database:", err)
	}

	// Initialize token service with JWT configuration
	tokenService, err := auth.NewTokenService(
		config.JWTSecret,
		config.JWTIssuer,
		24*time.Hour,
	)
	if err != nil {
		log.Fatalln("Error initializing token service:", err)
	}

	setupRoutes(&rest.RestHandler{
		App:          app,
		DB:           db,
		TokenService: tokenService,
	})

	if err := app.Listen(config.Host + ":" + config.ServerPort); err != nil {
		log.Fatalln("Error starting server:", err)
	}
}

func setupRoutes(restHandler *rest.RestHandler) {
	handlers.SetupHealthCheckRoute(restHandler)
	handlers.SetupUserRoutes(restHandler)
}
