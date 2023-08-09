package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/yornifpaz/back_noteapp/lib"
)

/**
*@package helpers
* function jwt created and validates token
 */

func GetTokenStringFromHeader(ctx *gin.Context) (tokenString string) {
	const BEARER_SCHEMA = "Bearer"
	authHeader := ctx.GetHeader("Authorization")
	tokenString = authHeader[len(BEARER_SCHEMA):]
	return
}

/*
	GetCurrentUserId

function  get  current user id  from the current token

	@return userId string
*/
func GetCurrentUserId(ctx *gin.Context) (userId string) {
	tokenString := GetTokenStringFromHeader(ctx)
	jwtLibrary := lib.NewJwtLibrary()
	token, _ := jwtLibrary.Parse(tokenString)
	userId = jwtLibrary.GetClaims(token)["sub"].(string)
	return
}
