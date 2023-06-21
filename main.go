package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"your-module-name/config"
	"your-module-name/controllers"
	"your-module-name/middleware"
	"your-module-name/models"
)

func main() {
	// Load configuration from YAML file
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Connect to database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		viper.GetString("database.host"),
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.name"),
		viper.GetInt("database.port"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate database schema
	err = db.AutoMigrate(&models.User{}, &models.Session{}, &models.Analytix{})
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	// Set up Gin router
	router := gin.Default()

	// Use middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Auth())

	// Set up user routes
	userController := controllers.NewUserController(db)
	router.POST("/users", userController.CreateUser)
	router.GET("/users", userController.ListUsers)
	router.GET("/users/:id", userController.GetUser)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)

	// Set up auth routes
	authController := controllers.NewAuthController(db)
	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)
	router.POST("/logout", authController.Logout)

	// Set up analytix routes
	analytixController := controllers.NewAnalytixController(db)
	router.POST("/analytix", analytixController.CreateAnalytix)
	router.GET("/analytix", analytixController.ListAnalytix)
	router.GET("/analytix/:id", analytixController.GetAnalytix)
	router.PUT("/analytix/:id", analytixController.UpdateAnalytix)
	router.DELETE("/analytix/:id", analytixController.DeleteAnalytix)

	// Start server
	addr := fmt.Sprintf(":%d", viper.GetInt("port"))
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.name"),
	)
}