package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	FrontendOrigin string
	AppPort        string
	JWTSecret      string

	// services
	AuthServiceOrigin string
	UMKMServiceOrigin string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		FrontendOrigin:    os.Getenv("FRONTEND_ORIGIN"),
		JWTSecret:         os.Getenv("JWT_SECRET"),
		AppPort:           os.Getenv("APP_PORT"),
		AuthServiceOrigin: os.Getenv("AUTH_SERVICE_ORIGIN"),
		UMKMServiceOrigin: os.Getenv("UMKM_SERVICE_ORIGIN"),
	}

	return cfg
}
