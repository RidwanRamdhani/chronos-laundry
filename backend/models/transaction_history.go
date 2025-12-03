package models

import "time"

// TransactionHistory tracks status changes of a transaction
type TransactionHistory struct {
	ID             uint              `gorm:"primaryKey" json:"id"`
	TransactionID  uint              `gorm:"not null;index" json:"transaction_id"`
	PreviousStatus TransactionStatus `gorm:"type:varchar(20)" json:"previous_status"`
	NewStatus      TransactionStatus `gorm:"type:varchar(20);not null" json:"new_status"`
	ChangedBy      string            `gorm:"type:varchar(255)" json:"changed_by"` // admin username
	Reason         string            `gorm:"type:text" json:"reason"`

	CreatedAt time.Time `json:"created_at"`
}

// TableName specifies the table name for TransactionHistory model
func (TransactionHistory) TableName() string {
	return "transaction_history"
}
