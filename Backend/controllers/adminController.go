package controllers

import (
	"keceox_modules/initializers"
	"keceox_modules/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AdminSeeAllActiveURLS(c *gin.Context) {
	var mappings []models.Url_mappings
	err := initializers.DB.Preload("User").Preload("User.UserRole").Where("Enabled = true AND expiry_date > ?", time.Now()).Find(&mappings)
	println(err)

	c.JSON(http.StatusOK, gin.H{
		"Code": mappings,
	})
}

func AdminSeeAllUsers(c *gin.Context) {
	var users []models.User
	err := initializers.DB.Preload("UserRole").Find(&users)
	println(err)

	c.JSON(http.StatusOK, gin.H{
		"Users": users,
	})
}

func AdminDeleteAccount(c *gin.Context) {
	id := c.Param("id")

	var u models.User
	initializers.DB.First(&u, "id = ?", id)

	initializers.DB.Delete(&u, u.ID)
	// DELETE FROM users WHERE id = 10;

	c.JSON(http.StatusOK, gin.H{
		"message": "get rekt lol",
	})

}

func AdminDisableURL(c *gin.Context) {
	id := c.Param("id")

	var url_mapping models.Url_mappings
	initializers.DB.First(&url_mapping, "Short_Code = ?", id)

	if url_mapping.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Doesn't exist or is already disabled.",
		})
		return
	}

	if !url_mapping.Enabled {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Doesn't exist or is already disabled.",
		})
		return
	}

	url_mapping.Enabled = false
	initializers.DB.Save(url_mapping)
	c.JSON(http.StatusOK, gin.H{
		"Message": "Succesfully disabled the URL permanently!",
	})
}
