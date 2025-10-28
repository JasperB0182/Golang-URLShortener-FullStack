package controllers

import (
	"fmt"
	"keceox_modules/initializers"
	"keceox_modules/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddToCredit(c *gin.Context) {
	var body struct {
		AddedCredit float64 `json:"addedCredit"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	user, _ := c.Get("user")
	u := user.(models.User)

	u.Credit += body.AddedCredit

	fmt.Printf("New credit: %.2f\n", u.Credit)

	if err := initializers.DB.Model(&u).Update("credit", u.Credit).Error; err != nil {
		c.JSON(500, gin.H{"error": "Database update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "Successfully added credit to account!",
		"credit":  u.Credit,
	})
}
