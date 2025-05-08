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

func (r *CategoryRepository) CategoriesList(dto *dtoAdmin.BaseAdminListDTO) (*PaginatedResponse[model.Category], error) {

	query := r.database.Model(&model.Category{}).
		Select("id, title, created_at")

	if dto.Search != nil {
		query = query.Where("title LIKE ?", "%"+*dto.Search+"%")
	}

	if dto.CreatedAtFrom != nil {
		query = query.Where("created_at >= ?", *dto.CreatedAtFrom)
	}

	if dto.CreatedAtTo != nil {
		query = query.Where("created_at <= ?", *dto.CreatedAtTo)
	}

	query = r.OrderBY(query, *dto.Sort, *dto.Direction)

	return r.Paginate(query, *dto.Page, *dto.Limit)
}
