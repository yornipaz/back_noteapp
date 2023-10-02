package models

import (
	"time"

	"github.com/yornifpaz/back_noteapp/lib"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	LogoutAt  time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Status    string
	FirstName string
	LastName  string
	Email     string `gorm:"unique" `
	Avatar    string
	IsActive  bool `gorm:"default:true"`
	Password  string
	Roles     []Role `gorm:"many2many:user_roles;"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = lib.NewIdLibrary().Create()
	return
}
