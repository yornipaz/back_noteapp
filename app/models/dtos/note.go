package dtos

import (
	"time"

	"github.com/lib/pq"
	"github.com/yornifpaz/back_noteapp/app/models"
)

type AddNote struct {
	Archived bool
	Title    string
	Content  string
	Tasks    []models.Task
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
	Id       string
}
type AddTask struct {
	NoteId string
	Task   string
}

type UpdateTask struct {
	Id        string
	Content   string
	Completed bool
	NoteRefer string
}
