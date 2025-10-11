package controllers

import (
	"fmt"
	"keceox_modules/initializers"
	"keceox_modules/models"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func generateShortenedURL(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	var result []byte

	for i := 0; i < length; i++ {
		index := seededRand.Intn(len(charset))
		result = append(result, charset[index])
	}

	return string(result)
}

func ShortenURL(c *gin.Context) {
	user, _ := c.Get("user")

	u := user.(models.User)

	var req struct {
		URL string
	}

	if c.BindJSON(&req) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to bind URL to body.",
		})
		return
	}

	fmt.Println(req.URL)

	newURL := generateShortenedURL(8)

	ShortenedURL := models.Url_mappings{UserID: u.ID, ShortCode: newURL, FullURL: req.URL, CreatedAt: time.Now(), Enabled: true}

	result := initializers.DB.Create(&ShortenedURL)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "FAILED TO SHORTEN URL!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Success": "Succesfully shortened URL!",
		"Code":    newURL,
	})
}

func GetOriginalURL(c *gin.Context) {
	id := c.Param("id")

	fmt.Println(id)

	var url_mapping models.Url_mappings
	initializers.DB.First(&url_mapping, "Short_Code = ?", id)

	if url_mapping.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "URL is currently disabled or doesn't exist.",
		})
		return
	}

	if url_mapping.Enabled == true {
		c.JSON(http.StatusOK, gin.H{
			"URL": url_mapping.FullURL,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "URL is currently disabled or doesn't exist.",
		})
		return
	}

}

func DisableURL(c *gin.Context) {
	id := c.Param("id")

	user, _ := c.Get("user")

	u := user.(models.User)

	var url_mapping models.Url_mappings
	initializers.DB.First(&url_mapping, "Short_Code = ?", id)

	if url_mapping.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "URL is not yours, doesn't exist or is already disabled.",
		})
		return
	}

	if u.ID == url_mapping.UserID {
		if !url_mapping.Enabled {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "URL is not yours, doesn't exist or is already disabled.",
			})
			return
		}

		url_mapping.Enabled = false
		initializers.DB.Save(url_mapping)
		c.JSON(http.StatusOK, gin.H{
			"Message": "Succesfully disabled the URL permanently!",
		})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "URL is not yours, doesn't exist or is already disabled.",
		})
		return
	}

}
