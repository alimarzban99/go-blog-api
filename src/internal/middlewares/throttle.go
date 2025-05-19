package middlewares

import (
	"github.com/didip/tollbooth/v7"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Throttle() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(10.0/60.0, nil)
	lmt.SetTokenBucketExpirationTTL(time.Minute) // مدت زمان حفظ توکن‌ها
	lmt.SetBurst(10)
	return func(ctx *gin.Context) {

		err := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": err.Error()})
			return
		}

		ctx.Next()

	}
}
