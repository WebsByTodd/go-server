package config

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

type Application struct {
	JWTSecret string
}

func ReadProperties() Application {
	err := godotenv.Load()
	if err != nil {
		slog.Error("error loading .env file")
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	return Application{
		JWTSecret: jwtSecret,
	}
}
