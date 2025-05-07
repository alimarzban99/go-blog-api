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

func (r *PostRepository) PostsList(dto *dtoAdmin.GetUserAdminListDTO) ([]admin.PostCollection, error) {

	var users []admin.PostCollection
	return users, nil
}
