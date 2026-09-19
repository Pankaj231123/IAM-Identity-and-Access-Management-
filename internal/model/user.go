package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
    ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email string    `json:"email" gorm:"unique;not null"`
	PasswordHash string    `json:"-" gorm:"not null"`
	FailedLoginAttempts int       `json:"failed_login_attempts" gorm:"default:0"`
	LockedUntil *time.Time `json:"locked_until" gorm:"default:null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}