package handlers

import (
	"secure-bank-api2/database"
	"secure-bank-api2/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Get All Karyawan
func GetKaryawans(c echo.Context) error {
	rows, err := database.DB.Query("SELECT id, nama, nip, posisi, created_at FROM karyawan")
	if err != nil {
		log.Println("Error fetching karyawan:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Error fetching data"})
	}
	defer rows.Close()

	var karyawans []models.Karyawan
	for rows.Next() {
		var k models.Karyawan
		if err := rows.Scan(&k.ID, &k.Nama, &k.NIP, &k.Posisi, &k.CreatedAt); err != nil {
			log.Println("Error scanning karyawan:", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Error reading data"})
		}
		karyawans = append(karyawans, k)
	}

	return c.JSON(http.StatusOK, karyawans)
}

// Create Karyawan
func CreateKaryawan(c echo.Context) error {
	var karyawan models.Karyawan
	if err := c.Bind(&karyawan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid request payload"})
	}

	_, err := database.DB.Exec("INSERT INTO karyawan (nama, nip, posisi) VALUES ($1, $2, $3)",
		karyawan.Nama, karyawan.NIP, karyawan.Posisi)
	if err != nil {
		log.Println("Error inserting karyawan:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Gagal menambahkan karyawan"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Karyawan berhasil ditambahkan"})
}

// Update Karyawan
func UpdateKaryawan(c echo.Context) error {
	id := c.Param("id")
	var karyawan models.Karyawan
	if err := c.Bind(&karyawan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid request payload"})
	}

	_, err := database.DB.Exec("UPDATE karyawan SET nama=$1, nip=$2, posisi=$3 WHERE id=$4",
		karyawan.Nama, karyawan.NIP, karyawan.Posisi, id)
	if err != nil {
		log.Println("Error updating karyawan:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Gagal memperbarui karyawan"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Karyawan berhasil diperbarui"})
}

// Delete Karyawan
func DeleteKaryawan(c echo.Context) error {
	id := c.Param("id")

	_, err := database.DB.Exec("DELETE FROM karyawan WHERE id=$1", id)
	if err != nil {
		log.Println("Error deleting karyawan:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Gagal menghapus karyawan"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Karyawan berhasil dihapus"})
}
