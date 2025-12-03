package models

import "gorm.io/gorm"

// Admin represents an admin user in the system
type Admin struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(255);uniqueIndex;not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"password,omitempty"`
	Email    string `gorm:"type:varchar(255);uniqueIndex" json:"email"`
	FullName string `gorm:"type:varchar(255)" json:"full_name"`

	CreatedAt int64          `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64          `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for Admin model
func (Admin) TableName() string {
	return "admins"
}
