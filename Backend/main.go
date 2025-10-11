package main

import (
	"keceox_modules/controllers"
	"keceox_modules/initializers"
	"keceox_modules/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)
	router.POST("/shorten", middleware.RequireAuth, controllers.ShortenURL)
	router.DELETE("/deleteaccount", middleware.RequireAuth, controllers.DeleteAccount)
	router.PUT("/changename", middleware.RequireAuth, controllers.ChangeName)

	router.GET("/link/:id", middleware.RequireAuth, controllers.GetOriginalURL)

	router.PUT("/disable/:id", middleware.RequireAuth, controllers.DisableURL)

	router.Run() // listens on 0.0.0.0:8080 by default
}
