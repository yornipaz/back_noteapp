package models

import (
	"time"

	"github.com/lib/pq"
	"github.com/yornifpaz/back_noteapp/lib"
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
	UserID      string
	Labels      pq.StringArray `gorm:"type:text[]"`
	Tasks       []Task         `gorm:"foreignKey:NoteRefer"`
}

func (note *Note) BeforeCreate(tx *gorm.DB) (err error) {
	note.ID = lib.NewIdLibrary().Create()
	return
}
