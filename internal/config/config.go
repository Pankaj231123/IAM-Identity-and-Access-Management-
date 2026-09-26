package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config holds all runtime configuration loaded from environment variables.
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	RedisAddr  string
	ServerPort string
	SessionTTL time.Duration
}

// getEnv returns the value of key, or fallback if it is empty.
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// Load reads configuration from the environment (and .env if present).
func Load() (*Config, error) {
	// Error ignored on purpose: in Docker/production there is no .env file,
	// values come from the real environment.
	_ = godotenv.Load()

	cfg := &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		RedisAddr:  getEnv("REDIS_ADDR", "localhost:6379"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}

	if cfg.DBUser == "" {
		return nil, fmt.Errorf("DB_USER is required")
	}
	if cfg.DBPassword == "" {
		return nil, fmt.Errorf("DB_PASSWORD is required")
	}
	if cfg.DBName == "" {
		return nil, fmt.Errorf("DB_NAME is required")
	}

	ttlHours, err := strconv.Atoi(getEnv("SESSION_TTL_HOURS", "720"))
	if err != nil {
		return nil, fmt.Errorf("invalid SESSION_TTL_HOURS: %w", err)
	}
	if ttlHours <= 0 {
		return nil, fmt.Errorf("SESSION_TTL_HOURS must be positive, got %d", ttlHours)
	}
	cfg.SessionTTL = time.Duration(ttlHours) * time.Hour

	return cfg, nil
}

// DSN builds the Postgres connection string.
// sslmode=disable is only acceptable for local Docker; change it for real deployments.
func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}