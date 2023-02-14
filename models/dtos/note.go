package dtos

import (
	"time"

	"github.com/lib/pq"
	"github.com/yornifpaz/back_noteapp/models"
)

type AddNote struct {
	Reminder time.Time
	Archived bool
	Title    string
	Content  string
	Labels   pq.StringArray
	Tasks    []models.Task
	image    string
}
type UpdateNote struct {
	Reminder time.Time
	Archived bool
	Title    string
	Content  string
	Labels   pq.StringArray
	Tasks    []models.Task
	image    string
	UserID   string
}
