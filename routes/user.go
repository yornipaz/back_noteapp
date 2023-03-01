package routes

import (
	"github.com/gin-gonic/gin"
	usercontroller "github.com/yornifpaz/back_noteapp/controllers/userController"
	"github.com/yornifpaz/back_noteapp/middleware"
)

func userRoutes(app *gin.RouterGroup) {
	userGroup := app.Group("/user")
	userGroup.Use(middleware.Authenticate())
	userGroup.PATCH("", usercontroller.Update)
	userGroup.PATCH("/avatar", usercontroller.UpdateAvatar)
	userGroup.PATCH("/status", usercontroller.UpdateStatus)
	userGroup.PATCH("/password", usercontroller.UpdatePassword)
}
