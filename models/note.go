package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Note struct {
	ID          string `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Title       string         `json:"title"  `
	Description string         `json:"description"`
	UserID      string         `json:"user_id" binding:"required"`
	Tags        pq.StringArray `gorm:"type:text[]"`
}

func (note *Note) BeforeCreate(tx *gorm.DB) (err error) {
	note.ID = uuid.NewString()
	return
}
