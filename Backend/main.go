package main

import (
	"keceox_modules/controllers"
	"keceox_modules/initializers"
	"keceox_modules/middleware"
	"keceox_modules/seeders"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()

	seeders.RoleSeeder()
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, middleware.RequireAdmin, controllers.Validate)
	router.POST("/shorten", middleware.RequireAuth, controllers.ShortenURL)
	router.DELETE("/deleteaccount", middleware.RequireAuth, controllers.DeleteAccount)
	router.PUT("/changename", middleware.RequireAuth, controllers.ChangeName)

	router.GET("/link/:id", controllers.GetOriginalURL)

	router.PUT("/disable/:id", middleware.RequireAuth, controllers.DisableURL)

	router.GET("/getactive", middleware.RequireAuth, middleware.RequireAdmin, controllers.AdminSeeAllActiveURLS)
	router.PUT("/admindisableurl/:id", middleware.RequireAuth, middleware.RequireAdmin, controllers.AdminDisableURL)

	router.Run() // listens on 0.0.0.0:8080 by default
}
