package admin

import (
	dtoAdmin "github.com/alimarzban99/go-blog-api/internal/dtos/admin"
	"github.com/alimarzban99/go-blog-api/internal/repository"
	"github.com/alimarzban99/go-blog-api/internal/resources/admin"
)

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService() *PostService {
	return &PostService{repo: repository.NewPostRepository()}
}

func (s *PostService) PostsList(dto *dtoAdmin.GetUserAdminListDTO) ([]admin.PostCollection, error) {
	return s.repo.PostsList(dto)
}

func (s *PostService) Show(id int) (*admin.PostResource, error) {
	return s.repo.FindOne(id)
}

func (s *PostService) Store(dto *dtoAdmin.StorePostDTO) (*admin.PostResource, error) {
	return s.repo.Create(dto)
}

func (s *PostService) Update(id int, dto *dtoAdmin.UpdatePostDTO) error {
	return s.repo.Update(id, dto)
}

func (s *PostService) Destroy(id int) error {
	return s.repo.Destroy(id)
}
