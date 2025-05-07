package client

import (
	"github.com/alimarzban99/go-blog-api/internal/service/client"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *client.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{service: client.NewUserService()}
}
func (h *UserHandler) Profile(c *gin.Context) {
}

func (h *UserHandler) Update(c *gin.Context) {}
