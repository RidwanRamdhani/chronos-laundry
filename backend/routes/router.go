package routes

import (
	"time"

	"github.com/RidwanRamdhani/chronos-laundry/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	authController *controllers.AuthController,
	transactionController *controllers.TransactionController,
	servicePriceController *controllers.ServicePriceController,
) *gin.Engine {

	r := gin.Default()

	// Disable automatic trailing slash redirect to prevent CORS issues
	r.RedirectTrailingSlash = false

	// CORS Middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")

	// Auth
	AuthRoutes(api, authController)

	// Transactions
	TransactionRoutes(api, transactionController)

	// Service Prices
	SetupServicePriceRoutes(r, servicePriceController)

	return r
}
