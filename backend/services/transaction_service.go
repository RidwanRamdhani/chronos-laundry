package services

import (
	"fmt"

	"github.com/RidwanRamdhani/chronos-laundry/backend/models"
	"github.com/RidwanRamdhani/chronos-laundry/backend/repositories"
	"github.com/RidwanRamdhani/chronos-laundry/backend/utils"
)

// TransactionService handles transaction business logic
type TransactionService struct {
	transactionRepo *repositories.TransactionRepository
	historyRepo     *repositories.TransactionHistoryRepository
}

// NewTransactionService creates a new transaction service
func NewTransactionService(
	transactionRepo *repositories.TransactionRepository,
	historyRepo *repositories.TransactionHistoryRepository,
) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
		historyRepo:     historyRepo,
	}
}

// CreateTransaction creates a new transaction
func (s *TransactionService) CreateTransaction(transaction *models.Transaction) error {
	// Generate unique transaction code
	transaction.TransactionCode = utils.GenerateTransactionCode()

	// Set initial status to Antrian
	transaction.Status = models.StatusAntrian

	// Create transaction
	err := s.transactionRepo.CreateTransaction(transaction)
	if err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}

	// Record initial status in history
	history := &models.TransactionHistory{
		TransactionID:  transaction.ID,
		PreviousStatus: "",
		NewStatus:      models.StatusAntrian,
		ChangedBy:      "system",
		Reason:         "Transaction created",
	}
	_ = s.historyRepo.CreateHistory(history)

	return nil
}

// GetTransaction retrieves a transaction by ID
func (s *TransactionService) GetTransaction(id uint) (*models.Transaction, error) {
	transaction, err := s.transactionRepo.GetTransactionByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve transaction: %w", err)
	}
	if transaction == nil {
		return nil, fmt.Errorf("transaction not found")
	}
	return transaction, nil
}

// GetTransactionByCode retrieves a transaction by code (for tracking)
func (s *TransactionService) GetTransactionByCode(code string) (*models.Transaction, error) {
	if !utils.IsValidTransactionCode(code) {
		return nil, fmt.Errorf("invalid transaction code format")
	}

	transaction, err := s.transactionRepo.GetTransactionByCode(code)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve transaction: %w", err)
	}
	if transaction == nil {
		return nil, fmt.Errorf("transaction not found")
	}
	return transaction, nil
}

// UpdateTransaction updates a transaction
func (s *TransactionService) UpdateTransaction(transaction *models.Transaction) error {
	err := s.transactionRepo.UpdateTransaction(transaction)
	if err != nil {
		return fmt.Errorf("failed to update transaction: %w", err)
	}
	return nil
}

// UpdateTransactionStatus updates transaction status with workflow validation
func (s *TransactionService) UpdateTransactionStatus(id uint, newStatus models.TransactionStatus, adminUsername string, reason string) error {
	// Get current transaction
	transaction, err := s.GetTransaction(id)
	if err != nil {
		return err
	}

	// Validate status transition
	if !isValidStatusTransition(transaction.Status, newStatus) {
		return fmt.Errorf("invalid status transition from %s to %s", transaction.Status, newStatus)
	}

	// Update status
	err = s.transactionRepo.UpdateTransactionStatus(id, newStatus)
	if err != nil {
		return fmt.Errorf("failed to update status: %w", err)
	}

	// Record in history
	history := &models.TransactionHistory{
		TransactionID:  id,
		PreviousStatus: transaction.Status,
		NewStatus:      newStatus,
		ChangedBy:      adminUsername,
		Reason:         reason,
	}
	_ = s.historyRepo.CreateHistory(history)

	return nil
}

// DeleteTransaction deletes a transaction
func (s *TransactionService) DeleteTransaction(id uint) error {
	err := s.transactionRepo.DeleteTransaction(id)
	if err != nil {
		return fmt.Errorf("failed to delete transaction: %w", err)
	}
	return nil
}

// GetAllTransactions retrieves all transactions with pagination
func (s *TransactionService) GetAllTransactions(limit, offset int, status string) ([]models.Transaction, int64, error) {
	transactions, total, err := s.transactionRepo.GetAllTransactions(limit, offset, status)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to retrieve transactions: %w", err)
	}
	return transactions, total, nil
}

// isValidStatusTransition checks if a status transition is valid
func isValidStatusTransition(currentStatus, newStatus models.TransactionStatus) bool {
	validTransitions := map[models.TransactionStatus][]models.TransactionStatus{
		models.StatusAntrian:     {models.StatusMencuci, models.StatusSelesai},     // Antrian -> Mencuci or Selesai (cancel)
		models.StatusMencuci:     {models.StatusMenyetrika, models.StatusSelesai},  // Mencuci -> Menyetrika or Selesai (cancel)
		models.StatusMenyetrika:  {models.StatusSiapDiambil, models.StatusSelesai}, // Menyetrika -> SiapDiambil or Selesai (cancel)
		models.StatusSiapDiambil: {models.StatusSelesai},                           // SiapDiambil -> Selesai
		models.StatusSelesai:     {},                                               // Selesai is final state
	}

	transitions, exists := validTransitions[currentStatus]
	if !exists {
		return false
	}

	for _, valid := range transitions {
		if valid == newStatus {
			return true
		}
	}
	return false
}

// GetDashboardStats returns dashboard statistics
func (s *TransactionService) GetDashboardStats() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	antrian, _ := s.transactionRepo.CountTransactionsByStatus(models.StatusAntrian)
	mencuci, _ := s.transactionRepo.CountTransactionsByStatus(models.StatusMencuci)
	menyetrika, _ := s.transactionRepo.CountTransactionsByStatus(models.StatusMenyetrika)
	siapDiambil, _ := s.transactionRepo.CountTransactionsByStatus(models.StatusSiapDiambil)
	selesai, _ := s.transactionRepo.CountTransactionsByStatus(models.StatusSelesai)

	stats["antrian"] = antrian
	stats["mencuci"] = mencuci
	stats["menyetrika"] = menyetrika
	stats["siap_diambil"] = siapDiambil
	stats["selesai"] = selesai
	stats["total"] = antrian + mencuci + menyetrika + siapDiambil + selesai

	return stats, nil
}
