package models

import (
	"time"

	"github.com/yornifpaz/back_noteapp/lib"
	"gorm.io/gorm"
)

type Label struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string         `json:"title" binding:"required"`
	UserID    string         `json:"user_id" binding:"required"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (label *Label) BeforeCreate(tx *gorm.DB) (err error) {
	label.ID = lib.NewIdLibrary().Create()
	return
}
