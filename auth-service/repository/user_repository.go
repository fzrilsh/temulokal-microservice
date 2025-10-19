package repository

import (
	"errors"

	"temulokal-microservice/auth-service/model"
	"temulokal-microservice/auth-service/utils/passwords"
	"temulokal-microservice/shared-service/httpclient"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*model.User, error)
	SaveUser(user *model.User) error
}

type userRepository struct {
	db         *gorm.DB
	HTTPClient *httpclient.Client
}

// Constructor
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db:         db,
		HTTPClient: httpclient.New(),
	}
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	result := r.db.Where("email = ?", email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (r *userRepository) SaveUser(user *model.User) error {
	hashedPassword, err := passwords.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	result := r.db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
