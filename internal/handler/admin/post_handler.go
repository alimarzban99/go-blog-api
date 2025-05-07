package admin

import (
	dtoAdmin "github.com/alimarzban99/go-blog-api/internal/dtos/admin"
	"github.com/alimarzban99/go-blog-api/internal/service/admin"
	"github.com/alimarzban99/go-blog-api/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type PostHandler struct {
	service *admin.PostService
}

func NewPostHandler() *PostHandler {
	return &PostHandler{service: admin.NewPostService()}
}

func (h *PostHandler) Index(ctx *gin.Context) {

	dto := new(dtoAdmin.GetUserAdminListDTO)
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}
	dto.SetDefaults()

	posts, err := h.service.PostsList(dto)

	if err != nil {
		response.ErrorResponse(ctx, err.Error())
	}

	response.SuccessResponse(ctx, posts)
}

func (h *PostHandler) Show(ctx *gin.Context) {
	PostId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	post, err := h.service.Show(PostId)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
	}

	response.SuccessResponse(ctx, post)
}

func (h *PostHandler) Store(ctx *gin.Context) {
	dto := new(dtoAdmin.StorePostDTO)
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	result, err := h.service.Store(dto)

	if err != nil {
		response.ErrorResponse(ctx, err.Error())
	}

	response.CreatedResponse(ctx, result)
}

func (h *PostHandler) Update(ctx *gin.Context) {
	dto := new(dtoAdmin.UpdatePostDTO)
	PostId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	err = h.service.Update(PostId, dto)

	if err != nil {
		response.ErrorResponse(ctx, err.Error())
	}

	response.UpdateResponse(ctx, nil)
}

func (h *PostHandler) Destroy(ctx *gin.Context) {
	PostId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	err := h.service.Destroy(PostId)

	if err != nil {
		response.ErrorResponse(ctx, err.Error())
	}

	response.UpdateResponse(ctx, nil)
}
