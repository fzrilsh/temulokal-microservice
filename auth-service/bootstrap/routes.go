package bootstrap

import (
	"time"

	"temulokal-microservice/auth-service/config"
	"temulokal-microservice/auth-service/handler"
	"temulokal-microservice/auth-service/repository"
	"temulokal-microservice/auth-service/usecase"
	"temulokal-microservice/auth-service/utils/database"
	"temulokal-microservice/shared-service/jwt"

	"github.com/gofiber/fiber/v2"
	gocache "github.com/patrickmn/go-cache"
)

func RegisterRoutes(app *fiber.App) {
	cfg := config.Load()
	cache := gocache.New(5*time.Minute, 10*time.Minute)

	db := database.Connect(cfg)
	jwtManager := jwt.NewJWTManager(cfg.JWTSecret, time.Hour*24*5)

	emailRepo := repository.NewEmailRepository(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPSender, cfg.SMTPPassword)
	userRepo := repository.NewUserRepository(db, cache)
	authUsecase := usecase.NewAuthUsecase(userRepo)

	loginHandler := handler.NewLoginHandler(authUsecase, jwtManager)
	registerHandler := handler.NewRegisterHandler(authUsecase, jwtManager, emailRepo)

	auth := app.Group("/auth")
	auth.Post("/login", loginHandler.Handler)
	auth.Post("/register", registerHandler.Handler)

	// TODO: umkm index
}
