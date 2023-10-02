package routes

import (
	"github.com/gin-gonic/gin"

	userfactory "github.com/yornifpaz/back_noteapp/app/factory/userFactory"
	usercontroller "github.com/yornifpaz/back_noteapp/app/http/controllers/userController"
	"github.com/yornifpaz/back_noteapp/app/http/middleware"
	userrepository "github.com/yornifpaz/back_noteapp/app/repositories/userRepository"

	"gorm.io/gorm"
)

type IRouteUser interface {
	setupRoutesUser()
}

// setup implements IRoute
func (r *Route) setupRoutesUser() {
	userRoutesGroup := r.router.Group(r.path)
	controller := getUserController(r.db)
	userRoutesGroup.Use(middleware.Authenticate())
	userRoutesGroup.PATCH("", controller.Update())
	userRoutesGroup.PATCH("/avatar", controller.UpdateAvatar())
	userRoutesGroup.PATCH("/status", controller.UpdateStatus())
	userRoutesGroup.PATCH("/password", controller.UpdatePassword())

}

// setup implements IRoute
func getUserController(db *gorm.DB) (controller usercontroller.IUserController) {
	repository := userrepository.NewUserRepository(db)
	factory := userfactory.NewUserFactory()
	controller = usercontroller.NewUserController(repository, factory)
	return
}

func newUserRoutes(router *gin.RouterGroup, db *gorm.DB,
	path string) IRouteUser {
	return &Route{
		router: router,
		db:     db,
		path:   path,
	}
}
