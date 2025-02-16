package handlers

import (
	"account-service/database"
	"account-service/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Get All Master Accounts
func GetMasterAccounts(c echo.Context) error {
	rows, err := database.DB.Query("SELECT id, kode_akun, nama_akun, tipe FROM master_account")
	if err != nil {
		log.Println("Error fetching master accounts:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Error fetching data"})
	}
	defer rows.Close()

	var accounts []models.MasterAccount
	for rows.Next() {
		var acc models.MasterAccount
		if err := rows.Scan(&acc.ID, &acc.KodeAkun, &acc.NamaAkun, &acc.Tipe); err != nil {
			log.Println("Error scanning master account:", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Error reading data"})
		}
		accounts = append(accounts, acc)
	}

	return c.JSON(http.StatusOK, accounts)
}

// Create Master Account
func CreateMasterAccount(c echo.Context) error {
	var account models.MasterAccount
	if err := c.Bind(&account); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid request payload"})
	}

	_, err := database.DB.Exec("INSERT INTO master_account (kode_akun, nama_akun, tipe) VALUES ($1, $2, $3)",
		account.KodeAkun, account.NamaAkun, account.Tipe)
	if err != nil {
		log.Println("Error inserting master account:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Gagal menambahkan akun"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Master Account berhasil ditambahkan"})
}
