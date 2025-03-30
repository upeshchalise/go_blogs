package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/upeshchalise/go_blogs/pkg/utils/jwt"
)

func VerifyJWT(Next gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(401, gin.H{"error": "Authorization header is missing"})
			ctx.Abort()
			return
		}

		if !strings.HasPrefix(token, "Bearer ") {
			ctx.JSON(401, gin.H{"error": "Authorization header format must be Bearer {token}"})
			ctx.Abort()
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")

		token, err := jwt.ValidateJwtToken(token, "thisissecrettoken")
		if err != nil {
			ctx.JSON(401, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}
		ctx.Set("userId", token)
		Next(ctx)
	}
}
