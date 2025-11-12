package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName   string
	AppPort   string
	AppEnv    string
	AppOrigin string

	DBHost string
	DBUser string
	DBPass string
	DBName string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		AppName:   os.Getenv("UMKM_APP_NAME"),
		AppPort:   os.Getenv("UMKM_APP_PORT"),
		AppEnv:    os.Getenv("APP_ENV"),
		AppOrigin: os.Getenv("APP_ORIGIN"),

		DBHost: os.Getenv("DB_HOST"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),
	}

	if cfg.AppPort == "" {
		log.Fatal("UMKM_APP_PORT is required")
	}

	return cfg
}
