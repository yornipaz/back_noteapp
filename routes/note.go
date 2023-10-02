package routes

import (
	"github.com/gin-gonic/gin"

	notefactory "github.com/yornifpaz/back_noteapp/app/factory/noteFactory"
	notecontroller "github.com/yornifpaz/back_noteapp/app/http/controllers/noteController"
	"github.com/yornifpaz/back_noteapp/app/http/middleware"
	noterepository "github.com/yornifpaz/back_noteapp/app/repositories/noteRepository"

	"gorm.io/gorm"
)

type IRouteNote interface {
	setupRoutesNote()
}

// setup implements IRoute
func (r *Route) setupRoutesNote() {
	controller := getNoteController(r.db)
	r.router.Use(middleware.Authenticate())
	r.router.GET("", controller.GetAll())
	r.router.POST("", controller.Create())
	r.router.PATCH("", controller.Update())
	r.router.PATCH("/archived", controller.UpdateArchived())
	r.router.DELETE("/:id", controller.Delete())

}

// setup implements IRoute
func getNoteController(db *gorm.DB) (controller notecontroller.INoteController) {
	repository := noterepository.NewNoteRepository(db)
	factory := notefactory.NewNoteFactory()
	controller = notecontroller.NewNoteController(repository, factory)
	return
}

func newNoteRoutes(router *gin.RouterGroup, db *gorm.DB,
	path string) IRouteNote {
	return &Route{
		router: router,
		db:     db,
		path:   path,
	}
}
