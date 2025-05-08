package client

import (
	dtoAdmin "github.com/alimarzban99/go-blog-api/internal/dtos/admin"
	"github.com/alimarzban99/go-blog-api/internal/model"
	"github.com/alimarzban99/go-blog-api/internal/repository"
	"github.com/alimarzban99/go-blog-api/internal/resources/admin"
)

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService() *PostService {
	return &PostService{repo: repository.NewPostRepository()}
}

func (s *PostService) FilterList(dto *dtoAdmin.StorePostDTO) (*admin.PostResource, error) {
	return s.repo.Create(dto)
}

func (s *PostService) PostsList(dto *dtoAdmin.BaseAdminListDTO) (*repository.PaginatedResponse[model.Post], error) {
	return s.repo.ClientPostsList(dto)
}

func (s *PostService) Show(id int) (*admin.PostResource, error) {
	return s.repo.FindOne(id)
}
