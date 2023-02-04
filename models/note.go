package models

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title       string `json:"title" binding:"required" `
	Description string `json:"description" binding:"required"`
	UserID      uint   `json:"user_id" binding:"required"`
	Tags        string `json:"tags" `
}
