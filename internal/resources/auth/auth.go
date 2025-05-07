package auth

type CodeResponse struct {
	Code string `json:"code" binding:"required,min=4,max=4"`
}
