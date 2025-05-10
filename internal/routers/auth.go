package routers

import (
	"github.com/alimarzban99/go-blog-api/internal/handler/auth"
	"github.com/alimarzban99/go-blog-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup) {
	authRouter := r.Group("auth")

	authHandler := auth.NewAuthHandler()
	authRouter.POST("get-verification-code", authHandler.GetVerificationCode)
	authRouter.POST("verify", authHandler.Verify)
	authRouter.GET("logout", middlewares.Authentication("admin"), authHandler.Logout)
}
