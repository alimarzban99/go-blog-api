package middlewares

import (
	"github.com/alimarzban99/go-blog-api/internal/enums"
	"github.com/alimarzban99/go-blog-api/internal/repository"
	"github.com/alimarzban99/go-blog-api/pkg"
	"github.com/alimarzban99/go-blog-api/pkg/response"
	"github.com/gin-gonic/gin"
	"strings"
)

func Authentication(kind string) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		claimMap := map[string]interface{}{}
		auth := ctx.GetHeader("Authorization")
		if auth == "" {
			response.AuthenticationErrorResponse(ctx, "Authentication Is Required")
			return
		}

		token := strings.Split(auth, " ")
		claimMap, err := pkg.GetClaims(token[1])
		if err != nil {
			response.AuthenticationErrorResponse(ctx, "Invalid Token")
			return
		}

		repo := repository.NewUserRepository()

		userIDFloat, _ := claimMap["user_id"].(float64)
		user, err := repo.FindOne(int(userIDFloat))

		if (kind == "admin" && true != true) || (kind == "client" && enums.Status(user.Status) == enums.Active) {
			response.AuthorizationErrorResponse(ctx, "No Access Permission")
		}

		ctx.Set("user", user)
		ctx.Next()

	}
}
