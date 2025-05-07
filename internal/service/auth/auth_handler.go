package auth

import (
	"errors"
	"github.com/alimarzban99/go-blog-api/config"
	"github.com/alimarzban99/go-blog-api/internal/dtos/admin"
	"github.com/alimarzban99/go-blog-api/internal/dtos/auth"
	"github.com/alimarzban99/go-blog-api/internal/repository"
	adminResources "github.com/alimarzban99/go-blog-api/internal/resources/admin"
	"github.com/alimarzban99/go-blog-api/pkg"
	"math/rand"
	"time"
)

type Service struct {
	repo     *repository.VerificationCodeRepository
	userRepo *repository.UserRepository
}

func NewAuthService() *Service {
	return &Service{
		repo:     repository.NewVerificationCodeRepository(),
		userRepo: repository.NewUserRepository(),
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

func (s *Service) Verify(dto *auth.VerifyOTPCodeDTO) (*auth.TokenDetail, error) {

	var user *adminResources.UserResource
	var err error

	isMobileExists := s.existsByMobile(dto.Mobile)
	if !isMobileExists {
		user, err = s.userRepo.Create(&admin.StoreUserDTO{Mobile: &dto.Mobile})
	} else {
		user, err = s.userRepo.FindByMobile(dto.Mobile)
	}

	if err != nil {
		return nil, err
	}
	isCodeValid := s.checkOTPCode(dto)
	if !isCodeValid {
		return nil, errors.New("code invalid")
	}

	return pkg.GenerateToken(user.ID)
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
