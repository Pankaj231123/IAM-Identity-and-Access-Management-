package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pankajroy/iam-service/internal/model"
	"gorm.io/gorm"
)

// UserRepository defines database operations for users.
type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindByID(id uuid.UUID) (*model.User, error)
	IncrementFailedLogins(id uuid.UUID) error
	ResetFailedLogins(id uuid.UUID) error
	LockUntil(id uuid.UUID, until time.Time) error
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a user repository backed by GORM.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		var postgresErr *pgconn.PgError
		if errors.As(err, &postgresErr) && postgresErr.Code == "23505" {
			return fmt.Errorf("create user: %w", ErrEmailTaken)
		}
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return fmt.Errorf("create user: %w", ErrEmailTaken)
		}
		return fmt.Errorf("create user: %w", err)
	}
	return nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("find user by email: %w", ErrNotFound)
		}
		return nil, fmt.Errorf("find user by email: %w", err)
	}
	return &user, nil
}

func (r *userRepository) FindByID(id uuid.UUID) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("find user by ID: %w", ErrNotFound)
		}
		return nil, fmt.Errorf("find user by ID: %w", err)
	}
	return &user, nil
}

func (r *userRepository) IncrementFailedLogins(id uuid.UUID) error {
	result := r.db.Model(&model.User{}).
		Where("id = ?", id).
		UpdateColumn("failed_login_count", gorm.Expr("failed_login_count + ?", 1))
	if result.Error != nil {
		return fmt.Errorf("increment failed logins: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("increment failed logins: %w", ErrNotFound)
	}
	return nil
}

func (r *userRepository) ResetFailedLogins(id uuid.UUID) error {
	result := r.db.Model(&model.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"failed_login_count": 0,
			"locked_until":       nil,
		})
	if result.Error != nil {
		return fmt.Errorf("reset failed logins: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("reset failed logins: %w", ErrNotFound)
	}
	return nil
}

func (r *userRepository) LockUntil(id uuid.UUID, until time.Time) error {
	result := r.db.Model(&model.User{}).
		Where("id = ?", id).
		Update("locked_until", until)
	if result.Error != nil {
		return fmt.Errorf("lock user: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("lock user: %w", ErrNotFound)
	}
	return nil
}
