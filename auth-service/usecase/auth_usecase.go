package usecase

import (
	"errors"

	"temulokal-microservice/auth-service/model"
	"temulokal-microservice/auth-service/repository"
	"temulokal-microservice/auth-service/utils/passwords"
)

type AuthUsecase struct {
	userRepo repository.UserRepository
}

type UserResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type RegisterInput struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func NewAuthUsecase(userRepo repository.UserRepository) *AuthUsecase {
	return &AuthUsecase{
		userRepo: userRepo,
	}
}

func (u *AuthUsecase) Register(user *RegisterInput) (*UserResponse, error) {
	exists, _ := u.userRepo.FindByEmail(user.Email)
	if exists != nil {
		return nil, errors.New("email already registered")
	}

	data := &model.User{
		FullName: user.FullName,
		Email:    user.Email,
		Password: user.Password,
	}

	err := u.userRepo.SaveUser(data)
	if err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:       data.ID,
		FullName: data.FullName,
		Email:    user.Email,
	}, nil
}

func (u *AuthUsecase) Login(email string, password string) (*model.User, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid")
	}

	if passwords.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid")
	}

	return user, nil
}
