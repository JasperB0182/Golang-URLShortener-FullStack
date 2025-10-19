package controllers

import (
	"keceox_modules/initializers"
	"keceox_modules/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// GET EMAIL/PASSWORD
	var body struct {
		Email    string
		Password string
		Name     string
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
	user := models.User{Email: body.Email, Password: string(hash), Name: body.Name, RoleID: 1}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "FAILED TO CREATE USER!",
		})
	}

	var newuser models.User
	initializers.DB.First(&newuser, "email = ?", body.Email)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": newuser.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "FAILED TO CREATE TOKEN!",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", tokenString, 3600*24, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"Success": "Succesfully created user and logged in!",
	})
}

func Login(c *gin.Context) {
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

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email/Password!",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email/Password!",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "FAILED TO CREATE TOKEN!",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", tokenString, 3600*24, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})

}

func Logout(c *gin.Context) {
	c.SetCookie("Auth", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

func AdminCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func DeleteAccount(c *gin.Context) {
	user, _ := c.Get("user")

	u := user.(models.User)

	initializers.DB.Delete(&u, u.ID)
	// DELETE FROM users WHERE id = 10;

	c.JSON(http.StatusOK, gin.H{
		"message": "get rekt",
	})

}

func ChangeName(c *gin.Context) {
	var Name struct {
		Name string
	}

	if err := c.BindJSON(&Name); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	user, _ := c.Get("user")
	u := user.(models.User)
	u.Name = Name.Name
	initializers.DB.Save(&u)
}
