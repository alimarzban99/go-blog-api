package repository

import (
	dtoAdmin "github.com/alimarzban99/go-blog-api/internal/dtos/admin"
	"github.com/alimarzban99/go-blog-api/internal/model"
	"github.com/alimarzban99/go-blog-api/internal/resources/admin"
	"github.com/alimarzban99/go-blog-api/pkg/database"
)

type UserRepository struct {
	*Repository[model.User, dtoAdmin.StoreUserDTO, dtoAdmin.UpdateUserDTO, admin.UserResource]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		&Repository[model.User, dtoAdmin.StoreUserDTO, dtoAdmin.UpdateUserDTO, admin.UserResource]{
			database: database.GetDB(),
		},
	}
}

func (r UserRepository) AdminList(dto *dtoAdmin.BaseAdminListDTO) (*PaginatedResponse[model.User], error) {

	query := r.database.Model(&model.User{}).
		Select("id, first_name, last_name, mobile, email, is_admin, created_at")

	if dto.Search != nil {
		query = query.Where("first_name LIKE ?", "%"+*dto.Search+"%")
	}

	if dto.Search != nil {
		query = query.Where("last_name LIKE ?", "%"+*dto.Search+"%")
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

func (r UserRepository) CheckExistsByMobile(mobile string) (bool, error) {
	var exists bool

	err := r.database.
		Model(&model.User{}).
		Select("count(*) > 0").
		Where("mobile=?", mobile).
		Find(&exists).
		Error

	if err != nil {
		return false, nil
	}

	return exists, nil
}

func (r UserRepository) FindByMobile(mobile string) (*admin.UserResource, error) {

	res := &admin.UserResource{}
	err := r.database.
		Model(model.User{}).
		Where("mobile=?", mobile).
		Find(res).
		Error

	return res, err
}
