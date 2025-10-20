package db

import (
    "log"
    "os"
    "time"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        dsn = "host=localhost user=myuser password=mypass dbname=mydb port=5432 sslmode=disable"
    }
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Configure connection pool
    sqlDB, err := database.DB()
    if err != nil {
        log.Fatalf("Failed to get sql DB: %v", err)
    }
    sqlDB.SetMaxOpenConns(20)
    sqlDB.SetMaxIdleConns(5)
    sqlDB.SetConnMaxLifetime(30 * time.Minute)

    DB = database
    log.Println("Connected to PostgreSQL")
}

func GetDB() *gorm.DB {
    return DB
}
