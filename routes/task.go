package routes

import (
	"github.com/gin-gonic/gin"

	taskfactory "github.com/yornifpaz/back_noteapp/app/factory/taskFactory"
	taskcontroller "github.com/yornifpaz/back_noteapp/app/http/controllers/taskController"
	"github.com/yornifpaz/back_noteapp/app/http/middleware"
	noterepository "github.com/yornifpaz/back_noteapp/app/repositories/noteRepository"
	taskrepository "github.com/yornifpaz/back_noteapp/app/repositories/taskRepository"
	"github.com/yornifpaz/back_noteapp/lib"

	"gorm.io/gorm"
)

type IRouteTask interface {
	setupRoutesTask()
}

// setup implements IRoute
func (r *Route) setupRoutesTask() {
	controller := getTaskController(r.db)
	taskGroup := r.router.Group(r.path)
	taskGroup.Use(middleware.Authenticate())
	taskGroup.POST("", controller.Create())
	taskGroup.PATCH("", controller.Update())
	taskGroup.DELETE("/:id", controller.Delete())

}

// setup implements IRoute
func getTaskController(db *gorm.DB) (controller taskcontroller.ITaskController) {
	repository := taskrepository.NewTaskRepository(db)
	noteRepository := noterepository.NewNoteRepository(db)
	idLibrary := lib.NewIdLibrary()
	factory := taskfactory.NewTaskFactory(idLibrary)
	controller = taskcontroller.NewTaskController(repository, factory, noteRepository)
	return
}

func newTaskRoutes(router *gin.RouterGroup, db *gorm.DB,
	path string) IRouteTask {
	return &Route{
		router: router,
		db:     db,
		path:   path,
	}
}
