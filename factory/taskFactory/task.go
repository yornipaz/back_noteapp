package taskfactory

import (
	"github.com/google/uuid"
	"github.com/yornifpaz/back_noteapp/models"
)

type ITaskFactory interface {
	Create(content string, noteId string) (task models.Task)
}

type TaskFactory struct{}

// Create implements ITaskFactory
func (*TaskFactory) Create(content string, noteId string) (task models.Task) {
	task = models.Task{
		Completed: false,
		Content:   content,
		Id:        uuid.NewString(),
		NoteRefer: noteId,
	}
	return
}

func NewtaskFactory() ITaskFactory {
	return &TaskFactory{}
}
