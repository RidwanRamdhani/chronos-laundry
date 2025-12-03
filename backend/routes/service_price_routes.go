package routes

import (
	"github.com/RidwanRamdhani/chronos-laundry/backend/controllers"
	"github.com/RidwanRamdhani/chronos-laundry/backend/middlewares"
	"github.com/gin-gonic/gin"
)

// SetupServicePriceRoutes sets up service price routes
func SetupServicePriceRoutes(router *gin.Engine, servicePriceController *controllers.ServicePriceController) {
	// Public routes (no authentication required)
	public := router.Group("/api")
	{
		// Get all service types (for dropdown)
		public.GET("/service-types", servicePriceController.GetServiceTypes)

		// Get all service prices
		public.GET("/service-prices", servicePriceController.GetAllServicePrices)

		// Get service prices by type (for dropdown based on selected service type)
		public.GET("/service-prices/by-type", servicePriceController.GetServicePricesByType)

		// Get single service price
		public.GET("/service-prices/:id", servicePriceController.GetServicePrice)
	}

	// Protected routes (authentication required)
	protected := router.Group("/api/service-prices")
	protected.Use(middlewares.AuthMiddleware())
	{
		// Create service price
		protected.POST("", servicePriceController.CreateServicePrice)

		// Update service price
		protected.PUT("/:id", servicePriceController.UpdateServicePrice)

		// Delete service price
		protected.DELETE("/:id", servicePriceController.DeleteServicePrice)

		// Deactivate service price
		protected.PATCH("/:id/deactivate", servicePriceController.DeactivateServicePrice)

		// Activate service price
		protected.PATCH("/:id/activate", servicePriceController.ActivateServicePrice)
	}
}
