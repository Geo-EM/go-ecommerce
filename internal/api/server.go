package api

import (
	"e-commerce/config"
	"e-commerce/internal/api/rest"
	"e-commerce/internal/api/rest/handlers"
	"log"

	"github.com/gofiber/fiber/v3"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	setupRoutes(&rest.RestHandler{App: app})

	if err := app.Listen("localhost:" + config.ServerPort); err != nil {
		log.Fatalln("Error starting server:", err)
	}
}

func setupRoutes(restHandler *rest.RestHandler) {
	handlers.SetupHealthCheckRoute(restHandler)
	handlers.SetupUserRoutes(restHandler)
}
