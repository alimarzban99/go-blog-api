package routers

import (
	"github.com/alimarzban99/go-blog-api/internal/handler/admin"
	"github.com/alimarzban99/go-blog-api/internal/handler/client"
	"github.com/alimarzban99/go-blog-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	adminRouter := r.Group("admin")
	{
		userRouter := adminRouter.Group("user").Use(middlewares.Authentication("admin"))
		userHandler := admin.NewUserHandler()
		userRouter.GET("", userHandler.Index)
		userRouter.GET(":id", userHandler.Show)
		userRouter.POST("", userHandler.Store)
		userRouter.PUT(":id", userHandler.Update)
		userRouter.DELETE(":id", userHandler.Destroy)
	}

	clientRouter := r.Group("client")
	{
		userRouter := clientRouter.Group("user").Use(middlewares.Authentication("client"))
		userHandler := client.NewUserHandler()
		userRouter.GET("profile", userHandler.Profile)
		userRouter.PUT("", userHandler.Update)
	}
}
