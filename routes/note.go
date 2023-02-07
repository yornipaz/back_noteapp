package routes

import (
	"github.com/gin-gonic/gin"
	notecontroller "github.com/yornifpaz/back_noteapp/controllers/noteController"
	"github.com/yornifpaz/back_noteapp/middleware"
)

func noteRoutes(app *gin.RouterGroup) {
	noteGroup := app.Group("/notes")
	noteGroup.Use(middleware.Authenticate)
	noteGroup.GET("", notecontroller.GetAll)
	noteGroup.POST("", notecontroller.Create)
	noteGroup.PATCH("/:id", notecontroller.Update)
	noteGroup.DELETE("/:id", notecontroller.Delete)

}
