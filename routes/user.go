package routes

import (
	"github.com/gin-gonic/gin"
	usercontroller "github.com/yornifpaz/back_noteapp/controllers/userController"
	"github.com/yornifpaz/back_noteapp/middleware"
)

func userRoutes(app *gin.RouterGroup) {
	userGroup := app.Group("user")
	userGroup.POST("/register", usercontroller.Register)
	userGroup.POST("/login", usercontroller.Login)
	userGroup.GET("/validate", middleware.Authenticate, usercontroller.Validate)
	userGroup.GET("/logout", middleware.Authenticate, usercontroller.Logout)

}
