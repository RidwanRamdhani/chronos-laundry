package models

import (
	"time"

	"gorm.io/gorm"
)

// ServicePrice represents the pricing for laundry services
type ServicePrice struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	ServiceType string  `gorm:"type:varchar(50);not null;index:idx_service_item" json:"service_type"` // "reguler", "express"
	ItemName    string  `gorm:"type:varchar(100);not null;index:idx_service_item" json:"item_name"`   // "kemeja_cuci_setrika", "celana_cuci", etc.
	Description string  `gorm:"type:varchar(255)" json:"description"`                                 // Human-readable description
	Price       float64 `gorm:"not null" json:"price"`
	IsActive    bool    `gorm:"default:true" json:"is_active"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for ServicePrice model
func (ServicePrice) TableName() string {
	return "service_prices"
}
