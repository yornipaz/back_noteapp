package config

import "github.com/yornifpaz/back_noteapp/app/models"

var initialModels = []interface{}{
	&models.User{},
	&models.Note{},
	&models.Label{},
	&models.Task{},
	&models.Permission{},
	&models.Role{},
}
