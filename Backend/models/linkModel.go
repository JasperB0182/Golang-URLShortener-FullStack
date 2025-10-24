package models

import "time"

type Url_mappings struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	User       User   `gorm:"foreignKey:UserID"`
	ShortCode  string `gorm:"uniqueIndex"`
	FullURL    string
	CreatedAt  time.Time
	Enabled    bool
	ExpiryDate time.Time
	UsageCount float64
}
