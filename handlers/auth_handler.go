package handlers

import (
	"secure-bank-api2/database"
	"secure-bank-api2/middleware"
	"log"
	"net/http"
	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login Handler
func Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid request payload"})
	}

	var userID int
	err := database.DB.QueryRow("SELECT id FROM users WHERE username=$1 AND password=$2", req.Username, req.Password).Scan(&userID)
	if err != nil {
		log.Println("Login failed:", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"remark": "Invalid credentials"})
	}

	token, err := middleware.GenerateToken(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Gagal membuat token"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}
