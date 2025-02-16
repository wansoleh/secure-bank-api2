package handlers

import (
	"account-service/database"
	"account-service/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Get All Jurnal Umum
func GetJurnalUmum(c echo.Context) error {
	rows, err := database.DB.Query("SELECT id, tanggal, no_rekening, id_master_account, tipe, nominal, keterangan, id_karyawan, input_by FROM jurnal_umum")
	if err != nil {
		log.Println("Error fetching jurnal umum:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Error fetching data"})
	}
	defer rows.Close()

	var jurnalEntries []models.JurnalUmum
	for rows.Next() {
		var entry models.JurnalUmum
		if err := rows.Scan(&entry.ID, &entry.Tanggal, &entry.NoRekening, &entry.IDMasterAccount, &entry.Tipe, &entry.Nominal, &entry.Keterangan, &entry.IDKaryawan, &entry.InputBy); err != nil {
			log.Println("Error scanning jurnal umum:", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Error reading data"})
		}
		jurnalEntries = append(jurnalEntries, entry)
	}

	return c.JSON(http.StatusOK, jurnalEntries)
}

// Create Jurnal Umum Entry
func CreateJurnalUmum(c echo.Context) error {
	var jurnal models.JurnalUmum
	if err := c.Bind(&jurnal); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid request payload"})
	}

	_, err := database.DB.Exec("INSERT INTO jurnal_umum (no_rekening, id_master_account, tipe, nominal, keterangan, id_karyawan, input_by) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		jurnal.NoRekening, jurnal.IDMasterAccount, jurnal.Tipe, jurnal.Nominal, jurnal.Keterangan, jurnal.IDKaryawan, jurnal.InputBy)
	if err != nil {
		log.Println("Error inserting jurnal umum:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Gagal menambahkan jurnal umum"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Jurnal Umum berhasil ditambahkan"})
}
