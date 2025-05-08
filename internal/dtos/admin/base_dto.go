package admin

import "time"

type BaseAdminListDTO struct {
	Search *string `form:"search" binding:"omitempty,min=3"`
	//TODO// fix validate dateTime
	CreatedAtFrom *time.Time `form:"created_at_from" binding:"omitempty,datetime=2006-01-02"`
	CreatedAtTo   *time.Time `form:"created_at_to" binding:"omitempty,datetime=2006-01-02"`
	Limit         *int       `form:"limit"`
	Sort          *string    `form:"sort" binding:"required,oneof=id created_at"`
	Direction     *string    `form:"direction" binding:"required,oneof=asc desc"`
	Page          *int       `form:"page"`
}

func (dto *BaseAdminListDTO) SetDefaults() {
	if dto.Limit == nil {
		defaultLimit := 50
		dto.Limit = &defaultLimit
	}
	if dto.Sort == nil {
		defaultSort := "id"
		dto.Sort = &defaultSort
	}
	if dto.Direction == nil {
		defaultDir := "desc"
		dto.Direction = &defaultDir
	}
	if dto.Page == nil {
		defaultPage := 1
		dto.Page = &defaultPage
	}
}
