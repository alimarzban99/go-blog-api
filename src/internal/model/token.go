package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Token struct {
	ID        string `gorm:"type:char(36);primaryKey"` // تغییر به `CHAR(36)` برای MySQL
	UserID    uint
	Revoked   *bool
	ExpiresAt time.Time `gorm:"null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
	User      *User     `gorm:"foreignKey:UserID;OnDelete:CASCADE"`
}

func (token *Token) BeforeCreate(tx *gorm.DB) (err error) {
	if token.ID == "" {
		token.ID = uuid.New().String()
	}
	return
}
