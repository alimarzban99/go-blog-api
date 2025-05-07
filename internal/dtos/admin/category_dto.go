package admin

type StoreCategoryDTO struct {
	FirstName *string `json:"first_name" binding:"required"`
	LastName  *string `json:"last_name" binding:"required"`
	Mobile    *string `json:"mobile" binding:"required"`
	IsAdmin   bool    `json:"is_admin"`
	Email     *string `json:"email" binding:"required,email"`
	Status    *string `json:"status" binding:"required"`
}

type UpdateCategoryDTO struct {
	FirstName string `json:"first_name" binding:"omitempty,min=3"`
	LastName  string `json:"last_name" binding:"omitempty,datetime=2006-01-02"`
	Mobile    string `json:"mobile" binding:"omitempty,datetime=2006-01-02"`
	IsAdmin   bool   `json:"is_admin"`
	Email     string `json:"email" binding:"required,oneof=id created_at"`
	Status    string `json:"status" binding:"required,oneof=asc desc"`
}
