package admin

type StoreAndUpdatePostDTO struct {
	UserID      *int    `json:"user_id"`
	CategoryID  *int    `json:"category_id" binding:"required"`
	Title       *string `json:"title" binding:"required"`
	Slug        *string `json:"slug" binding:"required"`
	Description *string `json:"description" binding:"required"`
	Image       string  `json:"image" binding:"required"`
	Status      string  `json:"status" binding:"required"`
}
