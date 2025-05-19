package client

import (
	dtoAdmin "github.com/alimarzban99/go-blog-api/internal/dtos/admin"
	"github.com/alimarzban99/go-blog-api/internal/repository"
	"github.com/alimarzban99/go-blog-api/internal/resources/admin"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{repo: repository.NewUserRepository()}
}

func (s UserService) Profile(id int) (*admin.UserResource, error) {
	return s.repo.FindOne(id)
}

func (s UserService) Update(id int, dto *dtoAdmin.UpdateUserDTO) error {
	return s.repo.Update(id, dto)
}
