package repository

import (
	dtoAdmin "github.com/alimarzban99/go-blog-api/internal/dtos/admin"
	"github.com/alimarzban99/go-blog-api/internal/model"
	"github.com/alimarzban99/go-blog-api/internal/resources/admin"
	"github.com/alimarzban99/go-blog-api/pkg/database"
)

type PostRepository struct {
	*Repository[model.Post, dtoAdmin.StorePostDTO, dtoAdmin.UpdatePostDTO, admin.PostResource]
}

func NewPostRepository() *PostRepository {
	return &PostRepository{
		&Repository[model.Post, dtoAdmin.StorePostDTO, dtoAdmin.UpdatePostDTO, admin.PostResource]{
			database: database.GetDB(),
		},
	}
}

func (r *PostRepository) AdminPostsList(dto *dtoAdmin.BaseAdminListDTO) (*PaginatedResponse[model.Post], error) {

	query := r.database.Model(&model.Post{}).
		Select("id, title, slug, description, email, hits, category_id, user_id, created_at")

	if dto.Search != nil {
		query = query.Where("title LIKE ?", "%"+*dto.Search+"%")
	}

	if dto.Search != nil {
		query = query.Where("slug LIKE ?", "%"+*dto.Search+"%")
	}

	if dto.Search != nil {
		query = query.Where("description LIKE ?", "%"+*dto.Search+"%")
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
func (r *PostRepository) ClientPostsList(dto *dtoAdmin.BaseAdminListDTO) (*PaginatedResponse[model.Post], error) {

	query := r.database.Model(&model.Post{}).
		Select("id, title, slug, description, email, hits, category_id, user_id, created_at")

	query = r.OrderBY(query, *dto.Sort, *dto.Direction)

	return r.Paginate(query, *dto.Page, *dto.Limit)
}
