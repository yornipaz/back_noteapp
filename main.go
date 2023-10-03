package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yornifpaz/back_noteapp/app/http/middleware"
	"github.com/yornifpaz/back_noteapp/config"
	"github.com/yornifpaz/back_noteapp/routes"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {

	newConfigurationApplication := config.NewConfigurationApplication()
	dbApplication, errors := newConfigurationApplication.Init()
	if len(errors) > 0 {
		for _, err := range errors {
			fmt.Println(err.Error())
		}

	}
	db = dbApplication

}

func main() {
	app := gin.Default()
	app.Use(middleware.CORSMiddleware())
	routes.NewApplicationRouter(app, db).Setup()
	app.Run()
}
