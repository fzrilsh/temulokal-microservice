package bootstrap

import (
	"fmt"
	"strings"
	"temulokal-microservice/auth-service/config"
	"temulokal-microservice/shared-service/logger"
	"temulokal-microservice/shared-service/response"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func StartServer() {
	cfg := config.Load()

	app := fiber.New(fiber.Config{
		AppName:       cfg.AppName,
		Prefork:       cfg.AppEnv == "production",
		CaseSensitive: true,
		StrictRouting: true,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			logger.Error(err.Error())
			return response.Error(c, fiber.StatusInternalServerError, "Internal server error", err.Error())
		},
	})

	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.AppOrigin,
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization, Cookie",
	}))

	RegisterRoutes(app)

	env := strings.ToUpper(cfg.AppEnv)
	logger.Success(fmt.Sprintf("🚀 Auth service running in %s on port %s\n", env, cfg.AppPort))

	app.Listen(":" + cfg.AppPort)
}
