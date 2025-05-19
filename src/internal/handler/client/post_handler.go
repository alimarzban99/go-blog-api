package client

import (
	"github.com/alimarzban99/go-blog-api/internal/service/client"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service *client.PostService
}

func NewPostHandler() *PostHandler {
	return &PostHandler{service: client.NewPostService()}
}
func (h *PostHandler) FilterList(c *gin.Context) {}

func (h *PostHandler) Index(c *gin.Context) {
}

func (h *PostHandler) Show(c *gin.Context) {}
