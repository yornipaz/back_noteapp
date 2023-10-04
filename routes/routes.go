package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yornifpaz/back_noteapp/app/models"
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
		c.JSON(http.StatusOK, models.APIResponse{
			Message: "Api work in progress",
			Status:  http.StatusOK,
			Data:    nil,
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
