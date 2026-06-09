package main

import (
	"fmt"
	"keceox_modules/controllers"
	"keceox_modules/initializers"
	"keceox_modules/middleware"
	"keceox_modules/seeders"
	"os"
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
	router.Use(middleware.RateLimiter())

	info, err := os.Stat("logs.txt")
	if os.IsNotExist(err) {
		LogFile, error := os.Create("logs.txt")
		fmt.Println(info)
		if error != nil {
		}

		defer LogFile.Close()
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("CORS")},
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

	api := router.Group("/api")
	{
		api.POST("/signup", controllers.Signup)
		api.POST("/login", controllers.Login)
		api.POST("/logout", middleware.RequireAuth, controllers.Logout)
		api.GET("/validate", middleware.RequireAuth, controllers.Validate)
		api.POST("/shorten", middleware.RequireAuth, controllers.ShortenURL)
		api.DELETE("/deleteaccount", middleware.RequireAuth, controllers.DeleteAccount)
		api.PUT("/changename", middleware.RequireAuth, controllers.ChangeName)
		api.PUT("/changeemail", middleware.RequireAuth, controllers.ChangeEmail)
		api.PUT("/changepassword", middleware.RequireAuth, controllers.ChangePassword)

		api.GET("/link/:id", controllers.GetOriginalURL)
		api.GET("/getmyurls", middleware.RequireAuth, controllers.GetAllMyURLS)

		api.PUT("/disable/:id", middleware.RequireAuth, controllers.DisableURL)
		api.PUT("/enable/:id", middleware.RequireAuth, controllers.EnableURL)

		api.GET("/admincheck", middleware.RequireAuth, middleware.RequireAdmin, controllers.AdminCheck)

		api.GET("/getactive", middleware.RequireAuth, middleware.RequireAdmin, controllers.AdminSeeAllActiveURLS)
		api.GET("/getusers", middleware.RequireAuth, middleware.RequireAdmin, controllers.AdminSeeAllUsers)
		api.DELETE("/deleteaccountadmin/:id", middleware.RequireAuth, middleware.RequireAdmin, controllers.AdminDeleteAccount)
		api.PUT("/admindisableurl/:id", middleware.RequireAuth, middleware.RequireAdmin, controllers.AdminDeleteURL)
		api.PUT("/admindisablemultipleurl", middleware.RequireAuth, middleware.RequireAdmin, controllers.AdminDisableMultipleURL)

		api.PUT("/addtocredit", middleware.RequireAuth, controllers.AddToCredit)
		api.GET("/getcredit", middleware.RequireAuth, controllers.GetCreditAndURLs)
	}

	router.Run() // listens on 0.0.0.0:8080 by default
}
