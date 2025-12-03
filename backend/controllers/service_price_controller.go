package controllers

import (
	"net/http"
	"strconv"

	"github.com/RidwanRamdhani/chronos-laundry/backend/models"
	"github.com/RidwanRamdhani/chronos-laundry/backend/services"
	"github.com/RidwanRamdhani/chronos-laundry/backend/utils"
	"github.com/gin-gonic/gin"
)

// ServicePriceController handles service price endpoints
type ServicePriceController struct {
	servicePriceService *services.ServicePriceService
}

// NewServicePriceController creates a new service price controller
func NewServicePriceController(servicePriceService *services.ServicePriceService) *ServicePriceController {
	return &ServicePriceController{servicePriceService: servicePriceService}
}

// CreateServicePriceRequest represents a create service price request
type CreateServicePriceRequest struct {
	ServiceType string  `json:"service_type" binding:"required"`
	ItemName    string  `json:"item_name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
}

// CreateServicePrice creates a new service price
func (c *ServicePriceController) CreateServicePrice(ctx *gin.Context) {
	var req CreateServicePriceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(ctx, "Invalid request body: "+err.Error())
		return
	}

	servicePrice := &models.ServicePrice{
		ServiceType: req.ServiceType,
		ItemName:    req.ItemName,
		Description: req.Description,
		Price:       req.Price,
		IsActive:    true,
	}

	err := c.servicePriceService.CreateServicePrice(servicePrice)
	if err != nil {
		utils.BadRequest(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Service price created successfully", servicePrice)
}

// GetServicePrice retrieves a service price by ID
func (c *ServicePriceController) GetServicePrice(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequest(ctx, "Invalid service price ID")
		return
	}

	servicePrice, err := c.servicePriceService.GetServicePrice(uint(id))
	if err != nil {
		utils.NotFound(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Service price retrieved successfully", servicePrice)
}

// GetAllServicePrices retrieves all active service prices
func (c *ServicePriceController) GetAllServicePrices(ctx *gin.Context) {
	servicePrices, err := c.servicePriceService.GetAllServicePrices()
	if err != nil {
		utils.InternalServerError(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Service prices retrieved successfully", servicePrices)
}

// GetServicePricesByType retrieves service prices by service type
func (c *ServicePriceController) GetServicePricesByType(ctx *gin.Context) {
	serviceType := ctx.Query("service_type")
	if serviceType == "" {
		utils.BadRequest(ctx, "service_type query parameter is required")
		return
	}

	servicePrices, err := c.servicePriceService.GetServicePricesByType(serviceType)
	if err != nil {
		utils.InternalServerError(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Service prices retrieved successfully", servicePrices)
}

// GetServiceTypes retrieves all unique service types
func (c *ServicePriceController) GetServiceTypes(ctx *gin.Context) {
	serviceTypes, err := c.servicePriceService.GetServiceTypes()
	if err != nil {
		utils.InternalServerError(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Service types retrieved successfully", serviceTypes)
}

// UpdateServicePriceRequest represents an update service price request
type UpdateServicePriceRequest struct {
	ServiceType string  `json:"service_type"`
	ItemName    string  `json:"item_name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	IsActive    *bool   `json:"is_active"`
}

// UpdateServicePrice updates a service price
func (c *ServicePriceController) UpdateServicePrice(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequest(ctx, "Invalid service price ID")
		return
	}

	var req UpdateServicePriceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(ctx, "Invalid request body: "+err.Error())
		return
	}

	// Get existing service price
	servicePrice, err := c.servicePriceService.GetServicePrice(uint(id))
	if err != nil {
		utils.NotFound(ctx, err.Error())
		return
	}

	// Update fields
	if req.ServiceType != "" {
		servicePrice.ServiceType = req.ServiceType
	}
	if req.ItemName != "" {
		servicePrice.ItemName = req.ItemName
	}
	if req.Description != "" {
		servicePrice.Description = req.Description
	}
	if req.Price > 0 {
		servicePrice.Price = req.Price
	}
	if req.IsActive != nil {
		servicePrice.IsActive = *req.IsActive
	}

	err = c.servicePriceService.UpdateServicePrice(servicePrice)
	if err != nil {
		utils.InternalServerError(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Service price updated successfully", servicePrice)
}

// DeleteServicePrice deletes a service price
func (c *ServicePriceController) DeleteServicePrice(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequest(ctx, "Invalid service price ID")
		return
	}

	err = c.servicePriceService.DeleteServicePrice(uint(id))
	if err != nil {
		utils.InternalServerError(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Service price deleted successfully", nil)
}

// DeactivateServicePrice deactivates a service price
func (c *ServicePriceController) DeactivateServicePrice(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequest(ctx, "Invalid service price ID")
		return
	}

	err = c.servicePriceService.DeactivateServicePrice(uint(id))
	if err != nil {
		utils.InternalServerError(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Service price deactivated successfully", nil)
}

// ActivateServicePrice activates a service price
func (c *ServicePriceController) ActivateServicePrice(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequest(ctx, "Invalid service price ID")
		return
	}

	err = c.servicePriceService.ActivateServicePrice(uint(id))
	if err != nil {
		utils.InternalServerError(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Service price activated successfully", nil)
}
