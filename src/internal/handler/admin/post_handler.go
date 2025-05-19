package admin

import (
	dtoAdmin "github.com/alimarzban99/go-blog-api/internal/dtos/admin"
	"github.com/alimarzban99/go-blog-api/internal/service/admin"
	"github.com/alimarzban99/go-blog-api/pkg/filer"
	"github.com/alimarzban99/go-blog-api/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type PostHandler struct {
	service *admin.PostService
	filer   *filer.Filer
}

func NewPostHandler() *PostHandler {
	return &PostHandler{service: admin.NewPostService(), filer: filer.NewFiler()}
}

func (h *PostHandler) Index(ctx *gin.Context) {
	dto := new(dtoAdmin.BaseAdminListDTO)
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}
	dto.SetDefaults()

	posts, err := h.service.PostsList(dto)

	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
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
	dto := new(dtoAdmin.StoreAndUpdatePostDTO)
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
	dto := new(dtoAdmin.StoreAndUpdatePostDTO)
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

func (h *PostHandler) Upload(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		response.ErrorResponse(ctx, "No file received")
		return
	}

	file, err := h.filer.Uploader(fileHeader)

	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.CreatedResponse(ctx, file)
}
