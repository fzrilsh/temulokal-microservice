package handler

import (
	"temulokal-microservice/auth-service/repository"
	"temulokal-microservice/auth-service/usecase"
	"temulokal-microservice/auth-service/utils/validator"
	"temulokal-microservice/shared-service/jwt"
	"temulokal-microservice/shared-service/response"

	"github.com/gofiber/fiber/v2"
)

type RegisterHandler struct {
	usecase    *usecase.AuthUsecase
	jwtManager *jwt.JWTManager
	emailRepo  repository.EmailRepository
}

// constructor
func NewRegisterHandler(u *usecase.AuthUsecase, jwtManager *jwt.JWTManager, emailRepo repository.EmailRepository) *RegisterHandler {
	return &RegisterHandler{
		usecase:    u,
		jwtManager: jwtManager,
		emailRepo:  emailRepo,
	}
}

// POST /auth/register
func (h *RegisterHandler) Handler(c *fiber.Ctx) error {
	var input usecase.RegisterInput

	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	if err := validator.ValidateStruct(input); err != nil {
		return response.Error(c, fiber.StatusUnprocessableEntity, "Invalid fields", err)
	}

	user, err := h.usecase.Register(&input)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	go func() {
		h.emailRepo.SendEmail(repository.EmailData{
			To:      user.Email,
			Subject: "TemuLokal - Email Verification",
			Body:    "",
		})
	}()

	return response.Success(c, fiber.StatusCreated, "You are registered successfully", nil)
}
