package handler

import (
	"temulokal-microservice/auth-service/usecase"
	"temulokal-microservice/shared-service/jwt"
	"temulokal-microservice/shared-service/response"
	"temulokal-microservice/shared-service/validator"

	"github.com/gofiber/fiber/v2"
)

type LoginHandler struct {
	usecase    *usecase.AuthUsecase
	jwtManager *jwt.JWTManager
}

// constructor
func NewLoginHandler(u *usecase.AuthUsecase, jwtManager *jwt.JWTManager) *LoginHandler {
	return &LoginHandler{
		usecase:    u,
		jwtManager: jwtManager,
	}
}

// POST /auth/login
func (h *LoginHandler) Handler(c *fiber.Ctx) error {
	var input usecase.LoginInput

	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	if err := validator.ValidateStruct(input); err != nil {
		return response.Error(c, fiber.StatusUnprocessableEntity, "Invalid fields", err)
	}

	user, err := h.usecase.Login(input.Email, input.Password)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "Invalid email or password", nil)
	}

	token, err := h.jwtManager.Generate(user.ID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Something went wrong, please try again", nil)
	}

	return response.Success(c, fiber.StatusOK, "Logged in successfully", fiber.Map{
		"token":     token,
		"full_name": user.FullName,
		"email":     user.Email,
	})
}
