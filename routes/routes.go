package routes

import (
	"github.com/gin-gonic/gin"
)

func AppRoutes(app *gin.Engine) {
	v1 := app.Group("api/v1")
	authRoutes(v1)
	userRoutes(v1)
	noteRoutes(v1)
	labelRoutes(v1)

}
