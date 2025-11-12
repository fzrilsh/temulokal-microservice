package bootstrap

import (
	"fmt"
	"temulokal-microservice/gateway/config"
	"temulokal-microservice/gateway/middleware"
	"temulokal-microservice/shared-service/jwt"
	"temulokal-microservice/shared-service/logger"
	"temulokal-microservice/shared-service/response"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func StartProxy() {
	cfg := config.Load()
	jwtManager := jwt.NewJWTManager(cfg.JWTSecret, time.Hour*24*5)

	app := fiber.New(fiber.Config{
		AppName: "Gateway",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			logger.Error(err.Error())
			return response.Error(c, fiber.StatusInternalServerError, "Internal server error", err.Error())
		},
	})

	app.Use(recover.New())
	allowedOrigins := cfg.FrontendOrigin
	allowCreds := true
	if allowedOrigins == "" {
		allowedOrigins = "*"
		allowCreds = false // browsers block credentials with wildcard origin
	}
	app.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With, X-User-ID",
		ExposeHeaders:    "Content-Length, X-User-ID",
		AllowCredentials: allowCreds,
		MaxAge:           300,
	}))
	app.Use(middleware.RateLimit())

	RegisterRoutes(app, cfg, jwtManager)

	logger.Success(fmt.Sprintf("Gateway proxy running on port %s\n", cfg.AppPort))
	app.Listen(":" + cfg.AppPort)
}
