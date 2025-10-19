package db

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    dsn := "host=localhost user=myuser password=mypass dbname=mydb port=5432 sslmode=disable"
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    DB = database
    log.Println("Connected to PostgreSQL")
}

func GetDB() *gorm.DB {
    return DB
}
