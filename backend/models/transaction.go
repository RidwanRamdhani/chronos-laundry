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
	TotalWeight     float64              `json:"total_weight"` // in kg
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

// TransactionItem represents an item in a transaction
type TransactionItem struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	TransactionID uint    `gorm:"not null;index" json:"transaction_id"`
	ServiceType   string  `gorm:"not null" json:"service_type"` // e.g., "cuci", "setrika", "cuci_setrika"
	ItemName      string  `gorm:"not null" json:"item_name"`    // e.g., "kemeja", "celana", "selimut"
	Quantity      int     `gorm:"default:1" json:"quantity"`
	UnitPrice     float64 `json:"unit_price"`
	Subtotal      float64 `json:"subtotal"`

	CreatedAt int64          `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64          `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for TransactionItem model
func (TransactionItem) TableName() string {
	return "transaction_items"
}

// TransactionHistory tracks status changes of a transaction
type TransactionHistory struct {
	ID             uint              `gorm:"primaryKey" json:"id"`
	TransactionID  uint              `gorm:"not null;index" json:"transaction_id"`
	PreviousStatus TransactionStatus `json:"previous_status"`
	NewStatus      TransactionStatus `gorm:"not null" json:"new_status"`
	ChangedBy      string            `json:"changed_by"` // admin username
	Reason         string            `gorm:"type:text" json:"reason"`

	CreatedAt int64 `gorm:"autoCreateTime:milli" json:"created_at"`
}

// TableName specifies the table name for TransactionHistory model
func (TransactionHistory) TableName() string {
	return "transaction_history"
}
