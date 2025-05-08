package admin

import (
	dtoAdmin "github.com/alimarzban99/go-blog-api/internal/dtos/admin"
	"github.com/alimarzban99/go-blog-api/internal/model"
	"github.com/alimarzban99/go-blog-api/internal/repository"
	"github.com/alimarzban99/go-blog-api/internal/resources/admin"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService() *CategoryService {
	return &CategoryService{repo: repository.NewCategoryRepository()}
}

func (s *CategoryService) CategoriesList(dto *dtoAdmin.BaseAdminListDTO) (*repository.PaginatedResponse[model.Category], error) {
	return s.repo.CategoriesList(dto)
}

func (s *CategoryService) Show(id int) (*admin.CategoryResource, error) {
	return s.repo.FindOne(id)
}

func (s *CategoryService) Store(dto *dtoAdmin.StoreCategoryDTO) (*admin.CategoryResource, error) {
	return s.repo.Create(dto)
}

func (s *CategoryService) Update(id int, dto *dtoAdmin.UpdateCategoryDTO) error {
	return s.repo.Update(id, dto)
}

func (s *CategoryService) Destroy(id int) error {
	return s.repo.Destroy(id)
}
