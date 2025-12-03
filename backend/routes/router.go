package routes

import (
	"github.com/RidwanRamdhani/chronos-laundry/backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	authController *controllers.AuthController,
	transactionController *controllers.TransactionController,
	servicePriceController *controllers.ServicePriceController,
) *gin.Engine {

	r := gin.Default()

	api := r.Group("/api")

	// Auth
	AuthRoutes(api, authController)

	// Transactions
	TransactionRoutes(api, transactionController)

	// Service Prices
	SetupServicePriceRoutes(r, servicePriceController)

	return r
}
