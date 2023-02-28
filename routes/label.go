package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yornifpaz/back_noteapp/config"
	labelcontrollers "github.com/yornifpaz/back_noteapp/controllers/labelControllers"
	"github.com/yornifpaz/back_noteapp/middleware"
	labelrepository "github.com/yornifpaz/back_noteapp/repositories/labelRepository"
)

func labelRoutes(app *gin.RouterGroup) {

	db := config.DB
	repository := labelrepository.NewLabelRepository(db)
	controller := labelcontrollers.NewLabelController(repository)
	labelGroup := app.Group("/notes/label")
	labelGroup.Use(middleware.Authenticate)

	labelGroup.POST("", controller.Create())
	labelGroup.GET("/all", controller.GetAll())
	labelGroup.PATCH("", controller.Update())
	labelGroup.DELETE("/:id", controller.Delete())

}
