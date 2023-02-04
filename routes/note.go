package routes

import (
	"github.com/gin-gonic/gin"
	notecontroller "github.com/yornifpaz/back_noteapp/controllers/noteController"
	"github.com/yornifpaz/back_noteapp/middleware"
)

func noteRoutes(app *gin.RouterGroup) {
	noteGroup := app.Group("/notes")
	noteGroup.POST("", middleware.Authenticate, notecontroller.Create)
	noteGroup.GET("", middleware.Authenticate, notecontroller.GetAll)
	noteGroup.PATCH("/:id", middleware.Authenticate, notecontroller.Update)
	noteGroup.DELETE("/:id", middleware.Authenticate, notecontroller.Delete)

}
