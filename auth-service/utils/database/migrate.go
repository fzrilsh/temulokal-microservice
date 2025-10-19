package database

import (
	"fmt"

	"temulokal-microservice/auth-service/config"
	"temulokal-microservice/auth-service/model"
	"temulokal-microservice/shared-service/logger"
)

func Migrate(cfg *config.Config) {
	db := Connect(cfg)

	logger.Info("Running migrations...")
	err := db.AutoMigrate(
		&model.User{},
	)

	if err != nil {
		logger.Error(fmt.Sprintf("Migration failed: %v", err))
	}

	logger.Success("Migration completed")
}
