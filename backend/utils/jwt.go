package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TokenClaims represents JWT token claims
type TokenClaims struct {
	AdminID   uint   `json:"admin_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	ExpiresAt int64  `json:"exp"`
	jwt.RegisteredClaims
}

// GenerateToken generates a JWT token for an admin
func GenerateToken(adminID uint, username, email, fullName string, expiryHours int) (string, error) {
	jwtSecret, err := getSecretKey()
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(time.Duration(expiryHours) * time.Hour)
	claims := &TokenClaims{
		AdminID:   adminID,
		Username:  username,
		Email:     email,
		FullName:  fullName,
		ExpiresAt: expirationTime.Unix(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// VerifyToken verifies a JWT token and returns the claims
func VerifyToken(tokenString string) (*TokenClaims, error) {
	jwtSecret, err := getSecretKey()
	if err != nil {
		return nil, err
	}

	claims := &TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// getSecretKey retrieves JWT secret from environment or uses default
func getSecretKey() ([]byte, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, fmt.Errorf("JWT_SECRET environment variable is not set")
	}
	return []byte(secret), nil
}
