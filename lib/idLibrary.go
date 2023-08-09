package lib

import (
	"github.com/google/uuid"
)

type IIdLibrary interface {
	Create() (id string)
}
type IdLibrary struct {
}

// Create implements IIdLibrary.
func (*IdLibrary) Create() (id string) {
	return uuid.NewString()
}

func NewIdLibrary() IIdLibrary {
	return &IdLibrary{}
}
