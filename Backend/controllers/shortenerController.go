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
	const maxAllowedURLS = 10

	user, _ := c.Get("user")

	u := user.(models.User)

	var mappings []models.Url_mappings
	databaseQuery := initializers.DB.Where("user_id = ? AND Enabled = true", u.ID).Find(&mappings)

	fmt.Println(databaseQuery.RowsAffected)

	if databaseQuery.RowsAffected >= maxAllowedURLS {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "You have reached your maximum URLS! Delete some to get more.",
		})
		return
	}

	var req struct {
		URL        string
		ExpiryDate time.Time
	}

	if c.BindJSON(&req) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to bind URL to body.",
		})
		return
	}

	fmt.Println(req.URL)

	newURL := generateShortenedURL(8)

	if req.ExpiryDate.Before(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)) {
		req.ExpiryDate = time.Now().Add(time.Hour * 999999)
	}

	ShortenedURL := models.Url_mappings{UserID: u.ID, ShortCode: newURL, FullURL: req.URL, CreatedAt: time.Now(), Enabled: true, ExpiryDate: req.ExpiryDate}

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

	if url_mapping.ExpiryDate.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "URL has expired.",
		})
		return
	}

	if url_mapping.Enabled {
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

func GetAllMyURLS(c *gin.Context) {
	user, _ := c.Get("user")

	u := user.(models.User)

	var mappings []models.Url_mappings
	err := initializers.DB.Where("Enabled = ? AND user_id = ?", true, u.ID).Find(&mappings)
	fmt.Println(err)

	c.JSON(http.StatusOK, gin.H{
		"Code": mappings,
	})

}
