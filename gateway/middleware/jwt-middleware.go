package middleware

import (
	"fmt"
	"strings"
	"temulokal-microservice/shared-service/jwt"
	"temulokal-microservice/shared-service/response"

	"github.com/gofiber/fiber/v2"
)

func JWTAuth(jwtManager *jwt.JWTManager) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.Error(c, fiber.StatusUnauthorized, "Missing authorization header", nil)
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			return response.Error(c, fiber.StatusUnauthorized, "Invalid authorization format", nil)
		}

		token := parts[1]
		claims, err := jwtManager.Verify(token)

		if err != nil {
			return response.Error(c, fiber.StatusUnauthorized, "Invalid or expired token", nil)
		}

		c.Locals("user_id", claims.UserID)
		c.Request().Header.Set("X-User-ID", fmt.Sprintf("%d", claims.UserID))

		return c.Next()
	}
}
