package repositories

import (
	"github.com/RidwanRamdhani/chronos-laundry/backend/models"
	"gorm.io/gorm"
)

// AdminRepository handles admin database operations
type AdminRepository struct {
	db *gorm.DB
}

// NewAdminRepository creates a new admin repository
func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

// GetAdminByID retrieves an admin by ID
func (r *AdminRepository) GetAdminByID(id uint) (*models.Admin, error) {
	var admin models.Admin
	err := r.db.Where("id = ?", id).First(&admin).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &admin, err
}

// GetAdminByUsername retrieves an admin by username
func (r *AdminRepository) GetAdminByUsername(username string) (*models.Admin, error) {
	var admin models.Admin
	err := r.db.Where("username = ?", username).First(&admin).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &admin, err
}

// UpdateAdmin updates an admin
func (r *AdminRepository) UpdateAdmin(admin *models.Admin) error {
	return r.db.Save(admin).Error
}

// DeleteAdmin deletes an admin
func (r *AdminRepository) DeleteAdmin(id uint) error {
	return r.db.Delete(&models.Admin{}, id).Error
}

// GetAllAdmins retrieves all admins
func (r *AdminRepository) GetAllAdmins(limit, offset int) ([]models.Admin, int64, error) {
	var admins []models.Admin
	var total int64
	err := r.db.Model(&models.Admin{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = r.db.Limit(limit).Offset(offset).Find(&admins).Error
	return admins, total, err
}
