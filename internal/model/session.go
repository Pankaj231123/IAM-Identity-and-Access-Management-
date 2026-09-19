package model

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID  `json:"user_id" gorm:"type:uuid;not null"`
	TokenHash string     `json:"-" gorm:"not null"`
	IpAddress string     `json:"ip_address" gorm:"not null"`
	UserAgent string     `json:"user_agent" gorm:"not null"`
	RevodedAt *time.Time `json:"revoked_at" gorm:"default:null"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	ExpiresAt time.Time  `json:"expires_at" gorm:"not null"`
}
