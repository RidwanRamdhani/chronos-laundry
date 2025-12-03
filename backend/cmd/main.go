package main

import (
	"log"

	"github.com/RidwanRamdhani/chronos-laundry/backend/config"
	"github.com/RidwanRamdhani/chronos-laundry/backend/controllers"
	"github.com/RidwanRamdhani/chronos-laundry/backend/repositories"
	"github.com/RidwanRamdhani/chronos-laundry/backend/routes"
	"github.com/RidwanRamdhani/chronos-laundry/backend/services"
)

func main() {
	// Initialize DB connection
	if err := config.InitDB(); err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}

	db := config.GetDB()

	// Repositories
	adminRepo := repositories.NewAdminRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)
	historyRepo := repositories.NewTransactionHistoryRepository(db)

	// Services
	authService := services.NewAuthService(adminRepo)
	transactionService := services.NewTransactionService(transactionRepo, historyRepo)

	// Controllers
	authController := controllers.NewAuthController(authService)
	transactionController := controllers.NewTransactionController(transactionService)

	// Router
	r := routes.SetupRouter(authController, transactionController)
	r.Run(":8080")
}
