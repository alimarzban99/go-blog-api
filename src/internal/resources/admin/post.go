package admin

type PostResource struct {
	UserID      uint    `json:"user_id"`
	CategoryID  uint    `json:"category_id"`
	Title       string  `json:"title"`
	Slug        string  `json:"slug"`
	Description *string `json:"description"`
	Image       string  `json:"image"`
	Hits        uint    `json:"hits"`
}
