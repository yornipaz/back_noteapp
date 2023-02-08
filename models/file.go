package models

import "mime/multipart"

type File struct {
	File multipart.File `json:"avatar,omitempty" validate:"required"`
}
