package middlewares

import (
	"github.com/alimarzban99/go-blog-api/config"
	"github.com/alimarzban99/go-blog-api/pkg/logging"
	"github.com/alimarzban99/go-blog-api/pkg/response"
	"github.com/gin-gonic/gin"
)

func CustomRecovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				var errMessage string
				env := config.Config.App.Env

				if env == "production" {
					errMessage = "خطای داخلی سرور رخ داد. لطفاً بعداً دوباره تلاش کنید."
				} else {
					switch v := r.(type) {
					case string:
						errMessage = v
					case error:
						errMessage = v.Error()
					default:
						errMessage = "خطای ناشناخته‌ای رخ داد"
					}
				}
				logger := logging.NewLogger()
				logger.Error(logging.Recovery, errMessage)

				response.PanicResponse(ctx, errMessage)
			}
		}()
		ctx.Next()
	}
}
