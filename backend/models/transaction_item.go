package models

import "gorm.io/gorm"

// TransactionItem represents an item in a transaction
type TransactionItem struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	TransactionID uint    `gorm:"not null;index" json:"transaction_id"`
	ServiceType   string  `gorm:"type:varchar(50);not null" json:"service_type"` // e.g., "reguler", "express"
	ItemName      string  `gorm:"type:varchar(100);not null" json:"item_name"`   // e.g., "kemeja", "celana", "selimut"
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
