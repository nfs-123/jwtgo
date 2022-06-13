package middleware

import (
	auth "jwtgo/helpers"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenstring := ctx.GetHeader("Authorization")
		if tokenstring == "" {
			ctx.JSON(401, gin.H{
				"error": "request does not contain access token",
			})
			return
		}
		err := auth.ValidateToken(tokenstring)
		if err != nil {
			ctx.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.Next()
	}
}
