package models

type Role struct {
	ID          string       `gorm:"primarykey"`
	Name        string       `gorm:"unique"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}
