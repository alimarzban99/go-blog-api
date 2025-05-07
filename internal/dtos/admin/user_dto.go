package admin

import "time"

type StoreUserDTO struct {
	FirstName *string `json:"first_name" binding:"required"`
	LastName  *string `json:"last_name" binding:"required"`
	Mobile    *string `json:"mobile" binding:"required"`
	IsAdmin   bool    `json:"is_admin"`
	Email     *string `json:"email" binding:"required,email"`
	Status    *string `json:"status" binding:"required"`
}

type UpdateUserDTO struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Mobile    string `json:"mobile" binding:"required"`
	IsAdmin   bool   `json:"is_admin"`
	Email     string `json:"email" binding:"required"`
	Status    string `json:"status" binding:"required"`
}

type GetUserAdminListDTO struct {
	Search *string `form:"search" binding:"omitempty,min=3"`
	//TODO// fix validate dateTime
	CreatedAtFrom *time.Time `form:"created_at_from" binding:"omitempty,datetime=2006-01-02"`
	CreatedAtTo   *time.Time `form:"created_at_to" binding:"omitempty,datetime=2006-01-02"`
	Limit         *int       `form:"limit"`
	Sort          *string    `form:"sort" binding:"required,oneof=id created_at"`
	Direction     *string    `form:"direction" binding:"required,oneof=asc desc"`
	Page          *int       `form:"page"`
}

func (dto *GetUserAdminListDTO) SetDefaults() {
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
