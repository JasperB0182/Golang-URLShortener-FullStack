package controllers

import (
	"fmt"
	"keceox_modules/initializers"
	"keceox_modules/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func AdminSeeAllActiveURLS(c *gin.Context) {
	var mappings []models.Url_mappings
	err := initializers.DB.Preload("User").Preload("User.UserRole").Where("Enabled = true AND expiry_date > ?", time.Now()).Find(&mappings)

	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Couldn't retrieve codes",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Code": mappings,
	})
}

func AdminSeeAllUsers(c *gin.Context) {
	var users []models.User
	err := initializers.DB.Preload("UserRole").Find(&users)

	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Couldn't retrieve users",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Users": users,
	})
}

func AdminDeleteAccount(c *gin.Context) {
	id := c.Param("id")

	admin, _ := c.Get("user")

	a := admin.(models.User)

	var u models.User
	initializers.DB.First(&u, "id = ?", id)

	if u.RoleID != 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Couldn't delete admin",
		})
		return
	}

	result := initializers.DB.Where("user_id = ?", u.ID).Delete(&models.Url_mappings{})

	if result.Error != nil {
		fmt.Println(result.RowsAffected)

		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Couldn't delete user",
		})
		return

	}

	initializers.DB.Delete(&u, u.ID)
	// DELETE FROM users WHERE id = 10;

	LogFile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("ERROR WRITING TO LOGS!")
	}

	LogFile.WriteString("[" + time.Now().Format("Jan 2, 2006 3:04pm") + "] " + "Admin " + a.Email + " has deleted account with ID: " + id + "\n")

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted account",
	})

}

func AdminDeleteURL(c *gin.Context) {
	id := c.Param("id")
	user, _ := c.Get("user")

	u := user.(models.User)

	var url_mapping models.Url_mappings
	initializers.DB.First(&url_mapping, "Short_Code = ?", id)

	if url_mapping.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Doesn't exist or is already deleted.",
		})
		return
	}

	if !url_mapping.Enabled {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Doesn't exist or is already deleted.",
		})
		return
	}

	initializers.DB.Delete(url_mapping)

	LogFile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("ERROR WRITING TO LOGS!")
	}

	LogFile.WriteString("[" + time.Now().Format("Jan 2, 2006 3:04pm") + "] " + "Admin " + u.Email + " has disabled URL with shortcode: " + id + "\n")

	c.JSON(http.StatusOK, gin.H{
		"Message": "Succesfully deleted the URL permanently!",
	})
}

func AdminDisableMultipleURL(c *gin.Context) {
	var body struct {
		Codes []string `json:"codes"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, _ := c.Get("user")
	u := user.(models.User)

	LogFile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("ERROR OPENING LOG FILE!")
	}
	defer LogFile.Close()

	var succeeded []string
	var failed []string

	for _, code := range body.Codes {
		var url_mapping models.Url_mappings
		initializers.DB.First(&url_mapping, "Short_Code = ?", code)

		if url_mapping.ID == 0 || !url_mapping.Enabled {
			failed = append(failed, code)
			continue
		}

		url_mapping.Enabled = false
		if err := initializers.DB.Save(&url_mapping).Error; err != nil {
			failed = append(failed, code)
			continue
		}

		succeeded = append(succeeded, code)

		if LogFile != nil {
			LogFile.WriteString("[" + time.Now().Format("Jan 2, 2006 3:04pm") + "] " + "Admin " + u.Email + " has disabled URL with shortcode: " + url_mapping.ShortCode + "\n")
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"succeeded": succeeded,
		"failed":    failed,
	})
}
