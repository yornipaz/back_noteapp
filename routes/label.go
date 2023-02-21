package routes

import (
	"github.com/gin-gonic/gin"
	labelcontrollers "github.com/yornifpaz/back_noteapp/controllers/labelControllers"
	"github.com/yornifpaz/back_noteapp/middleware"
)

func labelRoutes(app *gin.RouterGroup) {
	labelGroup := app.Group("/notes/label")
	labelGroup.Use(middleware.Authenticate)
	labelGroup.POST("", labelcontrollers.Create)
	labelGroup.GET("/all", labelcontrollers.GetAll)
	labelGroup.PATCH("", labelcontrollers.Update)
	labelGroup.DELETE("/:id", labelcontrollers.Delete)

}
