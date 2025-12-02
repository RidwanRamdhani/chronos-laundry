package controllers

import (
	"net/http"

	"github.com/RidwanRamdhani/chronos-laundry/backend/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Token    string `json:"token"`
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	admin, token, err := c.authService.Login(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	response := LoginResponse{
		ID:       admin.ID,
		Username: admin.Username,
		Email:    admin.Email,
		FullName: admin.FullName,
		Token:    token,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"data":    response,
	})
}
