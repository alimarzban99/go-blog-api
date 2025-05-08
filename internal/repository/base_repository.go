package repository

import (
	"github.com/alimarzban99/go-blog-api/pkg/converter"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
)

type Repository[Model any, CrDTO any, UpDTO any, ResSingle any] struct {
	database *gorm.DB
}

type PaginatedResponse[T any] struct {
	Data            []T   `json:"data"`
	BaseData        []T   `json:"base_data"`
	Total           int64 `json:"total"`
	PerPage         int   `json:"per_page"`
	CurrentPage     int   `json:"current_page"`
	LastPage        int   `json:"last_page"`
	From            int   `json:"from"`
	To              int   `json:"to"`
	FirstPage       int   `json:"first_page"`
	HasNextPage     bool  `json:"has_next_page"`
	HasPreviousPage bool  `json:"has_previous_page"`
}

func (r *Repository[Model, CrDTO, UpDTO, ResSingle]) FindOne(id int) (*ResSingle, error) {
	var model Model

	err := r.database.
		Where("id=?", id).
		First(&model).
		Error

	if err != nil {
		return nil, err
	}

	return converter.TypeConverter[ResSingle](model)
}

func (r *Repository[Model, CrDTO, UpDTO, ResSingle]) Create(CreateDTO *CrDTO) (*ResSingle, error) {

	model, _ := converter.TypeConverter[Model](CreateDTO)

	err := r.database.
		Create(&model).
		Error

	if err != nil {
		return nil, err
	}
	return converter.TypeConverter[ResSingle](model)
}

func (r *Repository[Model, CrDTO, UpDTO, ResSingle]) Update(id int, UpdateDTO *UpDTO) error {

	updateMap, _ := converter.TypeConverter[map[string]interface{}](UpdateDTO)
	model := new(Model)

	err := r.database.
		Model(model).
		Where("id=?", id).
		Updates(*updateMap).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository[Model, CrDTO, UpDTO, ResSingle]) Destroy(id int) error {
	model := new(Model)
	err := r.database.Model(model).Where("id = ?", id).Delete(&model).Error
	return err
}

func (r *Repository[Model, CrDTO, UpDTO, ResSingle]) OrderBY(query *gorm.DB, sort string, direction string) *gorm.DB {

	sortDirection := direction

	query = query.Order(clause.OrderByColumn{
		Column: clause.Column{Name: sort},
		Desc:   sortDirection == "desc",
	})

	return query
}

func (r *Repository[Model, CrDTO, UpDTO, ResSingle]) Paginate(db *gorm.DB, page, limit int) (*PaginatedResponse[Model], error) {
	var items []Model
	var total int64

	offset := (page - 1) * limit

	err := db.Count(&total).
		Offset(offset).
		Limit(limit).
		Find(&items).Error

	if err != nil {
		return nil, err
	}

	lastPage := int(math.Ceil(float64(total) / float64(limit)))
	if lastPage == 0 {
		lastPage = 1
	}

	from := offset + 1
	to := offset + len(items)

	result := &PaginatedResponse[Model]{
		Data:            items,
		Total:           total,
		PerPage:         limit,
		CurrentPage:     page,
		LastPage:        lastPage,
		From:            from,
		To:              to,
		FirstPage:       1,
		HasNextPage:     page < lastPage,
		HasPreviousPage: page > 1,
	}

	return result, nil
}
