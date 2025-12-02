package services

import (
	"errors"

	"github.com/RidwanRamdhani/chronos-laundry/backend/models"
	"github.com/RidwanRamdhani/chronos-laundry/backend/repositories"
	"github.com/RidwanRamdhani/chronos-laundry/backend/utils"
)

type AuthService struct {
	adminRepo *repositories.AdminRepository
}

func NewAuthService(adminRepo *repositories.AdminRepository) *AuthService {
	return &AuthService{adminRepo: adminRepo}
}

func (s *AuthService) Login(username, password string) (*models.Admin, string, error) {
	// Retrieve admin from repository
	admin, err := s.adminRepo.GetAdminByUsername(username)
	if err != nil {
		return nil, "", errors.New("failed to retrieve admin data")
	}

	if admin == nil {
		return nil, "", errors.New("username not found")
	}

	// Verify password
	if !utils.VerifyPassword(admin.Password, password) {
		return nil, "", errors.New("incorrect password")
	}

	// Generate token (your utils/jwt.go)
	token, err := utils.GenerateToken(
		admin.ID,
		admin.Username,
		admin.Email,
		admin.FullName,
		6, // token valid for 6 hours
	)
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	return admin, token, nil
}
