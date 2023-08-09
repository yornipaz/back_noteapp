package taskfactory

import (
	"github.com/yornifpaz/back_noteapp/lib"
	"github.com/yornifpaz/back_noteapp/models"
)

type ITaskFactory interface {
	Create(content string, noteId string) (task models.Task)
}

type TaskFactory struct {
	id lib.IIdLibrary
}

// Create implements ITaskFactory
func (f *TaskFactory) Create(content string, noteId string) (task models.Task) {
	id := f.id.Create()
	task = models.Task{
		Completed: false,
		Content:   content,
		Id:        id,
		NoteRefer: noteId,
	}
	return
}

func NewTaskFactory(idLibrary lib.IIdLibrary) ITaskFactory {
	return &TaskFactory{idLibrary}
}
