package jwt

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	appEnv        string
	secretKey     string
	tokenDuration time.Duration
}

type UserClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func NewJWTManager(appEnv string, secretKey string, duration time.Duration) *JWTManager {
	return &JWTManager{
		appEnv:        appEnv,
		secretKey:     secretKey,
		tokenDuration: duration,
	}
}

// Generate token from user ID
func (j *JWTManager) Generate(userID uint) (string, error) {
	claims := UserClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

// Verify and parse token
func (j *JWTManager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

// Set token to cookie
func (j *JWTManager) SetCookie(UserID uint, c *fiber.Ctx) error {
	token, err := j.Generate(UserID)
	if err != nil {
		return err
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "freshvora_token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * 7) // 7 days

	// Security
	cookie.HTTPOnly = true
	cookie.Secure = j.appEnv == "production"
	cookie.SameSite = "None"
	cookie.Path = "/"

	c.Cookie(cookie)
	return nil
}
