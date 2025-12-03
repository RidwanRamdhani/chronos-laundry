package repositories

import (
	"github.com/RidwanRamdhani/chronos-laundry/backend/models"
	"gorm.io/gorm"
)

// ServicePriceRepository handles service price database operations
type ServicePriceRepository struct {
	db *gorm.DB
}

// NewServicePriceRepository creates a new service price repository
func NewServicePriceRepository(db *gorm.DB) *ServicePriceRepository {
	return &ServicePriceRepository{db: db}
}

// CreateServicePrice creates a new service price
func (r *ServicePriceRepository) CreateServicePrice(servicePrice *models.ServicePrice) error {
	return r.db.Create(servicePrice).Error
}

// GetServicePriceByID retrieves a service price by ID
func (r *ServicePriceRepository) GetServicePriceByID(id uint) (*models.ServicePrice, error) {
	var servicePrice models.ServicePrice
	err := r.db.Where("id = ?", id).First(&servicePrice).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &servicePrice, err
}

// GetServicePriceByTypeAndItem retrieves a service price by service type and item name
func (r *ServicePriceRepository) GetServicePriceByTypeAndItem(serviceType, itemName string) (*models.ServicePrice, error) {
	var servicePrice models.ServicePrice
	err := r.db.Where("service_type = ? AND item_name = ? AND is_active = ?", serviceType, itemName, true).
		First(&servicePrice).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &servicePrice, err
}

// GetAllServicePrices retrieves all active service prices
func (r *ServicePriceRepository) GetAllServicePrices() ([]models.ServicePrice, error) {
	var servicePrices []models.ServicePrice
	err := r.db.Where("is_active = ?", true).
		Order("service_type ASC, item_name ASC").
		Find(&servicePrices).Error
	return servicePrices, err
}

// GetServicePricesByType retrieves all service prices by service type
func (r *ServicePriceRepository) GetServicePricesByType(serviceType string) ([]models.ServicePrice, error) {
	var servicePrices []models.ServicePrice
	err := r.db.Where("service_type = ? AND is_active = ?", serviceType, true).
		Order("item_name ASC").
		Find(&servicePrices).Error
	return servicePrices, err
}

// GetServiceTypes retrieves all unique service types
func (r *ServicePriceRepository) GetServiceTypes() ([]string, error) {
	var serviceTypes []string
	err := r.db.Model(&models.ServicePrice{}).
		Where("is_active = ?", true).
		Distinct("service_type").
		Order("service_type ASC").
		Pluck("service_type", &serviceTypes).Error
	return serviceTypes, err
}

// UpdateServicePrice updates a service price
func (r *ServicePriceRepository) UpdateServicePrice(servicePrice *models.ServicePrice) error {
	return r.db.Save(servicePrice).Error
}

// DeleteServicePrice soft deletes a service price
func (r *ServicePriceRepository) DeleteServicePrice(id uint) error {
	return r.db.Delete(&models.ServicePrice{}, id).Error
}

// DeactivateServicePrice deactivates a service price
func (r *ServicePriceRepository) DeactivateServicePrice(id uint) error {
	return r.db.Model(&models.ServicePrice{}).Where("id = ?", id).Update("is_active", false).Error
}

// ActivateServicePrice activates a service price
func (r *ServicePriceRepository) ActivateServicePrice(id uint) error {
	return r.db.Model(&models.ServicePrice{}).Where("id = ?", id).Update("is_active", true).Error
}
