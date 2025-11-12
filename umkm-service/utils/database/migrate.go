package database

import (
	"fmt"
	"temulokal-microservice/shared-service/logger"
	"temulokal-microservice/umkm-service/config"
	"temulokal-microservice/umkm-service/model"
)

// Migrate connects to DB using cfg and runs automigrations, similar to auth-service style
func Migrate(cfg *config.Config) {
	db := Connect(cfg)

	logger.Info("Running UMKM migrations...")
	err := db.AutoMigrate(
		&model.UMKM{},
		&model.UMKMOwner{},
		&model.UMKMGallery{},
		&model.UMKMLocation{},
		&model.UMKMWorkHour{},
		&model.UMKMRating{},
	)
	if err != nil {
		logger.Error(fmt.Sprintf("Migration failed: %v", err))
		return
	}
	logger.Success("UMKM migration completed")
}
