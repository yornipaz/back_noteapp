package models

type Permission struct {
	ID   string `gorm:"primarykey"`
	Name string `gorm:"unique"`
}
