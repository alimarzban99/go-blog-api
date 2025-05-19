package auth

import (
	"crypto/rsa"
	"errors"
	"github.com/alimarzban99/go-blog-api/config"
	"github.com/alimarzban99/go-blog-api/internal/dtos/admin"
	"github.com/alimarzban99/go-blog-api/internal/dtos/auth"
	"github.com/alimarzban99/go-blog-api/internal/repository"
	adminResources "github.com/alimarzban99/go-blog-api/internal/resources/admin"
	"github.com/golang-jwt/jwt/v5"
	"math/rand"
	"os"
	"time"
)

type Service struct {
	repo      *repository.VerificationCodeRepository
	tokenRepo *repository.TokenRepository
	userRepo  *repository.UserRepository
}

var privateKey *rsa.PrivateKey

func NewAuthService() *Service {
	loadPrivateKey()
	return &Service{
		repo:      repository.NewVerificationCodeRepository(),
		userRepo:  repository.NewUserRepository(),
		tokenRepo: repository.NewTokenRepository(),
	}
}

func (s *Service) GetVerificationCode(dto *auth.GetOTPCodeDTO) {
	expireTime := config.Config.OTPCode.ExpireTime

	_, _ = s.repo.Create(&auth.CreateOTPCodeDTO{
		Code:     s.generateOTPCode(),
		Mobile:   dto.Mobile,
		ExpireAt: time.Now().Add(expireTime * time.Second),
	})

}

func (s *Service) Verify(dto *auth.VerifyOTPCodeDTO) (string, error) {

	var user *adminResources.UserResource
	var err error

	isMobileExists := s.existsByMobile(dto.Mobile)
	if !isMobileExists {
		user, err = s.userRepo.Create(&admin.StoreUserDTO{Mobile: &dto.Mobile})
	} else {
		user, err = s.userRepo.FindByMobile(dto.Mobile)
	}

	if err != nil {
		return "", err
	}
	isCodeValid := s.checkOTPCode(dto)
	if !isCodeValid {
		return "", errors.New("code invalid")
	}

	expiration := time.Now().Add(time.Hour * 24)
	tokenData, err := s.tokenRepo.Create(&auth.TokenCreate{UserID: uint(user.ID), ExpiresAt: expiration, Revoked: false})

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": expiration.Unix(),
		"jti": tokenData.ID,
	})

	tokenStr, _ := token.SignedString(privateKey)
	return tokenStr, nil
}

func (s *Service) Logout(jti string) {
	s.tokenRepo.ExpiredToken(jti)
}

func (s *Service) generateOTPCode() int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return int(uint(r.Intn(9000) + 1000))
}

func (s *Service) existsByMobile(mobile string) bool {

	exists, err := s.userRepo.CheckExistsByMobile(mobile)

	if err != nil || !exists {
		return false
	}
	return true
}

func (s *Service) checkOTPCode(dto *auth.VerifyOTPCodeDTO) bool {

	ok, err := s.repo.ValidCode(dto)
	if err != nil || !ok {
		return false
	}

	return true
}

func loadPrivateKey() {
	keyData, err := os.ReadFile("keys/private.pem")
	if err != nil {
		panic("Could not read private key: " + err.Error())
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		panic("Could not parse private key: " + err.Error())
	}

	privateKey = key
}
