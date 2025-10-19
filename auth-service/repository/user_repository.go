package repository

import (
	"encoding/json"
	"errors"
	"net/http"

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
	resp, err := r.HTTPClient.Get("http://localhost:8002/api/users/find-by-email?email="+email, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("user not found")
	}

	var user model.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
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
