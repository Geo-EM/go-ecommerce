package main

import (
	"e-commerce/config"
	"e-commerce/internal/api"
	"log"
)

func main() {
	conf, err := config.SetupEnv()

	if err != nil {
		log.Fatalf("Error in config loading: %v\n", err)
	}

	api.StartServer(conf)
}
