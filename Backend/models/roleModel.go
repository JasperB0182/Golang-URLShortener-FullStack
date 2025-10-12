package models

type UserRole struct {
	ID   uint   `gorm:"primaryKey"`
	Role string `gorm:"unique"`
}
