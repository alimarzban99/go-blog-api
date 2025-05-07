package client

type UserUpdateProfileDTO struct {
	Mobile string `json:"mobile" binding:"required,min=11,max=11"`
}
