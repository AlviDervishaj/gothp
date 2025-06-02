package utils

import (
	"log"
	"os"
)

// GetEnv loads an environment variable or returns a default value.
func GetEnv(key, defaultVal string) string {
	if val, found := os.LookupEnv(key); found {
		return val
	}
	return defaultVal
}

// Example usage: dbURL := utils.GetEnv("DATABASE_URL", "postgres://user:pass@localhost/db?sslmode=disable")
