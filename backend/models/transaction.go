package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// TransactionStatus represents the status of a laundry transaction
type TransactionStatus string

const (
	StatusQueued        TransactionStatus = "Queued"           // Queued
	StatusWashing       TransactionStatus = "Washing"          // Washing
	StatusIroning       TransactionStatus = "Ironing"          // Ironing
	StatusReadytoPickup TransactionStatus = "Ready to pick up" // Ready to pick up
	StatusCompleted     TransactionStatus = "Completed"        // Completed
)

// Transaction represents a laundry transaction
type Transaction struct {
	ID              uint                 `gorm:"primaryKey" json:"id"`
	TransactionCode string               `gorm:"type:varchar(50);uniqueIndex;not null" json:"transaction_code"` // Unique code for tracking
	CustomerName    string               `gorm:"type:varchar(255);not null" json:"customer_name"`
	CustomerPhone   string               `gorm:"type:varchar(20)" json:"customer_phone"`
	CustomerAddress string               `gorm:"type:text" json:"customer_address"`
	Notes           string               `gorm:"type:text" json:"notes"`
	Status          TransactionStatus    `gorm:"type:varchar(20);default:'antrian'" json:"status"`
	TotalPrice      float64              `json:"total_price"`
	IsPaid          bool                 `gorm:"default:false" json:"is_paid"`
	PickupDate      datatypes.Date       `json:"pickup_date"`
	CompletedAt     *time.Time           `json:"completed_at"`
	AdminID         uint                 `json:"admin_id"`
	Admin           *Admin               `gorm:"foreignKey:AdminID" json:"-"`
	Items           []TransactionItem    `gorm:"foreignKey:TransactionID" json:"items"`
	StatusHistory   []TransactionHistory `gorm:"foreignKey:TransactionID" json:"status_history"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for Transaction model
func (Transaction) TableName() string {
	return "transactions"
}
