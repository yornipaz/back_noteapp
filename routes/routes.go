package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IRouter interface {
	Setup()
}

type AppRoute struct {
	app *gin.Engine
	db  *gorm.DB
}
type Route struct {
	router *gin.RouterGroup
	db     *gorm.DB
	path   string
}

// setup implements IRoute
func (appRoute *AppRoute) Setup() {
	router := appRoute.app.Group("api/v1")
	authRoutes(router)
	userRoutes(router)
	routerNoteGroup := router.Group("note")
	newNoteRoutes(routerNoteGroup, appRoute.db, "").setupRoutesNote()
	newTaskRoutes(routerNoteGroup, appRoute.db, "/task").setupRoutesTask()
	newLabelRoutes(routerNoteGroup, appRoute.db, "/label").setupRoutesLabel()
}

func NewApplicationRouter(app *gin.Engine, db *gorm.DB) IRouter {
	return &AppRoute{
		app: app,
		db:  db,
	}
}
