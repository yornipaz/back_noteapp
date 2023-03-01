package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yornifpaz/back_noteapp/helpers"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("Authorization")
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		isValidToken := helpers.IsValidToken(tokenString)

		if isValidToken {

			ctx.Next()
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}

}
