package repository

import (
	dtoAdmin "github.com/alimarzban99/go-blog-api/internal/dtos/admin"
	"github.com/alimarzban99/go-blog-api/internal/model"
	"github.com/alimarzban99/go-blog-api/internal/resources/admin"
	"github.com/alimarzban99/go-blog-api/pkg/database"
)

type CategoryRepository struct {
	*Repository[model.Category, dtoAdmin.StoreCategoryDTO, dtoAdmin.UpdateCategoryDTO, admin.CategoryResource]
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		&Repository[model.Category, dtoAdmin.StoreCategoryDTO, dtoAdmin.UpdateCategoryDTO, admin.CategoryResource]{
			database: database.GetDB(),
		},
	}
}

func (r *CategoryRepository) CategoriesList(dto *dtoAdmin.GetUserAdminListDTO) ([]admin.CategoryCollection, error) {

	var users []admin.CategoryCollection
	return users, nil
}
