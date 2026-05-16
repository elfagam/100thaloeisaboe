package config

import (
	"os"
)

// Config holds all configuration parameters.
type Config struct {
	Port      string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	JWTSecret string
}

// LoadConfig loads configuration from environmental variables with fallback values.
func LoadConfig() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "127.0.0.1"
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "root"
	}

	dbPass := os.Getenv("DB_PASSWORD")
	if dbPass == "" {
		dbPass = "secret"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "rs100_db"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "super_secret_key_for_rs100_anniversary"
	}

	return Config{
		Port:      port,
		DBHost:    dbHost,
		DBPort:    dbPort,
		DBUser:    dbUser,
		DBPass:    dbPass,
		DBName:    dbName,
		JWTSecret: jwtSecret,
	}
}
