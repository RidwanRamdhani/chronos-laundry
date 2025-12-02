package repositories

import (
	"github.com/RidwanRamdhani/chronos-laundry/backend/models"
	"gorm.io/gorm"
)

// TransactionHistoryRepository handles transaction history database operations
type TransactionHistoryRepository struct {
	db *gorm.DB
}

// NewTransactionHistoryRepository creates a new transaction history repository
func NewTransactionHistoryRepository(db *gorm.DB) *TransactionHistoryRepository {
	return &TransactionHistoryRepository{db: db}
}

// CreateHistory creates a new transaction history record
func (r *TransactionHistoryRepository) CreateHistory(history *models.TransactionHistory) error {
	return r.db.Create(history).Error
}

// GetHistoryByTransactionID retrieves all history records for a transaction
func (r *TransactionHistoryRepository) GetHistoryByTransactionID(transactionID uint) ([]models.TransactionHistory, error) {
	var history []models.TransactionHistory
	err := r.db.Where("transaction_id = ?", transactionID).
		Order("created_at ASC").Find(&history).Error
	return history, err
}

// DeleteHistoryByTransactionID deletes all history records for a transaction
func (r *TransactionHistoryRepository) DeleteHistoryByTransactionID(transactionID uint) error {
	return r.db.Where("transaction_id = ?", transactionID).Delete(&models.TransactionHistory{}).Error
}
