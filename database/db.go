package database

import (
    "database/sql"
    "fmt"
    "os"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
    var err error

    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        panic(fmt.Sprintf("Failed to connect to database: %v", err))
    }

    err = DB.Ping()
    if err != nil {
        panic(fmt.Sprintf("Database is not reachable: %v", err))
    }

    fmt.Println("Database connected successfully")
}
