package auth

import (
	"github.com/google/uuid"
	"time"
)

type CodeResponse struct {
	Code string `json:"code" binding:"required,min=4,max=4"`
}

type TokenResponse struct {
	ID        uuid.UUID
	UserID    uint
	Revoked   bool
	ExpiresAt time.Time `gorm:"null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}
