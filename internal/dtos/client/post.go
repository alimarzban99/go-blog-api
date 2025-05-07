package client

type PostsListDTO struct {
	Mobile string `json:"mobile" binding:"required,min=11,max=11"`
	Code   string `json:"code" binding:"required,min=4,max=4"`
}
