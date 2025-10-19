package database

import (
	"fmt"
	"log"

	"temulokal-microservice/auth-service/config"
	"temulokal-microservice/shared-service/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get sql.DB: %v", err)
	}

	// Optional tuning
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)

	logger.Success("Connected to MySQL database")
	return db
}
