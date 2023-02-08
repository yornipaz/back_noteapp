package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Task struct {
	Id        string
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
	NoteRefer string
}
type Note struct {
	ID          string `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Reminder    time.Time      `json:"reminder"`
	Archived    bool           `json:"archived"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Title       string         `json:"title"  `
	Description string         `json:"description"`
	UserID      string         `json:"user_id" binding:"required"`
	Labels      pq.StringArray `gorm:"type:text[]"`
	Tasks       []Task         `gorm:"foreignKey:NoteRefer"`
}

func (note *Note) BeforeCreate(tx *gorm.DB) (err error) {
	note.ID = uuid.NewString()
	return
}
