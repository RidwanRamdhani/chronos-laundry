package routes

import (
	"github.com/RidwanRamdhani/chronos-laundry/backend/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup, authController *controllers.AuthController) {
	auth := rg.Group("/auth")

	auth.POST("/login", authController.Login)
}
