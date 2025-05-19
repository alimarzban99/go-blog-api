package admin

import (
	dtoAdmin "github.com/alimarzban99/go-blog-api/internal/dtos/admin"
	"github.com/alimarzban99/go-blog-api/internal/service/admin"
	"github.com/alimarzban99/go-blog-api/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserHandler struct {
	service *admin.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{service: admin.NewUserService()}
}

func (h *UserHandler) Index(ctx *gin.Context) {

	dto := new(dtoAdmin.BaseAdminListDTO)
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}
	dto.SetDefaults()

	users, err := h.service.UserList(dto)

	if err != nil {
		response.ErrorResponse(ctx, err.Error())
	}

	response.SuccessResponse(ctx, users)
}

func (h *UserHandler) Show(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	user, err := h.service.Show(userId)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
	}

	response.SuccessResponse(ctx, user)
}

func (h *UserHandler) Store(ctx *gin.Context) {
	dto := new(dtoAdmin.StoreUserDTO)
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
	}

	result, err := h.service.Store(dto)

	if err != nil {
		response.ErrorResponse(ctx, err.Error())
	}

	response.CreatedResponse(ctx, result)
}

func (h *UserHandler) Update(ctx *gin.Context) {
	dto := new(dtoAdmin.UpdateUserDTO)
	userId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
	}

	err = h.service.Update(userId, dto)

	if err != nil {
		response.ErrorResponse(ctx, err.Error())
	}

	response.UpdateResponse(ctx, nil)
}

func (h *UserHandler) Destroy(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	err := h.service.Destroy(userId)

	if err != nil {
		response.ErrorResponse(ctx, err.Error())
	}

	response.DeletedResponse(ctx)
}
