package taskrepository

import (
	"github.com/yornifpaz/back_noteapp/app/models"
	"gorm.io/gorm"
)

type ITaskRepository interface {
	GetById(id string) (task models.Task, err error)
	Save(task models.Task) (err error)
	Update(task models.Task, content string, completed bool) (err error)
	Delete(id string) (err error)
}

type TaskRepository struct {
	db *gorm.DB
}

// Delete implements ITaskRepository
func (r *TaskRepository) Delete(id string) (err error) {
	err = r.db.Delete(&models.Task{}, "id=?", id).Error
	return
}

// GetById implements ITaskRepository
func (r *TaskRepository) GetById(id string) (task models.Task, err error) {
	err = r.db.First(&task, "id=?", id).Error
	return
}

// Save implements ITaskRepository
func (r *TaskRepository) Save(task models.Task) (err error) {
	err = r.db.Save(&task).Error
	return
}

// Update implements ITaskRepository
func (r *TaskRepository) Update(task models.Task, content string, completed bool) (err error) {
	err = r.db.Model(&task).Updates(map[string]interface{}{"completed": completed, "content": content}).Error
	return
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &TaskRepository{db: db}
}
