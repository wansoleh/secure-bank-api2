package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	jwtMiddleware "github.com/labstack/echo-jwt/v4"
)

// Ambil JWT Secret dari Environment Variable
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Generate JWT Token
func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token berlaku 24 jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// JWT Middleware
func JWTMiddleware() echo.MiddlewareFunc {
	return jwtMiddleware.WithConfig(jwtMiddleware.Config{
		SigningKey:    jwtSecret,              // Kunci untuk verifikasi token
		SigningMethod: "HS256",                // Gunakan algoritma HS256
		TokenLookup:   "header:Authorization", // Ambil token dari header Authorization
	})
}
