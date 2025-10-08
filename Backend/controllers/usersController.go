package controllers

import (
	"keceox_modules/initializers"
	"keceox_modules/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// GET EMAIL/PASSWORD
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "FAILED TO READ BODY!",
		})

		return
	}

	// HASH IT

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "FAILED PASSWORD HASH!",
		})
	}

	// CREATE USER
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "FAILED TO CREATE USER!",
		})
	}

	// RESPOND

	c.JSON(http.StatusOK, gin.H{
		"Success": "Succesfully created user!",
	})
}
