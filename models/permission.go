package models

import (
	"github.com/yornifpaz/back_noteapp/lib"
	"gorm.io/gorm"
)

type Permission struct {
	ID   string `gorm:"primarykey"`
	Name string `gorm:"unique"`
}

func (permission *Permission) BeforeCreate(tx *gorm.DB) (err error) {
	permission.ID = lib.NewIdLibrary().Create()
	return
}
