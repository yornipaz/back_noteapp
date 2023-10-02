package config

import "github.com/yornifpaz/back_noteapp/app/models"

func syncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Note{}, &models.Label{}, &models.Task{}, &models.Permission{}, &models.Role{})

}
