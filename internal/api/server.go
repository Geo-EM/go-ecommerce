package api

import (
	"e-commerce/config"
	"e-commerce/internal/api/rest"
	"e-commerce/internal/api/rest/handlers"
	"e-commerce/internal/domain"
	"log"

	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connecting to database:", err)
	}

	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatalln("Error migrating database:", err)
	}

	setupRoutes(&rest.RestHandler{App: app, DB: db})

	if err := app.Listen(config.Host + ":" + config.ServerPort); err != nil {
		log.Fatalln("Error starting server:", err)
	}
	log.Printf("Starting server on %s:%s\n", config.Host, config.ServerPort)
}

func setupRoutes(restHandler *rest.RestHandler) {
	handlers.SetupHealthCheckRoute(restHandler)
	handlers.SetupUserRoutes(restHandler)
}
