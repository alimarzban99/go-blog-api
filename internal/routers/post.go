package routers

import (
	"github.com/alimarzban99/go-blog-api/internal/handler/admin"
	"github.com/alimarzban99/go-blog-api/internal/handler/client"
	"github.com/alimarzban99/go-blog-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func PostRouter(r *gin.RouterGroup) {
	adminRouter := r.Group("admin")
	{
		postRouter := adminRouter.Group("post").Use(middlewares.Authentication("admin"))
		postHandler := admin.NewPostHandler()
		postRouter.GET("", postHandler.Index)
		postRouter.GET(":id", postHandler.Show)
		postRouter.POST("", postHandler.Store)
		postRouter.PUT(":id", postHandler.Update)
		postRouter.DELETE(":id", postHandler.Destroy)
	}

	clientRouter := r.Group("client")
	{
		postRouter := clientRouter.Group("post")
		postHandler := client.NewPostHandler()
		postRouter.GET("", postHandler.Index)
		postRouter.PUT(":slug", postHandler.Show)
	}
}
