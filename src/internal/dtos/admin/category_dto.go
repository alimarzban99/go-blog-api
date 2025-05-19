package admin

type StoreAndUpdateCategoryDTO struct {
	Title  *string `json:"title" binding:"required"`
	Status *string `json:"status" binding:"required"`
}
