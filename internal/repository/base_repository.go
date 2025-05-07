package repository

import (
	"github.com/alimarzban99/go-blog-api/pkg/converter"
	"gorm.io/gorm"
)

type Repository[Model any, CrDTO any, UpDTO any, ResSingle any] struct {
	database *gorm.DB
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
