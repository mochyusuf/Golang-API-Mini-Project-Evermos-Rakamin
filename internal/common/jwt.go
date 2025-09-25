package common

import (
	"evermos_rakamin/internal/entity"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/net/context"
)

type JwtCustomClaims struct {
	ID       int64 `json:"id"`
	Name     string `json:"name"`
	Email	 string `json:"email"`
	IsAdmin  bool `json:"is_admin"`
	NoTelp   string `json:"no_telp"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(c context.Context, user *entity.User) (string, error) {
	expiredTime := time.Now().Local().Add(60 * time.Minute)
	claims := JwtCustomClaims{
		ID:    user.ID,
		Name:  user.Nama,
		Email: user.Email,
		IsAdmin:  user.IsAdmin,
		NoTelp: user.NoTelp,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if AppConfig == nil {
		return "", fmt.Errorf("AppConfig is not initialized")
	}

	encodedToken, err := token.SignedString([]byte(AppConfig.JWTSecretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return encodedToken, nil
}

// Helper function to get user from JWT token
func GetUserFromToken(c *fiber.Ctx) (*JwtCustomClaims, error) {
	user := c.Locals("user")
	if user == nil {
		return nil, fmt.Errorf("user not found in context")
	}

	token, ok := user.(*jwt.Token)
	if !ok {
		return nil, fmt.Errorf("invalid token format")
	}

	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims format")
	}

	return claims, nil
}