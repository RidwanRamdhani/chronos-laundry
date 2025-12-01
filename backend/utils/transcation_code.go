package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"time"
)

// GenerateTransactionCode generates a unique transaction code
// Format: CHRN-YYYYMMDD-XXXXX (where XXXXX is random alphanumeric)
func GenerateTransactionCode() string {
	timestamp := time.Now().Format("20060102")
	randomPart := generateRandomString(5)
	return fmt.Sprintf("CHRN-%s-%s", timestamp, randomPart)
}

// generateRandomString generates a random alphanumeric string of given length
func generateRandomString(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charsetLength := big.NewInt(int64(len(charset)))

	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			num = big.NewInt(int64(i % len(charset)))
		}
		result[i] = charset[num.Int64()]
	}
	return string(result)
}

// IsValidTransactionCode validates if a transaction code format is correct
func IsValidTransactionCode(code string) bool {
	parts := strings.Split(code, "-")
	return len(parts) == 3 && parts[0] == "CHRN" && len(parts[1]) == 8 && len(parts[2]) == 5
}
