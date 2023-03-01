package notefactory

import (
	"time"

	"github.com/lib/pq"
	"github.com/yornifpaz/back_noteapp/models"
)

type INoteFactory interface {
	Create(title string, userId string, content string) (note models.Note)
	Update(title string, userId string, content string, labels pq.StringArray, tasks []models.Task, reminder time.Time) (note models.Note)
	UpdateTask(tasks []models.Task, noteId string) (tasksUpdate []models.Task)
	UpdateArchived(archived bool) (note models.Note)
}

type NoteFactory struct{}

// UpdateArchived implements INoteFactory
func (*NoteFactory) UpdateArchived(archived bool) (note models.Note) {
	return models.Note{
		Archived:  archived,
		UpdatedAt: time.Now(),
	}
}

// UpdateTask implements INoteFactory
func (*NoteFactory) UpdateTask(tasks []models.Task, noteId string) (tasksUpdate []models.Task) {

	if len(tasks) > 0 {
		for _, task := range tasks {
			var updateTask = models.Task{
				Id:        task.Id,
				Completed: task.Completed,
				Content:   task.Content,
				NoteRefer: noteId,
			}
			tasksUpdate = append(tasksUpdate, updateTask)
		}

	}
	return
}

// Create implements INoteFactory
func (*NoteFactory) Create(title string, userId string, content string) (note models.Note) {
	note = models.Note{Title: title,
		Description: content,
		UserID:      userId,
		Archived:    false,
		Reminder:    time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now()}
	return
}

// Update implements INoteFactory
func (*NoteFactory) Update(title string, userId string, content string, labels pq.StringArray, tasks []models.Task, reminder time.Time) (note models.Note) {
	note = models.Note{UpdatedAt: time.Now(),
		Reminder:    reminder,
		Title:       title,
		Description: content,
		UserID:      userId,
		Labels:      labels,
		Tasks:       tasks}
	return
}

func NewNoteFactory() INoteFactory {
	return &NoteFactory{}
}
