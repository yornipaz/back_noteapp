package config

import "github.com/yornifpaz/back_noteapp/models"

func syncDatabase() {
	DB.AutoMigrate(&models.Note{}, &models.User{}, &models.Label{}, &models.Task{})

}
