package bootstrap

import (
	"temulokal-microservice/gateway/config"
	"temulokal-microservice/gateway/utils"
	"temulokal-microservice/shared-service/jwt"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, cfg *config.Config, jwtManager *jwt.JWTManager) error {
	app.All("/auth/*", func(c *fiber.Ctx) error {
		return utils.ProxyRequest(c, cfg.AuthServiceOrigin)
	})

	app.All("/umkm/*", func(c *fiber.Ctx) error {
		return utils.ProxyRequest(c, cfg.UMKMServiceOrigin)
	})

	return nil
}
