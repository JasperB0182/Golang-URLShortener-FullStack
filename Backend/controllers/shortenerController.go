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

	ShortenedURL := models.Url_mappings{ShortCode: newURL, FullURL: req.URL, CreatedAt: time.Now()}

	result := initializers.DB.Create(&ShortenedURL)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "FAILED TO SHORTEN URL!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Success": "Succesfully shortened URL!",
	})
}
