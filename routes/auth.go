package routes

import (
	"github.com/gin-gonic/gin"
	authcontrollers "github.com/yornifpaz/back_noteapp/controllers/authControllers"
	userfactory "github.com/yornifpaz/back_noteapp/factory/userFactory"
	"github.com/yornifpaz/back_noteapp/middleware"
	userrepository "github.com/yornifpaz/back_noteapp/repositories/userRepository"
	"gorm.io/gorm"
)

type IRouteAuth interface {
	setupRoutesAuth()
}

// setup implements IRoute
func (r *Route) setupRoutesAuth() {

	controller := getAuthController(r.db)
	authGroup := r.router.Group(r.path)
	authGroup.POST("/register", controller.Register())
	authGroup.POST("/login", controller.Login())
	authGroup.GET("/validate", middleware.Authenticate(), controller.Validate())
	authGroup.GET("/logout", middleware.Authenticate(), controller.Logout())

}

// setup implements IRoute
func getAuthController(db *gorm.DB) (controller authcontrollers.IAuthController) {
	repository := userrepository.NewUserRepository(db)
	factory := userfactory.NewUserFactory()
	controller = authcontrollers.NewAuthController(repository, factory)
	return
}

func newAuthRoutes(router *gin.RouterGroup, db *gorm.DB,
	path string) IRouteAuth {
	return &Route{
		router: router,
		db:     db,
		path:   path,
	}
}
