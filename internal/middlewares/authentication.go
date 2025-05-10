package middlewares

import (
	"crypto/rsa"
	"fmt"
	"github.com/alimarzban99/go-blog-api/internal/repository"
	"github.com/alimarzban99/go-blog-api/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"strings"
)

var publicKey *rsa.PublicKey

func loadPublicKey() {
	keyData, err := os.ReadFile("keys/public.pem")
	if err != nil {
		log.Fatalf("Failed to read public key: %v", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		log.Fatalf("Failed to parse public key: %v", err)
	}

	publicKey = key
}

func Authentication(kind string) gin.HandlerFunc {

	loadPublicKey()
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") || len(authHeader) < 8 {
			response.AuthenticationErrorResponse(ctx, "Authentication required")
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return publicKey, nil
		})

		if err != nil || !token.Valid {
			response.AuthenticationErrorResponse(ctx, "Unauthorized")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.AuthenticationErrorResponse(ctx, "invalid token")
			return
		}

		jti := claims["jti"].(string)

		repo := repository.NewTokenRepository()
		tokenExists, err := repo.FindToken(jti)
		if err != nil || !tokenExists {
			response.AuthenticationErrorResponse(ctx, "invalid token")
			return
		}

		userID := int(claims["sub"].(float64))
		userRepo := repository.NewUserRepository()
		user, err := userRepo.FindOne(userID)

		ctx.Set("user", user)
		ctx.Set("jti", jti)
		ctx.Next()

	}
}
