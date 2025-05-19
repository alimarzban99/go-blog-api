package admin

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
