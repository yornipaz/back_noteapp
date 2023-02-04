package config

import "github.com/yornifpaz/back_noteapp/models"

func syncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Note{})
}
