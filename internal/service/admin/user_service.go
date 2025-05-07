package admin

import (
	dtoAdmin "github.com/alimarzban99/go-blog-api/internal/dtos/admin"
	"github.com/alimarzban99/go-blog-api/internal/model"
	"github.com/alimarzban99/go-blog-api/internal/repository"
	"github.com/alimarzban99/go-blog-api/internal/resources/admin"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{repo: repository.NewUserRepository()}
}

func (s UserService) UserList(dto *dtoAdmin.GetUserAdminListDTO) ([]model.User, int64, error) {
	return s.repo.AdminList(dto)
}

func (s UserService) Show(id int) (*admin.UserResource, error) {
	return s.repo.FindOne(id)
}

func (s UserService) Store(dto *dtoAdmin.StoreUserDTO) (*admin.UserResource, error) {
	return s.repo.Create(dto)
}

func (s UserService) Update(id int, dto *dtoAdmin.UpdateUserDTO) error {
	return s.repo.Update(id, dto)
}

func (s UserService) Destroy(id int) error {
	return s.repo.Destroy(id)
}
