package repositories

import (
	"github.com/RidwanRamdhani/chronos-laundry/backend/models"
	"gorm.io/gorm"
)

// TransactionRepository handles transaction database operations
type TransactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository creates a new transaction repository
func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// CreateTransaction creates a new transaction with items
func (r *TransactionRepository) CreateTransaction(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

// GetTransactionByID retrieves a transaction by ID with preloaded relationships
func (r *TransactionRepository) GetTransactionByID(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Items").Preload("StatusHistory").Preload("Admin").
		Where("id = ?", id).First(&transaction).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &transaction, err
}

// GetTransactionByCode retrieves a transaction by transaction code
func (r *TransactionRepository) GetTransactionByCode(code string) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Items").Preload("StatusHistory").Preload("Admin").
		Where("transaction_code = ?", code).First(&transaction).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &transaction, err
}

// UpdateTransaction updates a transaction
func (r *TransactionRepository) UpdateTransaction(transaction *models.Transaction) error {
	return r.db.Save(transaction).Error
}

// DeleteTransaction soft deletes a transaction
func (r *TransactionRepository) DeleteTransaction(id uint) error {
	return r.db.Delete(&models.Transaction{}, id).Error
}

// GetAllTransactions retrieves all transactions with pagination
func (r *TransactionRepository) GetAllTransactions(limit, offset int) ([]models.Transaction, int64, error) {
	var transactions []models.Transaction
	var total int64
	err := r.db.Model(&models.Transaction{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = r.db.Preload("Items").Preload("Admin").
		Limit(limit).Offset(offset).
		Order("created_at DESC").
		Find(&transactions).Error
	return transactions, total, err
}

// GetTransactionsByStatus retrieves transactions by status with pagination
func (r *TransactionRepository) GetTransactionsByStatus(status models.TransactionStatus, limit, offset int) ([]models.Transaction, int64, error) {
	var transactions []models.Transaction
	var total int64
	err := r.db.Model(&models.Transaction{}).Where("status = ?", status).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = r.db.Preload("Items").Preload("Admin").
		Where("status = ?", status).
		Limit(limit).Offset(offset).
		Order("created_at DESC").
		Find(&transactions).Error
	return transactions, total, err
}

// GetTransactionsByCustomerPhone retrieves transactions by customer phone
func (r *TransactionRepository) GetTransactionsByCustomerPhone(phone string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Items").Preload("Admin").
		Where("customer_phone = ?", phone).
		Order("created_at DESC").
		Find(&transactions).Error
	return transactions, err
}

// UpdateTransactionStatus updates transaction status and creates history record
func (r *TransactionRepository) UpdateTransactionStatus(transactionID uint, newStatus models.TransactionStatus, changedBy, reason string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Get current transaction
		var transaction models.Transaction
		if err := tx.Where("id = ?", transactionID).First(&transaction).Error; err != nil {
			return err
		}

		previousStatus := transaction.Status

		// Update transaction status
		if err := tx.Model(&transaction).Update("status", newStatus).Error; err != nil {
			return err
		}

		// Create history record
		history := models.TransactionHistory{
			TransactionID:  transactionID,
			PreviousStatus: previousStatus,
			NewStatus:      newStatus,
			ChangedBy:      changedBy,
			Reason:         reason,
		}
		return tx.Create(&history).Error
	})
}

// UpdatePaymentStatus updates the payment status of a transaction
func (r *TransactionRepository) UpdatePaymentStatus(id uint, isPaid bool) error {
	return r.db.Model(&models.Transaction{}).Where("id = ?", id).Update("is_paid", isPaid).Error
}

// GetTransactionsByDateRange retrieves transactions within a date range
func (r *TransactionRepository) GetTransactionsByDateRange(startDate, endDate int64, limit, offset int) ([]models.Transaction, int64, error) {
	var transactions []models.Transaction
	var total int64
	query := r.db.Model(&models.Transaction{}).
		Where("created_at >= ? AND created_at <= ?", startDate, endDate)

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Preload("Items").Preload("Admin").
		Limit(limit).Offset(offset).
		Order("created_at DESC").
		Find(&transactions).Error
	return transactions, total, err
}

// GetUnpaidTransactions retrieves all unpaid transactions
func (r *TransactionRepository) GetUnpaidTransactions(limit, offset int) ([]models.Transaction, int64, error) {
	var transactions []models.Transaction
	var total int64
	err := r.db.Model(&models.Transaction{}).Where("is_paid = ?", false).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = r.db.Preload("Items").Preload("Admin").
		Where("is_paid = ?", false).
		Limit(limit).Offset(offset).
		Order("created_at DESC").
		Find(&transactions).Error
	return transactions, total, err
}

// GetTransactionsByAdminID retrieves transactions created by a specific admin
func (r *TransactionRepository) GetTransactionsByAdminID(adminID uint, limit, offset int) ([]models.Transaction, int64, error) {
	var transactions []models.Transaction
	var total int64
	err := r.db.Model(&models.Transaction{}).Where("admin_id = ?", adminID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = r.db.Preload("Items").Preload("Admin").
		Where("admin_id = ?", adminID).
		Limit(limit).Offset(offset).
		Order("created_at DESC").
		Find(&transactions).Error
	return transactions, total, err
}

// GetTransactionHistory retrieves status history for a transaction
func (r *TransactionRepository) GetTransactionHistory(transactionID uint) ([]models.TransactionHistory, error) {
	var history []models.TransactionHistory
	err := r.db.Where("transaction_id = ?", transactionID).
		Order("created_at DESC").
		Find(&history).Error
	return history, err
}

// SearchTransactions searches transactions by customer name or transaction code
func (r *TransactionRepository) SearchTransactions(keyword string, limit, offset int) ([]models.Transaction, int64, error) {
	var transactions []models.Transaction
	var total int64
	searchPattern := "%" + keyword + "%"
	query := r.db.Model(&models.Transaction{}).
		Where("customer_name LIKE ? OR transaction_code LIKE ?", searchPattern, searchPattern)

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Preload("Items").Preload("Admin").
		Limit(limit).Offset(offset).
		Order("created_at DESC").
		Find(&transactions).Error
	return transactions, total, err
}

// GetDashboardStats retrieves statistics for dashboard
func (r *TransactionRepository) GetDashboardStats() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// Total transactions
	var totalTransactions int64
	if err := r.db.Model(&models.Transaction{}).Count(&totalTransactions).Error; err != nil {
		return nil, err
	}
	stats["total_transactions"] = totalTransactions

	// Transactions by status
	var statusCounts []struct {
		Status models.TransactionStatus
		Count  int64
	}
	if err := r.db.Model(&models.Transaction{}).
		Select("status, count(*) as count").
		Group("status").
		Scan(&statusCounts).Error; err != nil {
		return nil, err
	}
	stats["status_counts"] = statusCounts

	// Total revenue
	var totalRevenue float64
	if err := r.db.Model(&models.Transaction{}).
		Where("is_paid = ?", true).
		Select("COALESCE(SUM(total_price), 0)").
		Scan(&totalRevenue).Error; err != nil {
		return nil, err
	}
	stats["total_revenue"] = totalRevenue

	// Unpaid amount
	var unpaidAmount float64
	if err := r.db.Model(&models.Transaction{}).
		Where("is_paid = ?", false).
		Select("COALESCE(SUM(total_price), 0)").
		Scan(&unpaidAmount).Error; err != nil {
		return nil, err
	}
	stats["unpaid_amount"] = unpaidAmount

	return stats, nil
}
