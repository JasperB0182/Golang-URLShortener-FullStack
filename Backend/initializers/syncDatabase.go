package initializers

import "keceox_modules/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
