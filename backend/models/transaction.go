package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// TransactionStatus represents the status of a laundry transaction
type TransactionStatus string

const (
	StatusAntrian     TransactionStatus = "antrian"      // Queued
	StatusMencuci     TransactionStatus = "mencuci"      // Washing
	StatusMenyetrika  TransactionStatus = "menyetrika"   // Ironing
	StatusSiapDiambil TransactionStatus = "siap_diambil" // Ready to pick up
	StatusSelesai     TransactionStatus = "selesai"      // Completed
)

// Transaction represents a laundry transaction
type Transaction struct {
	ID              uint                 `gorm:"primaryKey" json:"id"`
	TransactionCode string               `gorm:"uniqueIndex;not null" json:"transaction_code"` // Unique code for tracking
	CustomerName    string               `gorm:"not null" json:"customer_name"`
	CustomerPhone   string               `json:"customer_phone"`
	CustomerAddress string               `json:"customer_address"`
	Notes           string               `gorm:"type:text" json:"notes"`
	Status          TransactionStatus    `gorm:"default:'antrian'" json:"status"`
	TotalPrice      float64              `json:"total_price"`
	IsPaid          bool                 `gorm:"default:false" json:"is_paid"`
	PickupDate      datatypes.Date       `json:"pickup_date"`
	CompletedAt     *time.Time           `json:"completed_at"`
	AdminID         uint                 `json:"admin_id"`
	Admin           *Admin               `gorm:"foreignKey:AdminID" json:"-"`
	Items           []TransactionItem    `gorm:"foreignKey:TransactionID" json:"items"`
	StatusHistory   []TransactionHistory `gorm:"foreignKey:TransactionID" json:"status_history"`

	CreatedAt int64          `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64          `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for Transaction model
func (Transaction) TableName() string {
	return "transactions"
}
