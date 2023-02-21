package labelrepository

import (
	"time"

	"github.com/yornifpaz/back_noteapp/models"
	"gorm.io/gorm"
)

type ILabelRepository interface {
	GetById(id string) (label models.Label, err error)
	GetByTitle(title string) (label models.Label, err error)
	GetAll() (labels []models.Label, err error)
	Save(title string) (err error)
	Update(title string, label models.Label) (err error)
	Delete(id string) (err error)
}

type labelRepository struct {
	DB     *gorm.DB
	UserId string
}

// Delete implements ILabelRepository
func (r *labelRepository) Delete(id string) (err error) {
	err = r.DB.Delete(&models.Label{}, "id=?", id).Error
	return
}

// GetAll implements ILabelRepository
func (r *labelRepository) GetAll() (labels []models.Label, err error) {
	err = r.DB.Where("user_id = ?", r.UserId).Find(&labels).Error
	return
}

// GetById implements ILabelRepository
func (r *labelRepository) GetById(id string) (label models.Label, err error) {
	err = r.DB.First(&label, "id=?", id).Error

	return
}

// GetByTitle implements ILabelRepository
func (r *labelRepository) GetByTitle(title string) (label models.Label, err error) {
	err = r.DB.Where("user_id=?", r.UserId).Where("title=?", title).First(&label).Error
	return
}

// Save implements ILabelRepository
func (r *labelRepository) Save(title string) (err error) {
	label := models.Label{Title: title, UserID: r.UserId, UpdatedAt: time.Now(), CreatedAt: time.Now()}
	err = r.DB.Create(&label).Error
	return
}

// Update implements ILabelRepository
func (r *labelRepository) Update(title string, label models.Label) (err error) {
	updateLabel := models.Label{Title: title, UpdatedAt: time.Now()}

	err = r.DB.Model(&label).Updates(updateLabel).Error
	return
}

func NewLabelRepository(db *gorm.DB, userId string) ILabelRepository {
	return &labelRepository{
		DB:     db,
		UserId: userId,
	}
}
