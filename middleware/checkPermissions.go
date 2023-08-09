package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yornifpaz/back_noteapp/helpers"
	userrepository "github.com/yornifpaz/back_noteapp/repositories/userRepository"
	"gorm.io/gorm"
)

func CheckPermissions(permissionNames []string, db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := helpers.GetTokenStringFromHeader(ctx)
		// Verifica los permisos del usuario
		hasPermission := false
		for _, permissionName := range permissionNames {
			if checkUserPermission(userId, permissionName, db) {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Next()
	}
}
func checkUserPermission(userId string, permissionName string, db *gorm.DB) bool {
	userRepository := userrepository.NewUserRepository(db)
	user, err := userRepository.GetWithPermission(userId)
	if err != nil {
		return false
	}

	for _, role := range user.Roles {
		for _, permission := range role.Permissions {
			if permission.Name == permissionName {
				return true
			}
		}
	}

	return false
}
