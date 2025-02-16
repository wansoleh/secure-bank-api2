package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct untuk menyimpan konfigurasi aplikasi
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	APIPort    string
	LogLevel   string
}

// LoadConfig membaca konfigurasi dari file `.env`
func LoadConfig() *Config {
	// Load file .env jika tersedia
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	// Membaca variabel lingkungan
	config := &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "secure_bank"),
		JWTSecret:  getEnv("JWT_SECRET", "default_jwt_secret"),
		APIPort:    getEnv("API_PORT", "8080"),
		LogLevel:   getEnv("LOG_LEVEL", "INFO"),
	}

	return config
}

// getEnv membaca variabel lingkungan dengan nilai default jika tidak ditemukan
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// GetDatabaseDSN mengembalikan DSN (Data Source Name) untuk koneksi PostgreSQL
func (c *Config) GetDatabaseDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}
