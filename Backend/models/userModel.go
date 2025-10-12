package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	Name     string
	RoleID   uint
	UserRole UserRole `gorm:"foreignKey:RoleID;references:ID"`
}
