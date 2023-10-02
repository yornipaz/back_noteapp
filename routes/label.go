package routes

import (
	"github.com/gin-gonic/gin"

	labelfactory "github.com/yornifpaz/back_noteapp/app/factory/labelFactory"
	labelcontrollers "github.com/yornifpaz/back_noteapp/app/http/controllers/labelControllers"
	"github.com/yornifpaz/back_noteapp/app/http/middleware"
	labelrepository "github.com/yornifpaz/back_noteapp/app/repositories/labelRepository"

	"gorm.io/gorm"
)

type IRouteLabel interface {
	setupRoutesLabel()
}

// setup implements IRoute
func (r *Route) setupRoutesLabel() {
	controller := getLabelController(r.db)
	labelGroup := r.router.Group(r.path)
	labelGroup.Use(middleware.Authenticate())
	labelGroup.POST("", controller.Create())
	labelGroup.GET("/all", controller.GetAll())
	labelGroup.PATCH("", controller.Update())
	labelGroup.DELETE("/:id", controller.Delete())

}

// setup implements IRoute
func getLabelController(db *gorm.DB) (controller labelcontrollers.ILabelController) {
	repository := labelrepository.NewLabelRepository(db)
	factory := labelfactory.NewLabelFactory()
	controller = labelcontrollers.NewLabelController(repository, factory)
	return
}

func newLabelRoutes(router *gin.RouterGroup, db *gorm.DB,
	path string) IRouteLabel {
	return &Route{
		router: router,
		db:     db,
		path:   path,
	}
}
