package auth

import "time"

type GetOTPCodeDTO struct {
	Mobile string `json:"mobile" binding:"required,min=11,max=11"`
}

type CreateOTPCodeDTO struct {
	Mobile   string    `json:"mobile" binding:"required,min=11,max=11"`
	Code     int       `json:"code" binding:"required,min=4,max=4"`
	ExpireAt time.Time `json:"expire_at" binding:"required"`
}

type VerifyOTPCodeDTO struct {
	Mobile string `json:"mobile" binding:"required,min=11,max=11"`
	Code   string `json:"code" binding:"required,min=4,max=4"`
}

type TokenCreate struct {
	UserID    uint
	ExpiresAt time.Time
	Revoked   bool
}
