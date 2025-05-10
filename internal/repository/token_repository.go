package repository

import (
	"github.com/alimarzban99/go-blog-api/internal/dtos/auth"
	"github.com/alimarzban99/go-blog-api/internal/model"
	authResources "github.com/alimarzban99/go-blog-api/internal/resources/auth"
	"github.com/alimarzban99/go-blog-api/pkg/database"
)

type TokenRepository struct {
	*Repository[model.Token, auth.TokenCreate, auth.TokenCreate, authResources.TokenResponse]
}

func NewTokenRepository() *TokenRepository {
	return &TokenRepository{
		&Repository[model.Token, auth.TokenCreate, auth.TokenCreate, authResources.TokenResponse]{
			database: database.GetDB(),
		},
	}
}

func (r *TokenRepository) FindToken(jti string) (bool, error) {
	var exists bool
	err := r.database.
		Model(&model.Token{}).
		Select("count(*) > 0").
		Where("id = ? AND revoked = 0", jti).
		Find(&exists).Error
	return exists, err
}

func (r *TokenRepository) ExpiredToken(jti string) {
	r.database.
		Model(&model.Token{}).
		Where("id = ?", jti).
		Update("revoked", true)
}
