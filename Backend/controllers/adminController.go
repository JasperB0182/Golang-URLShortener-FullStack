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
