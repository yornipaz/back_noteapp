package routes

import (
	"github.com/gin-gonic/gin"
	authcontrollers "github.com/yornifpaz/back_noteapp/controllers/authControllers"
	"github.com/yornifpaz/back_noteapp/middleware"
)

func authRoutes(app *gin.RouterGroup) {
	authGroup := app.Group("/auth")

	authGroup.POST("/register", authcontrollers.Register)
	authGroup.POST("/login", authcontrollers.Login)
	authGroup.GET("/validate", middleware.Authenticate, authcontrollers.Validate)
	authGroup.GET("/logout", middleware.Authenticate, authcontrollers.Logout)

}
