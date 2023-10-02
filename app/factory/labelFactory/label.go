package labelfactory

import (
	"time"

	"github.com/yornifpaz/back_noteapp/app/models"
)

type ILabelFactory interface {
	Create(title string, userId string) (label models.Label)
	Update(title string) (label models.Label)
}

type LabelFactory struct{}

// Update implements ILabelFactory
func (*LabelFactory) Update(title string) (label models.Label) {
	label = models.Label{Title: title, UpdatedAt: time.Now()}
	return
}

// Create implements ILabelFactory
func (*LabelFactory) Create(title string, userId string) (label models.Label) {
	label = models.Label{Title: title, UserID: userId, UpdatedAt: time.Now(), CreatedAt: time.Now()}
	return
}

func NewLabelFactory() ILabelFactory {
	return &LabelFactory{}
}
