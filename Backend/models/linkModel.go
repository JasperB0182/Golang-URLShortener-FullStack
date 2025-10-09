package models

import "time"

type Url_mappings struct {
	ID        uint   `gorm:"primaryKey"`
	ShortCode string `gorm:"uniqueIndex"`
	FullURL   string
	CreatedAt time.Time
}
