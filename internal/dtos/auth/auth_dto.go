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

type TokenDetail struct {
	AccessToken            string `json:"access_token"`
	RefreshToken           string `json:"refresh_token"`
	AccessTokenExpireTime  int64  `json:"access_token_expire_time"`
	RefreshTokenExpireTime int64  `json:"refresh_token_expire_time"`
}
