package routers

import (
	"github.com/alimarzban99/go-blog-api/internal/handler/admin"
	"github.com/alimarzban99/go-blog-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func CategoryRouter(r *gin.RouterGroup) {
	adminRouter := r.Group("admin")
	{
		categoryRouter := adminRouter.Group("category").Use(middlewares.Authentication("admin"))
		categoryHandler := admin.NewCategoryHandler()
		categoryRouter.GET("", categoryHandler.Index)
		categoryRouter.GET(":id", categoryHandler.Show)
		categoryRouter.POST("", categoryHandler.Store)
		categoryRouter.PUT(":id", categoryHandler.Update)
		categoryRouter.DELETE(":id", categoryHandler.Destroy)
	}
}
