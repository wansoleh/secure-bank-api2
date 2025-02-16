package handlers

import (
	"secure-bank-api2/database"
	"secure-bank-api2/models"
	"log"
	"net/http"
	"github.com/labstack/echo/v4"
)

// Get Nasabah By ID
func GetNasabahByID(c echo.Context) error {
	id := c.Param("id")

	var nasabah models.Nasabah
	err := database.DB.QueryRow("SELECT id, user_id, nama, nik, no_hp, no_rekening, saldo, created_at FROM nasabah WHERE id=$1", id).
		Scan(&nasabah.ID, &nasabah.UserID, &nasabah.Nama, &nasabah.NIK, &nasabah.NoHP, &nasabah.NoRekening, &nasabah.Saldo, &nasabah.CreatedAt)

	if err != nil {
		log.Println("Nasabah not found:", err)
		return c.JSON(http.StatusNotFound, map[string]string{"remark": "Nasabah tidak ditemukan"})
	}

	return c.JSON(http.StatusOK, nasabah)
}

// Update Nasabah
func UpdateNasabah(c echo.Context) error {
	id := c.Param("id")
	var nasabah models.Nasabah

	if err := c.Bind(&nasabah); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid request payload"})
	}

	_, err := database.DB.Exec("UPDATE nasabah SET nama=$1, nik=$2, no_hp=$3 WHERE id=$4",
		nasabah.Nama, nasabah.NIK, nasabah.NoHP, id)

	if err != nil {
		log.Println("Error updating nasabah:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Gagal memperbarui nasabah"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Nasabah berhasil diperbarui"})
}

// Delete Nasabah
func DeleteNasabah(c echo.Context) error {
	id := c.Param("id")

	_, err := database.DB.Exec("DELETE FROM nasabah WHERE id=$1", id)
	if err != nil {
		log.Println("Error deleting nasabah:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Gagal menghapus nasabah"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Nasabah berhasil dihapus"})
}

// Get All Nasabah
func GetNasabahs(c echo.Context) error {
	rows, err := database.DB.Query("SELECT id, user_id, nama, nik, no_hp, no_rekening, saldo, created_at FROM nasabah")
	if err != nil {
		log.Println("Error fetching nasabah:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Error fetching data"})
	}
	defer rows.Close()

	var nasabahs []models.Nasabah
	for rows.Next() {
		var n models.Nasabah
		if err := rows.Scan(&n.ID, &n.UserID, &n.Nama, &n.NIK, &n.NoHP, &n.NoRekening, &n.Saldo, &n.CreatedAt); err != nil {
			log.Println("Error scanning nasabah:", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Error reading data"})
		}
		nasabahs = append(nasabahs, n)
	}

	return c.JSON(http.StatusOK, nasabahs)
}

// Create Nasabah
func CreateNasabah(c echo.Context) error {
	var nasabah models.Nasabah
	if err := c.Bind(&nasabah); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid request payload"})
	}

	// Cek apakah NIK atau No HP sudah digunakan
	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM nasabah WHERE nik=$1 OR no_hp=$2)", nasabah.NIK, nasabah.NoHP).Scan(&exists)
	if err != nil {
		log.Println("Error checking NIK/NoHP:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Internal server error"})
	}
	if exists {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "NIK atau No HP sudah digunakan"})
	}

	// Insert new nasabah
	err = database.DB.QueryRow(
		"INSERT INTO nasabah (user_id, nama, nik, no_hp, no_rekening, no_pin, saldo) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		nasabah.UserID, nasabah.Nama, nasabah.NIK, nasabah.NoHP, nasabah.NoRekening, nasabah.NoPIN, nasabah.Saldo,
	).Scan(&nasabah.ID)

	if err != nil {
		log.Println("Error inserting nasabah:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Gagal membuat akun"})
	}

	log.Println("INFO: Nasabah berhasil didaftarkan dengan No Rekening:", nasabah.NoRekening)
	return c.JSON(http.StatusOK, map[string]string{"no_rekening": nasabah.NoRekening})
}