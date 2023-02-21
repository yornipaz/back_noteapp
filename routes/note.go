package routes

import (
	"github.com/gin-gonic/gin"
	notecontroller "github.com/yornifpaz/back_noteapp/controllers/noteController"
	"github.com/yornifpaz/back_noteapp/middleware"
)

func noteRoutes(app *gin.RouterGroup) {
	noteGroup := app.Group("/notes")
	noteGroup.Use(middleware.Authenticate)
	noteGroup.GET("", notecontroller.GetAll)
	noteGroup.POST("", notecontroller.Create)
	noteGroup.PATCH("", notecontroller.Update)
	noteGroup.PATCH("/archived", notecontroller.UpdateArchived)
	noteGroup.DELETE("/:id", notecontroller.Delete)
	noteGroup.POST("/tasks", notecontroller.AddTasks)
	noteGroup.PATCH("/tasks", notecontroller.UpdateTask)
	noteGroup.DELETE("/tasks/:id", notecontroller.DeleteTask)

}
