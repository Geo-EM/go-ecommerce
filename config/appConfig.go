package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Host       string
	ServerPort string
	Dsn        string

	JWTSecret string
	JWTIssuer string
}

func SetupEnv() (conf AppConfig, err error) {
	if os.Getenv("APP_ENV") == "dev" {
		if err := godotenv.Load(); err != nil {
			return AppConfig{}, errors.New("error loading .env file")
		}
	}

	// Server
	host := os.Getenv("HOST")
	if len(host) < 1 {
		return AppConfig{}, errors.New("HOST env variable not specified")
	}
	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("HTTP_PORT env variable not specified")
	}

	// DB
	dsn := os.Getenv("DSN")
	if len(dsn) < 1 {
		return AppConfig{}, errors.New("DSN env variable not specified")
	}

	// JWT
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return AppConfig{}, errors.New("JWT_SECRET env variable not specified")
	}
	jwtIssuer := os.Getenv("JWT_ISSUER")
	if jwtIssuer == "" {
		return AppConfig{}, errors.New("JWT_ISSUER env variable not specified")
	}

	return AppConfig{
		Host:       host,
		ServerPort: httpPort,

		Dsn: dsn,

		JWTSecret: jwtSecret,
		JWTIssuer: jwtIssuer,
	}, nil
}
