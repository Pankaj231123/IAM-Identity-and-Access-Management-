package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pankajroy/iam-service/internal/model"
	"gorm.io/gorm"
)

// SessionRepository defines database operations for sessions.
type SessionRepository interface {
	Create(session *model.Session) error
	FindActiveByTokenHash(tokenHash string) (*model.Session, error)
	RevokeByID(id uuid.UUID) error
	RevokeAllForUser(userID uuid.UUID) error
	ListActiveForUser(userID uuid.UUID) ([]model.Session, error)
	TouchLastSeen(id uuid.UUID, lastSeenAt time.Time) error
	DeleteExpired() (int64, error)
}

type sessionRepository struct {
	db *gorm.DB
}

// NewSessionRepository creates a session repository backed by GORM.
func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepository{db: db}
}

func (r *sessionRepository) Create(session *model.Session) error {
	if err := r.db.Create(session).Error; err != nil {
		return fmt.Errorf("create session: %w", err)
	}
	return nil
}

func (r *sessionRepository) FindActiveByTokenHash(tokenHash string) (*model.Session, error) {
	var session model.Session
	err := r.db.Where("token_hash = ? AND revoked_at IS NULL AND expires_at > ?", tokenHash, time.Now()).
		First(&session).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("find active session: %w", ErrNotFound)
		}
		return nil, fmt.Errorf("find active session: %w", err)
	}
	return &session, nil
}

func (r *sessionRepository) RevokeByID(id uuid.UUID) error {
	result := r.db.Model(&model.Session{}).
		Where("id = ? AND revoked_at IS NULL", id).
		Update("revoked_at", time.Now())
	if result.Error != nil {
		return fmt.Errorf("revoke session: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("revoke session: %w", ErrNotFound)
	}
	return nil
}

func (r *sessionRepository) RevokeAllForUser(userID uuid.UUID) error {
	result := r.db.Model(&model.Session{}).
		Where("user_id = ? AND revoked_at IS NULL AND expires_at > ?", userID, time.Now()).
		Update("revoked_at", time.Now())
	if result.Error != nil {
		return fmt.Errorf("revoke all sessions for user: %w", result.Error)
	}
	return nil
}

func (r *sessionRepository) ListActiveForUser(userID uuid.UUID) ([]model.Session, error) {
	var sessions []model.Session
	err := r.db.Where("user_id = ? AND revoked_at IS NULL AND expires_at > ?", userID, time.Now()).
		Order("created_at DESC").
		Find(&sessions).Error
	if err != nil {
		return nil, fmt.Errorf("list active sessions: %w", err)
	}
	return sessions, nil
}

func (r *sessionRepository) TouchLastSeen(id uuid.UUID, lastSeenAt time.Time) error {
	result := r.db.Model(&model.Session{}).
		Where("id = ? AND revoked_at IS NULL AND expires_at > ?", id, time.Now()).
		Update("last_seen_at", lastSeenAt)
	if result.Error != nil {
		return fmt.Errorf("update session last seen time: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("update session last seen time: %w", ErrNotFound)
	}
	return nil
}

func (r *sessionRepository) DeleteExpired() (int64, error) {
	result := r.db.Where("expires_at <= ?", time.Now()).Delete(&model.Session{})
	if result.Error != nil {
		return 0, fmt.Errorf("delete expired sessions: %w", result.Error)
	}
	return result.RowsAffected, nil
}
