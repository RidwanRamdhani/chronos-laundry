package services

import (
	"fmt"

	"github.com/RidwanRamdhani/chronos-laundry/backend/models"
	"github.com/RidwanRamdhani/chronos-laundry/backend/repositories"
)

// ServicePriceService handles service price business logic
type ServicePriceService struct {
	servicePriceRepo *repositories.ServicePriceRepository
}

// NewServicePriceService creates a new service price service
func NewServicePriceService(servicePriceRepo *repositories.ServicePriceRepository) *ServicePriceService {
	return &ServicePriceService{servicePriceRepo: servicePriceRepo}
}

// CreateServicePrice creates a new service price
func (s *ServicePriceService) CreateServicePrice(servicePrice *models.ServicePrice) error {
	// Check if service price already exists
	existing, err := s.servicePriceRepo.GetServicePriceByTypeAndItem(servicePrice.ServiceType, servicePrice.ItemName)
	if err != nil {
		return fmt.Errorf("failed to check existing service price: %w", err)
	}
	if existing != nil {
		return fmt.Errorf("service price for %s - %s already exists", servicePrice.ServiceType, servicePrice.ItemName)
	}

	err = s.servicePriceRepo.CreateServicePrice(servicePrice)
	if err != nil {
		return fmt.Errorf("failed to create service price: %w", err)
	}
	return nil
}

// GetServicePrice retrieves a service price by ID
func (s *ServicePriceService) GetServicePrice(id uint) (*models.ServicePrice, error) {
	servicePrice, err := s.servicePriceRepo.GetServicePriceByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve service price: %w", err)
	}
	if servicePrice == nil {
		return nil, fmt.Errorf("service price not found")
	}
	return servicePrice, nil
}

// GetServicePriceByTypeAndItem retrieves a service price by service type and item name
func (s *ServicePriceService) GetServicePriceByTypeAndItem(serviceType, itemName string) (*models.ServicePrice, error) {
	servicePrice, err := s.servicePriceRepo.GetServicePriceByTypeAndItem(serviceType, itemName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve service price: %w", err)
	}
	if servicePrice == nil {
		return nil, fmt.Errorf("service price not found for %s - %s", serviceType, itemName)
	}
	return servicePrice, nil
}

// GetAllServicePrices retrieves all active service prices
func (s *ServicePriceService) GetAllServicePrices() ([]models.ServicePrice, error) {
	servicePrices, err := s.servicePriceRepo.GetAllServicePrices()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve service prices: %w", err)
	}
	return servicePrices, nil
}

// GetServicePricesByType retrieves all service prices by service type
func (s *ServicePriceService) GetServicePricesByType(serviceType string) ([]models.ServicePrice, error) {
	servicePrices, err := s.servicePriceRepo.GetServicePricesByType(serviceType)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve service prices: %w", err)
	}
	return servicePrices, nil
}

// GetServiceTypes retrieves all unique service types
func (s *ServicePriceService) GetServiceTypes() ([]string, error) {
	serviceTypes, err := s.servicePriceRepo.GetServiceTypes()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve service types: %w", err)
	}
	return serviceTypes, nil
}

// UpdateServicePrice updates a service price
func (s *ServicePriceService) UpdateServicePrice(servicePrice *models.ServicePrice) error {
	// Check if service price exists
	existing, err := s.servicePriceRepo.GetServicePriceByID(servicePrice.ID)
	if err != nil {
		return fmt.Errorf("failed to check existing service price: %w", err)
	}
	if existing == nil {
		return fmt.Errorf("service price not found")
	}

	err = s.servicePriceRepo.UpdateServicePrice(servicePrice)
	if err != nil {
		return fmt.Errorf("failed to update service price: %w", err)
	}
	return nil
}

// DeleteServicePrice deletes a service price
func (s *ServicePriceService) DeleteServicePrice(id uint) error {
	err := s.servicePriceRepo.DeleteServicePrice(id)
	if err != nil {
		return fmt.Errorf("failed to delete service price: %w", err)
	}
	return nil
}

// DeactivateServicePrice deactivates a service price
func (s *ServicePriceService) DeactivateServicePrice(id uint) error {
	err := s.servicePriceRepo.DeactivateServicePrice(id)
	if err != nil {
		return fmt.Errorf("failed to deactivate service price: %w", err)
	}
	return nil
}

// ActivateServicePrice activates a service price
func (s *ServicePriceService) ActivateServicePrice(id uint) error {
	err := s.servicePriceRepo.ActivateServicePrice(id)
	if err != nil {
		return fmt.Errorf("failed to activate service price: %w", err)
	}
	return nil
}

// ValidatePrice validates if the provided price matches the database price
func (s *ServicePriceService) ValidatePrice(serviceType, itemName string, providedPrice float64) (bool, float64, error) {
	servicePrice, err := s.servicePriceRepo.GetServicePriceByTypeAndItem(serviceType, itemName)
	if err != nil {
		return false, 0, fmt.Errorf("failed to retrieve service price: %w", err)
	}
	if servicePrice == nil {
		return false, 0, fmt.Errorf("service price not found for %s - %s", serviceType, itemName)
	}

	// Check if provided price matches database price
	isValid := servicePrice.Price == providedPrice
	return isValid, servicePrice.Price, nil
}
