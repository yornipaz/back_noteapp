package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yornifpaz/back_noteapp/app/http/middleware"
	"github.com/yornifpaz/back_noteapp/config"
	"github.com/yornifpaz/back_noteapp/routes"
)

func init() {
	config.Init()
}

func main() {
	app := gin.Default()
	app.Use(middleware.CORSMiddleware())
	routes.NewApplicationRouter(app, config.DB).Setup()
	app.Run()
}
