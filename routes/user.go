package routes

import (
	"github.com/gin-gonic/gin"
	usercontroller "github.com/yornifpaz/back_noteapp/controllers/userController"
	"github.com/yornifpaz/back_noteapp/middleware"
)

func authRoutes(app *gin.RouterGroup) {
	authGroup := app.Group("/auth")

	authGroup.POST("/register", usercontroller.Register)
	authGroup.POST("/login", usercontroller.Login)
	authGroup.GET("/validate", middleware.Authenticate, usercontroller.Validate)
	authGroup.GET("/logout", middleware.Authenticate, usercontroller.Logout)

}
func userRoutes(app *gin.RouterGroup) {
	userGroup := app.Group("/user")
	userGroup.Use(middleware.Authenticate)
	userGroup.PATCH("", usercontroller.Update)
	userGroup.PATCH("/avatar", usercontroller.UpdateAvatar)
}
