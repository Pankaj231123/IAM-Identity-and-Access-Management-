package repository

import (
	"fmt"
	"time"

	"github.com/pankajroy/iam-service/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// OpenPostgres opens a GORM connection to PostgreSQL and verifies that the
// database is reachable before returning it.
func OpenPostgres(cfg *config.Config) (*gorm.DB, error) {
	if cfg == nil {
		return nil, fmt.Errorf("database config is required")
	}

	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open PostgreSQL connection: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get underlying SQL connection: %w", err)
	}

	// Keep a small pool suitable for a single service instance; tune these
	// values for the expected deployment size.
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("ping PostgreSQL: %w", err)
	}

	return db, nil
}
