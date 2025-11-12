package handler

import (
	"temulokal-microservice/shared-service/response"
	"temulokal-microservice/umkm-service/usecase"

	"github.com/gofiber/fiber/v2"
)

type UMKMIndexHandler struct {
	usecase *usecase.UMKMUsecase
}

func NewUMKMIndexHandler(u *usecase.UMKMUsecase) *UMKMIndexHandler {
	return &UMKMIndexHandler{usecase: u}
}

// GET /umkm
func (h *UMKMIndexHandler) Handler(c *fiber.Ctx) error {
	data, err := h.usecase.List()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch UMKM data", nil)
	}
	return response.Success(c, fiber.StatusOK, "OK", data)
}
