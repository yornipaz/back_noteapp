package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yornifpaz/back_noteapp/helpers"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		const BEARER_SCHEMA = "Bearer"
		authHeader := ctx.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		isValidToken := helpers.IsValidToken(tokenString)

		if isValidToken {

			ctx.Next()
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}

}
