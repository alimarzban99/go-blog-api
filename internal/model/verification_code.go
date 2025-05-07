package model

import "time"

type VerificationCode struct {
	BaseModel
	Code     uint      `gorm:"not null" json:"code" binding:"required,numeric"`
	Mobile   string    `gorm:"type:varchar(15);not null;index" json:"mobile" binding:"required,e164"`
	ExpireAt time.Time `gorm:"not null" json:"expire_at"`
}
