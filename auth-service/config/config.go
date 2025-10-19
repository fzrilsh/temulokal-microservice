package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	FrontendOrigin string
	// App
	AppName   string
	AppPort   string
	AppEnv    string
	AppOrigin string

	// Database
	DBHost string
	DBUser string
	DBPass string
	DBName string

	// Email
	SMTPHost     string
	SMTPPort     int
	SMTPPassword string
	SMTPSender   string

	// JWT
	JWTSecret string
}

func Load() *Config {
	_ = godotenv.Load()
	SMTPPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	cfg := &Config{
		FrontendOrigin: os.Getenv("FRONTEND_ORIGIN"),
		AppName:        os.Getenv("APP_NAME"),
		AppPort:        os.Getenv("APP_PORT"),
		AppEnv:         os.Getenv("APP_ENV"),
		AppOrigin:      os.Getenv("APP_ORIGIN"),

		DBHost: os.Getenv("DB_HOST"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),

		SMTPHost:     os.Getenv("SMTP_HOST"),
		SMTPPort:     SMTPPort,
		SMTPPassword: os.Getenv("SMTP_PASSWORD"),
		SMTPSender:   os.Getenv("SMTP_SENDER"),

		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	if cfg.AppPort == "" {
		log.Fatal("APP_PORT is required")
	}

	return cfg
}
