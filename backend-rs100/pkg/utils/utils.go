package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
)

// HashPassword hashes a password string using SHA256 (simple dependency-free hashing).
func HashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}

// LogInfo prints informational logs.
func LogInfo(message string) {
	log.Printf("[INFO] %s", message)
}

// LogError prints error logs.
func LogError(err error) {
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
	}
}
