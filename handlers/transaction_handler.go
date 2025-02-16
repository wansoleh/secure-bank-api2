package handlers

import (
	"secure-bank-api2/database"
	"net/http"
	"github.com/labstack/echo/v4"
)

type Transaction struct {
	NoRekening string  `json:"no_rekening"`
	Nominal    float64 `json:"nominal"`
}

// Setor Tabungan (Deposit)
func Deposit(c echo.Context) error {
	var transaksi Transaction
	if err := c.Bind(&transaksi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid request payload"})
	}

	var saldo float64
	err := database.DB.QueryRow("SELECT saldo FROM nasabah WHERE no_rekening=$1", transaksi.NoRekening).Scan(&saldo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "No rekening tidak ditemukan"})
	}

	newSaldo := saldo + transaksi.Nominal
	database.DB.Exec("UPDATE nasabah SET saldo=$1 WHERE no_rekening=$2", newSaldo, transaksi.NoRekening)

	return c.JSON(http.StatusOK, map[string]interface{}{"saldo": newSaldo})
}

// Tarik Tabungan (Withdraw)
func Withdraw(c echo.Context) error {
	var transaksi Transaction
	if err := c.Bind(&transaksi); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid request payload"})
	}

	var saldo float64
	err := database.DB.QueryRow("SELECT saldo FROM nasabah WHERE no_rekening=$1", transaksi.NoRekening).Scan(&saldo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "No rekening tidak ditemukan"})
	}

	if transaksi.Nominal > saldo {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Saldo tidak cukup"})
	}

	newSaldo := saldo - transaksi.Nominal
	database.DB.Exec("UPDATE nasabah SET saldo=$1 WHERE no_rekening=$2", newSaldo, transaksi.NoRekening)

	return c.JSON(http.StatusOK, map[string]interface{}{"saldo": newSaldo})
}

// Get Saldo Nasabah
func GetSaldo(c echo.Context) error {
	noRekening := c.Param("no_rekening")

	var saldo float64
	err := database.DB.QueryRow("SELECT saldo FROM nasabah WHERE no_rekening=$1", noRekening).Scan(&saldo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "No rekening tidak ditemukan"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"saldo": saldo})
}
