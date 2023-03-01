package labelrepository

import (
	"github.com/yornifpaz/back_noteapp/models"
	"gorm.io/gorm"
)

type ILabelRepository interface {
	GetById(id string) (label models.Label, err error)
	GetByTitle(title string, userId string) (label models.Label, err error)
	GetAll(userId string) (labels []models.Label, err error)
	Save(label models.Label) (err error)
	Update(updateLabel models.Label, label models.Label) (err error)
	Delete(id string) (err error)
}

type LabelRepository struct {
	DB *gorm.DB
}

// Delete implements ILabelRepository
func (r *LabelRepository) Delete(id string) (err error) {
	err = r.DB.Delete(&models.Label{}, "id=?", id).Error
	return
}

// GetAll implements ILabelRepository
func (r *LabelRepository) GetAll(userId string) (labels []models.Label, err error) {
	err = r.DB.Where("user_id = ?", userId).Find(&labels).Error
	return
}

// GetById implements ILabelRepository
func (r *LabelRepository) GetById(id string) (label models.Label, err error) {
	err = r.DB.First(&label, "id=?", id).Error
	return
}

// GetByTitle implements ILabelRepository
func (r *LabelRepository) GetByTitle(title string, userId string) (label models.Label, err error) {
	err = r.DB.Where("user_id=?", userId).Where("title=?", title).First(&label).Error
	return
}

// Save implements ILabelRepository
func (r *LabelRepository) Save(label models.Label) (err error) {
	err = r.DB.Create(&label).Error
	return
}

// Update implements ILabelRepository
func (r *LabelRepository) Update(updateLabel models.Label, label models.Label) (err error) {
	err = r.DB.Model(&label).Updates(updateLabel).Error
	return
}

func NewLabelRepository(db *gorm.DB) ILabelRepository {
	return &LabelRepository{
		DB: db,
	}
}
