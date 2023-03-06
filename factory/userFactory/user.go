package userfactory

import (
	"time"

	"github.com/yornifpaz/back_noteapp/models"
)

type IUserFactory interface {
	Create(firstName string, lastName string, email string, password string, avatar string) (user models.User)
	Update(firstName string, lastName string, email string) (user models.User)
}

type UserFactory struct{}

// Update implements IUserFactory
func (*UserFactory) Update(firstName string, lastName string, email string) (user models.User) {
	user = models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		UpdatedAt: time.Now(),
	}
	return
}

// Create implements IUserFactory
func (*UserFactory) Create(firstName string, lastName string, email string, password string, avatar string) (user models.User) {
	user = models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		Avatar:    avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		LogoutAt:  time.Now(),
		Status:    "Created"}
	return
}

func NewUserFactory() IUserFactory {
	return &UserFactory{}
}
