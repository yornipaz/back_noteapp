package routes

import (
	"github.com/gin-gonic/gin"
	userfactory "github.com/yornifpaz/back_noteapp/app/factory/userFactory"
	authcontrollers "github.com/yornifpaz/back_noteapp/app/http/controllers/authControllers"
	"github.com/yornifpaz/back_noteapp/app/http/middleware"
	authrepository "github.com/yornifpaz/back_noteapp/app/repositories/authRepository"
	userrepository "github.com/yornifpaz/back_noteapp/app/repositories/userRepository"
	"github.com/yornifpaz/back_noteapp/config"
	"github.com/yornifpaz/back_noteapp/lib"
	"github.com/yornifpaz/back_noteapp/services"
	"github.com/yornifpaz/back_noteapp/templates"

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
	authGroup.POST("/forgot-password", controller.ForgotPassword())
	authGroup.PATCH("/reset-password/:resetToken", controller.ResetPassword())
	authGroup.GET("/validate", middleware.Authenticate(), controller.Validate())
	authGroup.GET("/logout", middleware.Authenticate(), controller.Logout())

}

// setup implements IRoute
func getAuthController(db *gorm.DB) (controller authcontrollers.IAuthController) {
	repository := userrepository.NewUserRepository(db)

	factory := userfactory.NewUserFactory()
	configEmail := config.NewConfigurationApplication().GetDefaultEmailConfig()
	emailLibrary := lib.NewEmailLibrary(lib.EmailConfig(configEmail))
	templates := templates.NewEmailTemplate()
	emailService := services.NewEmailService(emailLibrary, templates)
	encryptLibrary := lib.NewEncryptLibrary()
	JwtLibrary := lib.NewJwtLibrary()
	authRepository := authrepository.NewAuthRepository(JwtLibrary, encryptLibrary)
	controller = authcontrollers.NewAuthController(repository, factory, authRepository, emailService)
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
