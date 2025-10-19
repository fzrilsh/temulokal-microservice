package handler

import (
	"fmt"
	"temulokal-microservice/auth-service/config"
	"temulokal-microservice/auth-service/repository"
	"temulokal-microservice/auth-service/usecase"
	"temulokal-microservice/shared-service/jwt"
	"temulokal-microservice/shared-service/response"
	"temulokal-microservice/shared-service/validator"

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
	cfg := config.Load()
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

	url := fmt.Sprintf("%s/%s", cfg.FrontendOrigin, "login")
	h.emailRepo.SendEmail(repository.EmailData{
		To:      user.Email,
		Subject: "TemuLokal - Email Verification",
		Body: fmt.Sprintf(`<!DOCTYPE html>
			<html>
			<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Konfirmasi Akun TemuLokal</title>
			<link rel="preconnect" href="https://fonts.googleapis.com">
			<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
			<link href="https://fonts.googleapis.com/css2?family=Poppins:wght@400;600;700&display=swap" rel="stylesheet">
			<style>
				body {
				font-family: 'Poppins', 'Arial', sans-serif;
				background-color: #f8f9fa;
				margin: 0;
				padding: 20px;
				-webkit-font-smoothing: antialiased;
				-moz-osx-font-smoothing: grayscale;
				}
				.email-container {
				background-color: #ffffff;
				max-width: 500px;
				margin: 40px auto;
				border-radius: 16px;
				box-shadow: 0 8px 30px rgba(0, 0, 0, 0.07);
				padding: 40px;
				text-align: center;
				}
				.brand-logo {
				/* --- PERUBAHAN DI SINI --- */
				font-size: 30px; /* Diperbesar dari 24px */
				font-weight: 700;
				color: #343a40;
				letter-spacing: 1px; /* Sedikit penyesuaian untuk spasi */
				margin-bottom: 30px; /* Disesuaikan untuk keseimbangan */
				}
				.illustration {
				max-width: 150px;
				margin-bottom: 30px;
				}
				.content h2 {
				color: #212529;
				font-size: 26px;
				font-weight: 600;
				margin-top: 0;
				margin-bottom: 15px;
				}
				.content p {
				color: #6c757d;
				font-size: 16px;
				line-height: 1.8;
				margin-bottom: 35px;
				}
				.button {
				display: inline-block;
				background: linear-gradient(45deg, #E74C3C, #C0392B);
				color: #ffffff !important;
				text-decoration: none;
				padding: 15px 40px;
				border-radius: 50px;
				font-weight: 600;
				font-size: 16px;
				transition: all 0.3s ease;
				box-shadow: 0 5px 15px rgba(231, 76, 60, 0.3);
				}
				.button:hover {
				transform: translateY(-3px);
				box-shadow: 0 8px 20px rgba(231, 76, 60, 0.4);
				}
				.link-info {
				margin-top: 35px;
				font-size: 13px;
				color: #adb5bd;
				}
				.link-info a {
				color: #E74C3C;
				text-decoration: none;
				font-weight: 500;
				}
			</style>
			</head>
			<body>
			<div class="email-container">
				<div class="content">
				<h2>Konfirmasi Email Anda</h2>
				<p>
					Halo! Tinggal satu klik lagi untuk mengaktifkan akun TemuLokal Anda dan mulai menemukan UMKM terbaik di sekitar.
				</p>
				<a href="%s" class="button">Verifikasi Akun</a>
				</div>
				
				<div class="link-info">
				Jika tombol di atas tidak berfungsi, <a href="%s">klik di sini</a>.
				</div>
			</div>
			</body>
			</html>`, url, url),
	})

	return response.Success(c, fiber.StatusCreated, "You are registered successfully", nil)
}
