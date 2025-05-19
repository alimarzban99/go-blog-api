package admin

import (
	dtoAdmin "github.com/alimarzban99/go-blog-api/internal/dtos/admin"
	"github.com/alimarzban99/go-blog-api/internal/service/admin"
	"github.com/alimarzban99/go-blog-api/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CategoryHandler struct {
	service *admin.CategoryService
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{service: admin.NewCategoryService()}
}

func (h *CategoryHandler) Index(ctx *gin.Context) {

	dto := new(dtoAdmin.BaseAdminListDTO)
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}
	dto.SetDefaults()

	categories, err := h.service.CategoriesList(dto)

	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.SuccessResponse(ctx, categories)
}

func (h *CategoryHandler) Show(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	category, err := h.service.Show(categoryId)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.SuccessResponse(ctx, category)
}

func (h *CategoryHandler) Store(ctx *gin.Context) {
	dto := new(dtoAdmin.StoreAndUpdateCategoryDTO)
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	result, err := h.service.Store(dto)

	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.CreatedResponse(ctx, result)
}

func (h *CategoryHandler) Update(ctx *gin.Context) {
	dto := new(dtoAdmin.StoreAndUpdateCategoryDTO)
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	err = h.service.Update(categoryId, dto)

	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.UpdateResponse(ctx, nil)
}

func (h *CategoryHandler) Destroy(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	err := h.service.Destroy(categoryId)

	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.DeletedResponse(ctx)
}
