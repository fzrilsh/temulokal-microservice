package bootstrap

import (
	"temulokal-microservice/umkm-service/config"
	"temulokal-microservice/umkm-service/handler"
	"temulokal-microservice/umkm-service/repository"
	"temulokal-microservice/umkm-service/usecase"
	"temulokal-microservice/umkm-service/utils/database"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	cfg := config.Load()

	db := database.Connect(cfg)
	umkmRepo := repository.NewUMKMRepository(db)
	umkmUsecase := usecase.NewUMKMUsecase(umkmRepo)

	umkmIndexHandler := handler.NewUMKMIndexHandler(umkmUsecase)

	umkm := app.Group("/umkm")
	umkm.Get("/", umkmIndexHandler.Handler)
}
