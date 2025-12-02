package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/RidwanRamdhani/chronos-laundry/backend/models"
	"github.com/RidwanRamdhani/chronos-laundry/backend/services"
	"github.com/RidwanRamdhani/chronos-laundry/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

// TransactionController handles transaction endpoints
type TransactionController struct {
	transactionService *services.TransactionService
}

// NewTransactionController creates a new transaction controller
func NewTransactionController(transactionService *services.TransactionService) *TransactionController {
	return &TransactionController{transactionService: transactionService}
}

// CreateTransactionRequest represents a create transaction request
type CreateTransactionRequest struct {
	CustomerName    string                         `json:"customer_name" binding:"required"`
	CustomerPhone   string                         `json:"customer_phone" binding:"required"`
	CustomerAddress string                         `json:"customer_address"`
	Notes           string                         `json:"notes"`
	TotalPrice      float64                        `json:"total_price" binding:"required"`
	PickupDate      string                         `json:"pickup_date"`
	Items           []CreateTransactionItemRequest `json:"items" binding:"required,min=1"`
}

// CreateTransactionItemRequest represents a transaction item
type CreateTransactionItemRequest struct {
	ServiceType string  `json:"service_type" binding:"required"` // cuci, setrika, cuci_setrika
	ItemName    string  `json:"item_name" binding:"required"`    // kemeja, celana, selimut
	Quantity    int     `json:"quantity" binding:"required,gt=0"`
	UnitPrice   float64 `json:"unit_price" binding:"required"`
}

// CreateTransaction creates a new transaction
// @Summary Create Transaction
// @Description Create a new laundry transaction (admin only)
// @Tags Transactions
// @Security Bearer
// @Accept json
// @Produce json
// @Param body body CreateTransactionRequest true "Transaction data"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /api/v1/transactions [post]
func (c *TransactionController) CreateTransaction(ctx *gin.Context) {
	var req CreateTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(ctx, "Invalid request body: "+err.Error())
		return
	}

	// Parse pickup date if provided
	var pickupDate datatypes.Date
	if req.PickupDate != "" {
		parsedDate, err := time.Parse("2006-01-02", req.PickupDate)
		if err != nil {
			utils.BadRequest(ctx, "Invalid pickup date format, use YYYY-MM-DD")
			return
		}
		pickupDate = datatypes.Date(parsedDate)
	}

	// Create transaction items
	items := make([]models.TransactionItem, len(req.Items))
	for i, item := range req.Items {
		items[i] = models.TransactionItem{
			ServiceType: item.ServiceType,
			ItemName:    item.ItemName,
			Quantity:    item.Quantity,
			UnitPrice:   item.UnitPrice,
			Subtotal:    float64(item.Quantity) * item.UnitPrice,
		}
	}

	// Get admin ID from context (set by auth middleware)
	adminID := uint(0)
	if id, exists := ctx.Get("admin_id"); exists {
		adminID = id.(uint)
	}

	// Create transaction
	transaction := &models.Transaction{
		CustomerName:    req.CustomerName,
		CustomerPhone:   req.CustomerPhone,
		CustomerAddress: req.CustomerAddress,
		Notes:           req.Notes,
		TotalPrice:      req.TotalPrice,
		PickupDate:      pickupDate,
		Items:           items,
		AdminID:         adminID,
	}

	err := c.transactionService.CreateTransaction(transaction)
	if err != nil {
		utils.InternalServerError(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Transaction created successfully", map[string]interface{}{
		"id":               transaction.ID,
		"transaction_code": transaction.TransactionCode,
		"customer_name":    transaction.CustomerName,
		"customer_phone":   transaction.CustomerPhone,
		"status":           transaction.Status,
		"total_price":      transaction.TotalPrice,
		"is_paid":          transaction.IsPaid,
	})
}

// GetTransaction retrieves a transaction by ID (admin only)
// @Summary Get Transaction
// @Description Get transaction details by ID
// @Tags Transactions
// @Security Bearer
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /api/v1/transactions/{id} [get]
func (c *TransactionController) GetTransaction(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequest(ctx, "Invalid transaction ID")
		return
	}

	transaction, err := c.transactionService.GetTransaction(uint(id))
	if err != nil {
		utils.NotFound(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Transaction retrieved successfully", transaction)
}

// TrackTransaction retrieves transaction status by code (public, no auth required)
// @Summary Track Transaction
// @Description Get transaction status by entering transaction code
// @Tags Transactions
// @Accept json
// @Produce json
// @Param code path string true "Transaction Code (e.g., CHRN-20251201-ABC12)"
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /api/v1/track/{code} [get]
func (c *TransactionController) TrackTransaction(ctx *gin.Context) {
	code := ctx.Param("code")
	if code == "" {
		utils.BadRequest(ctx, "Transaction code is required")
		return
	}

	transaction, err := c.transactionService.GetTransactionByCode(code)
	if err != nil {
		utils.NotFound(ctx, err.Error())
		return
	}

	// Return simplified tracking info (no sensitive data)
	trackingInfo := map[string]interface{}{
		"transaction_code": transaction.TransactionCode,
		"customer_name":    transaction.CustomerName,
		"status":           transaction.Status,
		"total_price":      transaction.TotalPrice,
		"is_paid":          transaction.IsPaid,
		"pickup_date":      transaction.PickupDate,
		"items_count":      len(transaction.Items),
		"status_history":   transaction.StatusHistory,
		"created_at":       transaction.CreatedAt,
		"updated_at":       transaction.UpdatedAt,
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Transaction tracking retrieved successfully", trackingInfo)
}

// UpdateTransactionRequest represents an update transaction request
type UpdateTransactionRequest struct {
	CustomerName    string  `json:"customer_name"`
	CustomerPhone   string  `json:"customer_phone"`
	CustomerAddress string  `json:"customer_address"`
	Notes           string  `json:"notes"`
	TotalPrice      float64 `json:"total_price"`
	IsPaid          *bool   `json:"is_paid"` // pointer to distinguish between false and not provided
}

// UpdateTransaction updates a transaction (admin only)
// @Summary Update Transaction
// @Description Update transaction details
// @Tags Transactions
// @Security Bearer
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Param body body UpdateTransactionRequest true "Update data"
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /api/v1/transactions/{id} [put]
func (c *TransactionController) UpdateTransaction(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequest(ctx, "Invalid transaction ID")
		return
	}

	var req UpdateTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(ctx, "Invalid request body: "+err.Error())
		return
	}

	// Get existing transaction
	transaction, err := c.transactionService.GetTransaction(uint(id))
	if err != nil {
		utils.NotFound(ctx, err.Error())
		return
	}

	// Update fields
	if req.CustomerName != "" {
		transaction.CustomerName = req.CustomerName
	}
	if req.CustomerPhone != "" {
		transaction.CustomerPhone = req.CustomerPhone
	}
	if req.CustomerAddress != "" {
		transaction.CustomerAddress = req.CustomerAddress
	}
	if req.Notes != "" {
		transaction.Notes = req.Notes
	}
	if req.TotalPrice > 0 {
		transaction.TotalPrice = req.TotalPrice
	}
	if req.IsPaid != nil {
		transaction.IsPaid = *req.IsPaid
	}

	err = c.transactionService.UpdateTransaction(transaction)
	if err != nil {
		utils.InternalServerError(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Transaction updated successfully", transaction)
}

// UpdateStatusRequest represents a status update request
type UpdateStatusRequest struct {
	NewStatus string `json:"new_status" binding:"required"`
	Reason    string `json:"reason"`
}

// UpdateTransactionStatus updates transaction status (admin only)
// @Summary Update Transaction Status
// @Description Update transaction status with workflow validation
// @Tags Transactions
// @Security Bearer
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Param body body UpdateStatusRequest true "Status update data"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /api/v1/transactions/{id}/status [patch]
func (c *TransactionController) UpdateTransactionStatus(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequest(ctx, "Invalid transaction ID")
		return
	}

	var req UpdateStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(ctx, "Invalid request body: "+err.Error())
		return
	}

	// Get admin username from context
	adminUsername := "unknown"
	if username, exists := ctx.Get("admin_username"); exists {
		adminUsername = username.(string)
	}

	// Update status
	newStatus := models.TransactionStatus(req.NewStatus)
	err = c.transactionService.UpdateTransactionStatus(uint(id), newStatus, adminUsername, req.Reason)
	if err != nil {
		if err.Error() == "transaction not found" {
			utils.NotFound(ctx, err.Error())
			return
		}
		utils.BadRequest(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Transaction status updated successfully", map[string]interface{}{
		"id":     id,
		"status": req.NewStatus,
	})
}

// DeleteTransaction deletes a transaction (admin only)
// @Summary Delete Transaction
// @Description Delete a transaction
// @Tags Transactions
// @Security Bearer
// @Param id path int true "Transaction ID"
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /api/v1/transactions/{id} [delete]
func (c *TransactionController) DeleteTransaction(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequest(ctx, "Invalid transaction ID")
		return
	}

	err = c.transactionService.DeleteTransaction(uint(id))
	if err != nil {
		utils.InternalServerError(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Transaction deleted successfully", nil)
}

// GetAllTransactions retrieves all transactions with pagination and filtering (admin only)
// @Summary Get All Transactions
// @Description Get all transactions with pagination and optional status filter
// @Tags Transactions
// @Security Bearer
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param limit query int false "Items per page (default: 10)"
// @Param status query string false "Filter by status"
// @Success 200 {object} utils.Response
// @Router /api/v1/transactions [get]
func (c *TransactionController) GetAllTransactions(ctx *gin.Context) {
	page := 1
	limit := 10
	status := ""

	if p := ctx.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}

	if l := ctx.Query("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 && parsed <= 100 {
			limit = parsed
		}
	}

	if s := ctx.Query("status"); s != "" {
		status = s
	}

	offset := (page - 1) * limit
	transactions, total, err := c.transactionService.GetAllTransactions(limit, offset, status)
	if err != nil {
		utils.InternalServerError(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Transactions retrieved successfully", map[string]interface{}{
		"data":        transactions,
		"total":       total,
		"page":        page,
		"limit":       limit,
		"total_pages": (total + int64(limit) - 1) / int64(limit),
	})
}

// GetDashboard returns dashboard statistics (admin only)
// @Summary Get Dashboard Stats
// @Description Get transaction statistics for dashboard
// @Tags Dashboard
// @Security Bearer
// @Produce json
// @Success 200 {object} utils.Response
// @Router /api/v1/dashboard/stats [get]
func (c *TransactionController) GetDashboard(ctx *gin.Context) {
	stats, err := c.transactionService.GetDashboardStats()
	if err != nil {
		utils.InternalServerError(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Dashboard statistics retrieved successfully", stats)
}
