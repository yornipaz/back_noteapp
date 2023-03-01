package noterepository

import (
	"github.com/yornifpaz/back_noteapp/models"
	"gorm.io/gorm"
)

type INoteRepository interface {
	GetById(id string) (note models.Note, err error)
	GetAll(userId string) (notes []models.Note, err error)
	Save(note models.Note) (err error)
	Update(note models.Note, updateNote models.Note) (noteUpdate models.Note, err error)
	UpdateArchived(note models.Note, updateNote models.Note) (noteUpdate models.Note, err error)
	Delete(id string) (err error)
}

type NoteRepository struct {
	db *gorm.DB
}

// UpdateArchived implements INoteRepository
func (r *NoteRepository) UpdateArchived(note models.Note, updateNote models.Note) (noteUpdate models.Note, err error) {
	err = r.db.Model(&note).Updates(map[string]interface{}{"updated_at": updateNote.UpdatedAt, "archived": updateNote.Archived}).Error
	return note, err
}

// Delete implements INoteRepository
func (r *NoteRepository) Delete(id string) (err error) {
	err = r.db.Delete(&models.Note{}, "id=?", id).Error
	return
}

// GetAll implements INoteRepository
func (r *NoteRepository) GetAll(userId string) (notes []models.Note, err error) {
	err = r.db.Where("user_id = ?", userId).Preload("Tasks").Find(&notes).Error
	return
}

// GetById implements INoteRepository
func (r *NoteRepository) GetById(id string) (note models.Note, err error) {
	err = r.db.First(&note, "id=?", id).Error
	return
}

// Save implements INoteRepository
func (r *NoteRepository) Save(note models.Note) (err error) {
	err = r.db.Create(&note).Error
	return
}

// Update implements INoteRepository
func (r *NoteRepository) Update(note models.Note, updateNote models.Note) (noteUpdate models.Note, err error) {
	err = r.db.Model(&note).Updates(updateNote).Error
	return note, err
}

func NewNoteRepository(db *gorm.DB) INoteRepository {
	return &NoteRepository{db: db}
}
