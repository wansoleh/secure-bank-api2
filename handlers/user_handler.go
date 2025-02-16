package handlers

import (
	"secure-bank-api2/database"
	"secure-bank-api2/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Get All Users
func GetAllUsers(c echo.Context) error {
	rows, err := database.DB.Query("SELECT id, username, role, created_at FROM users")
	if err != nil {
		log.Println("Error fetching users:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Error fetching data"})
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Role, &user.CreatedAt); err != nil {
			log.Println("Error scanning user:", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Error reading data"})
		}
		users = append(users, user)
	}

	return c.JSON(http.StatusOK, users)
}

// Create User
func RegisterUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid request payload"})
	}

	_, err := database.DB.Exec("INSERT INTO users (username, password, role) VALUES ($1, $2, $3)",
		user.Username, user.Password, user.Role)
	if err != nil {
		log.Println("Error inserting user:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Gagal menambahkan user"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User berhasil ditambahkan"})
}
