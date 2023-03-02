package userrepository

import (
	"github.com/yornifpaz/back_noteapp/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetById(id string) (user models.User, err error)
	GetByEmail(email string) (user models.User, err error)
	Save(user models.User) (err error)
	Update(user models.User, updateUser models.User) (userUpdate models.User, err error)
}

type UserRepository struct {
	db *gorm.DB
}

// GetByEmail implements IUserRepository
func (r *UserRepository) GetByEmail(email string) (user models.User, err error) {
	err = r.db.First(&user, "email=?", email).Error
	return
}

// GetById implements IUserRepository
func (r *UserRepository) GetById(id string) (user models.User, err error) {
	err = r.db.First(&user, "id=?", id).Error
	return
}

// Save implements IUserRepository
func (r *UserRepository) Save(user models.User) (err error) {
	err = r.db.Create(&user).Error
	return
}

// Update implements IUserRepository
func (r *UserRepository) Update(user models.User, updateUser models.User) (userUpdate models.User, err error) {
	err = r.db.Model(&user).Updates(updateUser).Error
	return user, err
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}
