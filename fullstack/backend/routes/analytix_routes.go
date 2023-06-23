package routes

import (
	"net/http"
	"project/controllers"
	"project/middleware"
	"github.com/gin-gonic/gin"
)

func SetupAnalytixRoutes(router *gin.Engine) {
	analytixGroup := router.Group("/api/analytix")
	analytixGroup.Use(middleware.Authenticate()) // middleware to require authentication
	analytixGroup.Use(middleware.RequireUserRole("admin")) // middleware to require admin role

	// routes to handle CRUD operations on analytix data
	analytixGroup.GET("/", controllers.GetAllAnalytixData)
	analytixGroup.GET("/:id", controllers.GetAnalytixDataById)
	analytixGroup.POST("/", controllers.CreateAnalytixData)
	analytixGroup.PUT("/:id", controllers.UpdateAnalytixData)
	analytixGroup.DELETE("/:id", controllers.DeleteAnalytixData)
}
