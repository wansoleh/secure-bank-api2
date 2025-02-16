package handlers

import (
	"secure-bank-api2/database"
	"secure-bank-api2/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Get All Produk Tabungan
func GetProdukTabungan(c echo.Context) error {
	rows, err := database.DB.Query("SELECT id, nama_produk, deskripsi, bunga, biaya_administrasi, min_saldo FROM produk_tabungan")
	if err != nil {
		log.Println("Error fetching produk tabungan:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Error fetching data"})
	}
	defer rows.Close()

	var products []models.ProdukTabungan
	for rows.Next() {
		var prod models.ProdukTabungan
		if err := rows.Scan(&prod.ID, &prod.NamaProduk, &prod.Deskripsi, &prod.Bunga, &prod.BiayaAdministrasi, &prod.MinSaldo); err != nil {
			log.Println("Error scanning produk tabungan:", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Error reading data"})
		}
		products = append(products, prod)
	}

	return c.JSON(http.StatusOK, products)
}

// Create Produk Tabungan
func CreateProdukTabungan(c echo.Context) error {
	var produk models.ProdukTabungan
	if err := c.Bind(&produk); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid request payload"})
	}

	_, err := database.DB.Exec("INSERT INTO produk_tabungan (nama_produk, deskripsi, bunga, biaya_administrasi, min_saldo) VALUES ($1, $2, $3, $4, $5)",
		produk.NamaProduk, produk.Deskripsi, produk.Bunga, produk.BiayaAdministrasi, produk.MinSaldo)
	if err != nil {
		log.Println("Error inserting produk tabungan:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Gagal menambahkan produk tabungan"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Produk Tabungan berhasil ditambahkan"})
}
