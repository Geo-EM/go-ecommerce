package api

import (
	"e-commerce/config"
	"log"

	"github.com/gofiber/fiber/v3"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	if err := app.Listen("localhost:" + config.ServerPort); err != nil {
		log.Println("Error starting server:", err)
		return
	}
	log.Println("Server is running on port 9898")
}
