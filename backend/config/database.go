package config

import (
	"fmt"
	"os"

	"github.com/RidwanRamdhani/chronos-laundry/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() error {
	dbUser := getEnv("DB_USER")
	dbPassword := getEnv("DB_PASSWORD")
	dbHost := getEnv("DB_HOST")
	dbPort := getEnv("DB_PORT")
	dbName := getEnv("DB_NAME")

	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		return fmt.Errorf("database environment variables not fully set")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	DB = db

	// Auto migrate models
	return AutoMigrate()
}

// AutoMigrate runs all database migrations
func AutoMigrate() error {
	return DB.AutoMigrate(
		&models.Admin{},
		&models.Transaction{},
		&models.TransactionItem{},
		&models.TransactionHistory{},
	)
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}

// getEnv gets an environment variable
func getEnv(key string) string {
	return os.Getenv(key)
}
