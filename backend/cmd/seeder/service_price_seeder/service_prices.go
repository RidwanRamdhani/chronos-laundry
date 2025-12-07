package main

import (
	"log"

	"github.com/RidwanRamdhani/chronos-laundry/backend/config"
	"github.com/RidwanRamdhani/chronos-laundry/backend/models"
	"github.com/joho/godotenv"
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

	// Service prices data
	servicePrices := []models.ServicePrice{
		// Reguler - Cuci + Setrika
		{
			ServiceType: "reguler",
			ItemName:    "kemeja_cuci_setrika",
			Description: "Kemeja - Cuci + Setrika (Reguler)",
			Price:       5000,
			IsActive:    true,
		},
		{
			ServiceType: "reguler",
			ItemName:    "celana_cuci_setrika",
			Description: "Celana - Cuci + Setrika (Reguler)",
			Price:       4000,
			IsActive:    true,
		},
		{
			ServiceType: "reguler",
			ItemName:    "jaket_cuci_setrika",
			Description: "Jaket - Cuci + Setrika (Reguler)",
			Price:       7000,
			IsActive:    true,
		},
		{
			ServiceType: "reguler",
			ItemName:    "selimut_cuci_setrika",
			Description: "Selimut - Cuci + Setrika (Reguler)",
			Price:       10000,
			IsActive:    true,
		},
		{
			ServiceType: "reguler",
			ItemName:    "sprei_cuci_setrika",
			Description: "Sprei - Cuci + Setrika (Reguler)",
			Price:       8000,
			IsActive:    true,
		},

		// Reguler - Cuci Saja
		{
			ServiceType: "reguler",
			ItemName:    "kemeja_cuci",
			Description: "Kemeja - Cuci Saja (Reguler)",
			Price:       3000,
			IsActive:    true,
		},
		{
			ServiceType: "reguler",
			ItemName:    "celana_cuci",
			Description: "Celana - Cuci Saja (Reguler)",
			Price:       2500,
			IsActive:    true,
		},
		{
			ServiceType: "reguler",
			ItemName:    "jaket_cuci",
			Description: "Jaket - Cuci Saja (Reguler)",
			Price:       5000,
			IsActive:    true,
		},
		{
			ServiceType: "reguler",
			ItemName:    "selimut_cuci",
			Description: "Selimut - Cuci Saja (Reguler)",
			Price:       7000,
			IsActive:    true,
		},
		{
			ServiceType: "reguler",
			ItemName:    "sprei_cuci",
			Description: "Sprei - Cuci Saja (Reguler)",
			Price:       5000,
			IsActive:    true,
		},

		// Reguler - Setrika Saja
		{
			ServiceType: "reguler",
			ItemName:    "kemeja_setrika",
			Description: "Kemeja - Setrika Saja (Reguler)",
			Price:       2000,
			IsActive:    true,
		},
		{
			ServiceType: "reguler",
			ItemName:    "celana_setrika",
			Description: "Celana - Setrika Saja (Reguler)",
			Price:       1500,
			IsActive:    true,
		},
		{
			ServiceType: "reguler",
			ItemName:    "jaket_setrika",
			Description: "Jaket - Setrika Saja (Reguler)",
			Price:       3000,
			IsActive:    true,
		},
		{
			ServiceType: "reguler",
			ItemName:    "sprei_setrika",
			Description: "Sprei - Setrika Saja (Reguler)",
			Price:       3000,
			IsActive:    true,
		},

		// Express - Cuci + Setrika (Harga lebih mahal)
		{
			ServiceType: "express",
			ItemName:    "kemeja_cuci_setrika",
			Description: "Kemeja - Cuci + Setrika (Express)",
			Price:       8000,
			IsActive:    true,
		},
		{
			ServiceType: "express",
			ItemName:    "celana_cuci_setrika",
			Description: "Celana - Cuci + Setrika (Express)",
			Price:       6000,
			IsActive:    true,
		},
		{
			ServiceType: "express",
			ItemName:    "jaket_cuci_setrika",
			Description: "Jaket - Cuci + Setrika (Express)",
			Price:       10000,
			IsActive:    true,
		},
		{
			ServiceType: "express",
			ItemName:    "selimut_cuci_setrika",
			Description: "Selimut - Cuci + Setrika (Express)",
			Price:       15000,
			IsActive:    true,
		},
		{
			ServiceType: "express",
			ItemName:    "sprei_cuci_setrika",
			Description: "Sprei - Cuci + Setrika (Express)",
			Price:       12000,
			IsActive:    true,
		},

		// Express - Cuci Saja
		{
			ServiceType: "express",
			ItemName:    "kemeja_cuci",
			Description: "Kemeja - Cuci Saja (Express)",
			Price:       5000,
			IsActive:    true,
		},
		{
			ServiceType: "express",
			ItemName:    "celana_cuci",
			Description: "Celana - Cuci Saja (Express)",
			Price:       4000,
			IsActive:    true,
		},
		{
			ServiceType: "express",
			ItemName:    "jaket_cuci",
			Description: "Jaket - Cuci Saja (Express)",
			Price:       7000,
			IsActive:    true,
		},
		{
			ServiceType: "express",
			ItemName:    "selimut_cuci",
			Description: "Selimut - Cuci Saja (Express)",
			Price:       10000,
			IsActive:    true,
		},
		{
			ServiceType: "express",
			ItemName:    "sprei_cuci",
			Description: "Sprei - Cuci Saja (Express)",
			Price:       8000,
			IsActive:    true,
		},

		// Express - Setrika Saja
		{
			ServiceType: "express",
			ItemName:    "kemeja_setrika",
			Description: "Kemeja - Setrika Saja (Express)",
			Price:       3000,
			IsActive:    true,
		},
		{
			ServiceType: "express",
			ItemName:    "celana_setrika",
			Description: "Celana - Setrika Saja (Express)",
			Price:       2500,
			IsActive:    true,
		},
		{
			ServiceType: "express",
			ItemName:    "jaket_setrika",
			Description: "Jaket - Setrika Saja (Express)",
			Price:       4000,
			IsActive:    true,
		},
		{
			ServiceType: "express",
			ItemName:    "sprei_setrika",
			Description: "Sprei - Setrika Saja (Express)",
			Price:       5000,
			IsActive:    true,
		},
	}

	// Insert service prices
	for _, sp := range servicePrices {
		// Check if already exists
		var existing models.ServicePrice
		result := db.Where("service_type = ? AND item_name = ?", sp.ServiceType, sp.ItemName).First(&existing)

		if result.Error != nil {
			// Not found, create new
			if err := db.Create(&sp).Error; err != nil {
				log.Printf("Failed to create service price %s - %s: %v", sp.ServiceType, sp.ItemName, err)
			} else {
				log.Printf("Created service price: %s - %s (Rp %.0f)", sp.ServiceType, sp.ItemName, sp.Price)
			}
		} else {
			log.Printf("Service price already exists: %s - %s", sp.ServiceType, sp.ItemName)
		}
	}

	log.Println("Service prices seeding completed!")
}
