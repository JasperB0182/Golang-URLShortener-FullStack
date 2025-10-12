package seeders

import (
	"keceox_modules/initializers"
	"keceox_modules/models"

	"gorm.io/gorm"
)

func RoleSeeder() {
	roles := []models.UserRole{
		{Role: "User"},
		{Role: "Admin"},
	}

	for _, UserRole := range roles {
		var existing models.UserRole
		if err := initializers.DB.Where("Role = ?", UserRole.Role).First(&existing).Error; err == gorm.ErrRecordNotFound {
			initializers.DB.Create(&UserRole)
		}
	}
}
