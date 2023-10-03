package models

import (
	"time"

	"github.com/yornifpaz/back_noteapp/lib"
	"gorm.io/gorm"
)

// ? ForgotPasswordInput struct
type ForgotPasswordInput struct {
	Email string `json:"email" binding:"required"`
}

// ? ResetPasswordInput struct
type ResetPasswordInput struct {
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}
type UserStatus string

const (
	Active   UserStatus = "active"
	Inactive UserStatus = "inactive"
	// Otros estados según sea necesario
)

type User struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	LogoutAt  time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Status    UserStatus
	FirstName string
	LastName  string
	Email     string `gorm:"unique" `
	Avatar    string
	IsActive  bool `gorm:"default:true"`
	Password  string
	Verified  bool   `gorm:"default:false"`
	Roles     []Role `gorm:"many2many:user_roles;"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = lib.NewIdLibrary().Create()
	return
}
