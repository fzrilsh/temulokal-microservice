package repository

import (
	"temulokal-microservice/umkm-service/model"

	"gorm.io/gorm"
)

type UMKMRepository interface {
	FindAll() ([]model.UMKM, error)
}

type umkmRepository struct {
	db *gorm.DB
}

func NewUMKMRepository(db *gorm.DB) UMKMRepository {
	return &umkmRepository{db: db}
}

func (r *umkmRepository) FindAll() ([]model.UMKM, error) {
	var list []model.UMKM
	res := r.db.Preload("Owner").Preload("Gallery").Preload("Location").Preload("WorkHours").Preload("Ratings").Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}
	return list, nil
}
