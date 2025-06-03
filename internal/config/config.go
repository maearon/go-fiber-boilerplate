package config

import (
	"os"
)

type Config struct {
	DatabaseURL  string
	JWTSecret    string
	SMTPHost     string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
	AppURL       string
	Environment  string
}

func Load() *Config {
	return &Config{
		DatabaseURL:  getEnv("DATABASE_URL", "postgres://user:password@localhost/go_fiber_boilerplate?sslmode=disable"),
		JWTSecret:    getEnv("JWT_SECRET", "your-secret-key"),
		SMTPHost:     getEnv("SMTP_HOST", "localhost"),
		SMTPPort:     getEnv("SMTP_PORT", "587"),
		SMTPUsername: getEnv("SMTP_USERNAME", ""),
		SMTPPassword: getEnv("SMTP_PASSWORD", ""),
		AppURL:       getEnv("APP_URL", "http://localhost:3000"),
		Environment:  getEnv("ENVIRONMENT", "development"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
