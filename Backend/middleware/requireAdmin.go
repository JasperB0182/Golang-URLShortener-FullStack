package middleware

import (
	"keceox_modules/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireAdmin(c *gin.Context) {
	user, _ := c.Get("user")

	u := user.(models.User)

	if u.RoleID == 1 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	} else if u.RoleID == 2 {
		c.Next()
	}
}
