package main

import (
	"errors"
	"log"

	"github.com/RidwanRamdhani/chronos-laundry/backend/config"
	"github.com/RidwanRamdhani/chronos-laundry/backend/models"
	"github.com/RidwanRamdhani/chronos-laundry/backend/utils"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	// Load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Initialize DB connection
	if err := config.InitDB(); err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}

	db := config.GetDB()

	// Hash password for admin
	hashedPassword, err := utils.HashPassword("admin123")
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	// Admin account data
	admin := models.Admin{
		Username: "admin",
		Password: hashedPassword,
		Email:    "admin@chronos-laundry.com",
		FullName: "System Administrator",
	}

	// Check if admin already exists
	var existing models.Admin
	result := db.Where("username = ?", admin.Username).First(&existing)

	if result.Error != nil {
		// Check if error is "record not found"
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Not found, create new admin
			if err := db.Create(&admin).Error; err != nil {
				log.Fatalf("Failed to create admin account: %v", err)
			}
			log.Printf("✓ Admin account created successfully")
			log.Printf("  Username: %s", admin.Username)
			log.Printf("  Email: %s", admin.Email)
			log.Printf("  Password: admin123 (please change after first login)")
		} else {
			// Other database error
			log.Fatalf("Database error while checking admin: %v", result.Error)
		}
	} else {
		log.Printf("Admin account already exists: %s", admin.Username)
	}

	log.Println("Admin seeding completed!")
}
