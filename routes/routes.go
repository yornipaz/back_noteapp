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
	appRoute.app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Api work in progress",
		})
	})
	router := appRoute.app.Group("api/v1")
	routerNoteGroup := router.Group("notes")
	newAuthRoutes(router, appRoute.db, "/auth").setupRoutesAuth()
	newUserRoutes(router, appRoute.db, "/user").setupRoutesUser()
	newNoteRoutes(routerNoteGroup, appRoute.db, "").setupRoutesNote()
	newTaskRoutes(routerNoteGroup, appRoute.db, "/tasks").setupRoutesTask()
	newLabelRoutes(routerNoteGroup, appRoute.db, "/label").setupRoutesLabel()
}

func NewApplicationRouter(app *gin.Engine, db *gorm.DB) IRouter {
	return &AppRoute{
		app: app,
		db:  db,
	}
}
