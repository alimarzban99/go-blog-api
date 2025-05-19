package auth

import (
	authdto "github.com/alimarzban99/go-blog-api/internal/dtos/auth"
	"github.com/alimarzban99/go-blog-api/internal/service/auth"
	"github.com/alimarzban99/go-blog-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *auth.Service
}

func NewAuthHandler() *Handler {
	return &Handler{service: auth.NewAuthService()}
}

func (h *Handler) GetVerificationCode(ctx *gin.Context) {
	dto := new(authdto.GetOTPCodeDTO)
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}
	h.service.GetVerificationCode(dto)
	response.SuccessResponse(ctx, "Code sent")
}

func (h *Handler) Verify(ctx *gin.Context) {
	dto := new(authdto.VerifyOTPCodeDTO)
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}
	token, err := h.service.Verify(dto)
	if err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}
	response.SuccessResponse(ctx, struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	})
}

func (h *Handler) Logout(ctx *gin.Context) {
	jti := ctx.MustGet("jti").(string)
	h.service.Logout(jti)
	response.SuccessResponse(ctx, "Logged out")
}
