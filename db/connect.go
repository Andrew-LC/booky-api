package db

import (
	"fmt"
	"log"
	"os"
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getEnv(envName string, fallback string) string {
	if value := os.Getenv(envName); value != "" {
		return value
	} 
	return fallback
}

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		host := getEnv("DB_HOST", "localhost")
		port := getEnv("DB_PORT", "5432")
		user := getEnv("DB_USER", "myuser")
		password := getEnv("DB_PASSWORD", "mypass")
		dbname := getEnv("DB_DBNAME", "mydb")
		sslmode := getEnv("DB_SSLMODE", "disable")

		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			host, port, user, password, dbname, sslmode,
		)
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
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	DB = database
	log.Println("Connected to PostgreSQL")
}

func GetDB() *gorm.DB {
	return DB
}
