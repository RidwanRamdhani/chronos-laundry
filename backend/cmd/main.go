package main

import (
	"log"

	"github.com/RidwanRamdhani/chronos-laundry/backend/config"
	"github.com/RidwanRamdhani/chronos-laundry/backend/controllers"
	"github.com/RidwanRamdhani/chronos-laundry/backend/repositories"
	"github.com/RidwanRamdhani/chronos-laundry/backend/routes"
	"github.com/RidwanRamdhani/chronos-laundry/backend/services"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Initialize DB connection
	if err := config.InitDB(); err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}

	db := config.GetDB()

	// Repositories
	adminRepo := repositories.NewAdminRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)
	historyRepo := repositories.NewTransactionHistoryRepository(db)
	servicePriceRepo := repositories.NewServicePriceRepository(db)

	// Services
	authService := services.NewAuthService(adminRepo)
	transactionService := services.NewTransactionService(transactionRepo, historyRepo)
	servicePriceService := services.NewServicePriceService(servicePriceRepo)

	// Controllers
	authController := controllers.NewAuthController(authService)
	transactionController := controllers.NewTransactionController(transactionService, servicePriceService)
	servicePriceController := controllers.NewServicePriceController(servicePriceService)

	// Router
	r := routes.SetupRouter(authController, transactionController, servicePriceController)
	r.Run(":8080")
}
