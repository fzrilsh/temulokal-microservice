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
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	JWTSecret string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		AppName:   os.Getenv("APP_NAME"),
		AppPort:   os.Getenv("APP_PORT"),
		AppEnv:    os.Getenv("APP_ENV"),
		AppOrigin: os.Getenv("APP_ORIGIN"),
		DBHost:    os.Getenv("DB_HOST"),
		DBUser:    os.Getenv("DB_USER"),
		DBPass:    os.Getenv("DB_PASS"),
		DBName:    os.Getenv("DB_NAME"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	if cfg.AppPort == "" {
		log.Fatal("APP_PORT is required")
	}

	return cfg
}
