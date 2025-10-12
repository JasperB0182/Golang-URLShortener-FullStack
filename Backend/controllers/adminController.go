package controllers

import (
	"keceox_modules/initializers"
	"keceox_modules/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminSeeAllActiveURLS(c *gin.Context) {
	var mappings []models.Url_mappings
	err := initializers.DB.Where("Enabled = true").Find(&mappings)
	println(err)

	c.JSON(http.StatusOK, gin.H{
		"Code": mappings,
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
