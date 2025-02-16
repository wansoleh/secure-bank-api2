package main

import (
	"secure-bank-api2/database"
	"secure-bank-api2/routes"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	// Initialize database
	database.InitDB()

	// Initialize Echo framework
	e := echo.New()

	// Setup Routes
	routes.SetupRoutes(e)

	// Get port from environment variables
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	fmt.Println("Server running on port", port)
	e.Logger.Fatal(e.Start(":" + port))
}
