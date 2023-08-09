package models

import (
	"github.com/yornifpaz/back_noteapp/lib"
	"gorm.io/gorm"
)

type Role struct {
	ID          string `gorm:"primarykey"`
	Name        string `gorm:"unique"`
	Permissions []Permission
}

func (role *Role) BeforeCreate(tx *gorm.DB) (err error) {
	role.ID = lib.NewIdLibrary().Create()
	return
}
