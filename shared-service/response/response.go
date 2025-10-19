package response

import "github.com/gofiber/fiber/v2"

type JSONResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func Success(c *fiber.Ctx, statuscode int, message string, data interface{}) error {
	response := &JSONResponse{
		Status:  "success",
		Message: message,
	}

	if data != nil {
		response.Data = data
	}

	return c.Status(statuscode).JSON(response)
}

func Error(c *fiber.Ctx, statusCode int, message string, details interface{}) error {
	response := &JSONResponse{
		Status:  "error",
		Message: message,
	}

	if details != nil {
		response.Error = details
	}

	return c.Status(statusCode).JSON(response)
}
