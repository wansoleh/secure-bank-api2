package routes

import (
	"secure-bank-api2/handlers"
	"secure-bank-api2/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	// Users Routes
	e.GET("/users", handlers.GetUsers, middleware.JWTMiddleware())
	e.POST("/user", handlers.CreateUser, middleware.JWTMiddleware())

	// Produk Tabungan Routes
	e.GET("/produk_tabungan", handlers.GetProdukTabungan, middleware.JWTMiddleware())
	e.POST("/produk_tabungan", handlers.CreateProdukTabungan, middleware.JWTMiddleware())

	// Jurnal Umum Routes
	e.GET("/jurnal_umum", handlers.GetJurnalUmum, middleware.JWTMiddleware())
	e.POST("/jurnal_umum", handlers.CreateJurnalUmum, middleware.JWTMiddleware())

	// Nasabah Routes
	e.GET("/nasabahs", handlers.GetNasabahs, middleware.JWTMiddleware())
	e.GET("/nasabah/:id", handlers.GetNasabahByID, middleware.JWTMiddleware())
	e.PUT("/nasabah/:id", handlers.UpdateNasabah, middleware.JWTMiddleware())
	e.DELETE("/nasabah/:id", handlers.DeleteNasabah, middleware.JWTMiddleware())

	// Karyawan Routes
	e.GET("/karyawans", handlers.GetKaryawans, middleware.JWTMiddleware())
	e.POST("/karyawan", handlers.CreateKaryawan, middleware.JWTMiddleware())
	e.PUT("/karyawan/:id", handlers.UpdateKaryawan, middleware.JWTMiddleware())
	e.DELETE("/karyawan/:id", handlers.DeleteKaryawan, middleware.JWTMiddleware())

	// Transaksi Setor & Tarik
	e.POST("/tabung", handlers.Deposit, middleware.JWTMiddleware())
	e.POST("/tarik", handlers.Withdraw, middleware.JWTMiddleware())

	// Cek Saldo
	e.GET("/saldo/:no_rekening", handlers.GetSaldo, middleware.JWTMiddleware())


	// Auth Routes
	e.POST("/login", handlers.Login)

	
	// Master Account Routes
	e.GET("/master_accounts", handlers.GetMasterAccounts, middleware.JWTMiddleware())
	e.POST("/master_account", handlers.CreateMasterAccount, middleware.JWTMiddleware())

	
	// Log Aktivitas
	e.GET("/logs", handlers.GetLogs, middleware.JWTMiddleware())

}


