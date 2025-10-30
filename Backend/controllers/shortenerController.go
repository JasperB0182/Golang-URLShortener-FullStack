package controllers

import (
	"fmt"
	"keceox_modules/initializers"
	"keceox_modules/models"
	"math/rand"
	"net/http"
	"net/url"
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
	databaseQuery := initializers.DB.Where(
		"user_id = ? AND Enabled = ? AND expiry_date >= ?",
		u.ID, true, time.Now(),
	).Find(&mappings)

	var req struct {
		URL         string    `json:"URL"`
		ExpiryDate  time.Time `json:"ExpiryDate"`
		UsedCredits bool      `json:"usedCredits"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind URL to body."})
		return
	}

	if databaseQuery.RowsAffected >= maxAllowedURLS && u.Credit < 10 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "You have reached your maximum URLs! Delete some to get more, or buy some credit!",
		})
		return
	}

	check, err := url.ParseRequestURI(req.URL)
	if err != nil || check == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entered text is not a valid URL."})
		return
	}

	newURL := generateShortenedURL(8)

	if req.UsedCredits {
		if u.Credit < 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough credits!"})
			return
		}
	} else {
		if u.RoleID == 1 {
			maxExpiry := time.Now().AddDate(0, 1, 0)
			if req.ExpiryDate.After(maxExpiry) || req.ExpiryDate.IsZero() {
				req.ExpiryDate = maxExpiry
			}
		}
	}

	if req.ExpiryDate.Before(time.Now()) {
		if u.RoleID == 1 {
			req.ExpiryDate = time.Now().AddDate(0, 1, 0)
		} else {
			req.ExpiryDate = time.Now().Add(time.Hour * 999999)
		}
	}

	ShortenedURL := models.Url_mappings{
		UserID:     u.ID,
		ShortCode:  newURL,
		FullURL:    req.URL,
		CreatedAt:  time.Now(),
		Enabled:    true,
		ExpiryDate: req.ExpiryDate,
		UsageCount: 0,
	}

	overMaxURLs := databaseQuery.RowsAffected >= maxAllowedURLS

	if u.RoleID != 2 {
		var creditsNeeded float64 = 0

		if overMaxURLs {
			creditsNeeded += 10
		}

		if req.UsedCredits {
			creditsNeeded += 10
		}

		if u.Credit < creditsNeeded {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough credits!"})
			return
		}

		u.Credit -= creditsNeeded
	}

	if err := initializers.DB.Create(&ShortenedURL).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "FAILED TO SHORTEN URL!"})
		return
	}

	if err := initializers.DB.Model(&u).Update("credit", u.Credit).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update credit"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Success": "Successfully shortened URL!",
		"Code":    newURL,
	})
}

func GetOriginalURL(c *gin.Context) {
	id := c.Param("id")

	fmt.Println(id)

	var url_mapping models.Url_mappings
	initializers.DB.First(&url_mapping, "Short_Code = ?", id)

	url_mapping.UsageCount += 1

	initializers.DB.Save(url_mapping)

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

func EnableURL(c *gin.Context) {
	id := c.Param("id")

	user, _ := c.Get("user")

	u := user.(models.User)

	var mappings []models.Url_mappings
	databaseQuery := initializers.DB.Where(
		"user_id = ? AND Enabled = ? AND expiry_date >= ?",
		u.ID, true, time.Now(),
	).Find(&mappings)

	if databaseQuery.RowsAffected >= 10 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "You have reached your max amount of links!",
		})
		return
	}

	var url_mapping models.Url_mappings
	initializers.DB.First(&url_mapping, "Short_Code = ?", id)

	if url_mapping.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "URL is not yours, doesn't exist or is already enabled.",
		})
		return
	}

	if u.ID == url_mapping.UserID {
		if url_mapping.Enabled {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "URL is not yours, doesn't exist or is already enabled.",
			})
			return
		}

		url_mapping.Enabled = true
		initializers.DB.Save(url_mapping)
		c.JSON(http.StatusOK, gin.H{
			"Message": "Succesfully enabled the URL!",
		})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "URL is not yours, doesn't exist or is already enabled.",
		})
		return
	}

}

func GetAllMyURLS(c *gin.Context) {
	user, _ := c.Get("user")

	u := user.(models.User)

	var checkmappings []models.Url_mappings
	databaseQuery := initializers.DB.Where(
		"user_id = ? AND Enabled = ? AND expiry_date >= ?",
		u.ID, true, time.Now(),
	).Find(&checkmappings)

	var mappings []models.Url_mappings
	err := initializers.DB.Where("Enabled = ? AND user_id = ? AND expiry_date > ?", true, u.ID, time.Now()).Find(&mappings)
	fmt.Println(err)

	var mappings2 []models.Url_mappings
	hi := initializers.DB.Where("Enabled = ? AND user_id = ? AND expiry_date > ?", false, u.ID, time.Now()).Find(&mappings2)
	fmt.Println(hi)

	c.JSON(http.StatusOK, gin.H{
		"Code":         mappings,
		"Disabledcode": mappings2,
		"AmountOfURLs": databaseQuery.RowsAffected,
	})

}
