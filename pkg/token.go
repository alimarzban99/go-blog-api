package pkg

import (
	"errors"
	"github.com/alimarzban99/go-blog-api/config"
	"github.com/alimarzban99/go-blog-api/internal/dtos/auth"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(userId int) (*auth.TokenDetail, error) {
	accessToken := &auth.TokenDetail{}
	accessToken.AccessTokenExpireTime = time.Now().Add(config.Config.JWTConfig.AccessTokenExpireDuration * time.Minute).Unix()
	accessToken.RefreshTokenExpireTime = time.Now().Add(config.Config.JWTConfig.RefreshTokenExpireDuration * time.Minute).Unix()

	accessTokenClaims := jwt.MapClaims{
		"user_id": userId,
		"exp":     accessToken.AccessTokenExpireTime,
	}

	jwtAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	var err error
	accessToken.AccessToken, err = jwtAccessToken.SignedString([]byte(config.Config.JWTConfig.Secret))
	if err != nil {
		return nil, err
	}

	refreshTokenClaims := jwt.MapClaims{
		"user_id": userId,
		"exp":     accessToken.RefreshTokenExpireTime,
	}
	jwtRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	accessToken.RefreshToken, err = jwtRefreshToken.SignedString([]byte(config.Config.JWTConfig.RefreshSecret))
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func VerifyToken(token string) (*jwt.Token, error) {
	at, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.Config.JWTConfig.Secret), nil
	})

	if err != nil {
		return nil, err
	}
	return at, nil
}

func GetClaims(token string) (claimMap map[string]interface{}, err error) {
	claimMap = map[string]interface{}{}
	verifyToken, err := VerifyToken(token)
	if err != nil {
		return nil, err
	}
	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid {
		for k, v := range claims {
			claimMap[k] = v
		}
		return claimMap, nil
	}
	return nil, errors.New("invalid token")
}
