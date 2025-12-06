package routes

import (
	"github.com/RidwanRamdhani/chronos-laundry/backend/controllers"
	"github.com/RidwanRamdhani/chronos-laundry/backend/middlewares"
	"github.com/gin-gonic/gin"
)

func TransactionRoutes(rg *gin.RouterGroup, controller *controllers.TransactionController) {
	tr := rg.Group("/transactions")
	tr.Use(middlewares.AuthMiddleware())

	// CRUD + dashboard
	tr.POST("", controller.CreateTransaction)
	tr.GET("", controller.GetAllTransactions)
	tr.GET("/dashboard", controller.GetDashboard)

	tr.GET("/:id", controller.GetTransaction)
	tr.PUT("/:id", controller.UpdateTransaction)
	tr.DELETE("/:id", controller.DeleteTransaction)

	// Update status
	tr.PUT("/:id/status", controller.UpdateTransactionStatus)

	// Public tracking (tanpa auth)
	rg.GET("/track/:code", controller.TrackTransaction)
}
