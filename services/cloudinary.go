package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/yornifpaz/back_noteapp/helpers"
	"github.com/yornifpaz/back_noteapp/models"
)

var (
	validate = validator.New()
)

type mediaUpload interface {
	FileUpload(file models.File) (string, error)
}

type media struct{}

func NewMediaUpload() mediaUpload {
	return &media{}
}

func (s *media) FileUpload(file models.File) (string, error) {
	//validate
	err := validate.Struct(file)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, err := helpers.ImageUpload(file.File)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}
