package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yornifpaz/back_noteapp/app/helpers"
	"github.com/yornifpaz/back_noteapp/lib"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtLibrary := lib.NewJwtLibrary()
		tokenString := helpers.GetTokenStringFromHeader(ctx)
		isValidToken := jwtLibrary.ValidateToken(tokenString)
		if !isValidToken {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Next()
	}

}
