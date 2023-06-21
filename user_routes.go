package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/username/repo/utils"
	"github.com/username/repo/controllers"
	"github.com/username/repo/models"
)

func UserRoutes(router *gin.RouterGroup) {

	router.POST("/register", func(c *gin.Context) {
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}
		if err := controllers.Register(user); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
	})

	router.POST("/login", func(c *gin.Context) {
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}
		token, err := controllers.Login(user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	userRoutes := router.Group("/user", utils.RequireTokenAuth())
	{
		userRoutes.GET("/", func(c *gin.Context) {
			users, err := controllers.GetAllUsers()
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"users": users})
		})

		userRoutes.GET("/:id", func(c *gin.Context) {
			user, err := controllers.GetUserById(c.Param("id"))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"user": user})
		})

		userRoutes.PUT("/:id", func(c *gin.Context) {
			var user models.User
			if err := c.BindJSON(&user); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
				return
			}
			updatedUser, err := controllers.UpdateUser(c.Param("id"), user)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"user": updatedUser})
		})

		userRoutes.DELETE("/:id", func(c *gin.Context) {
			if err := controllers.DeleteUser(c.Param("id")); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
		})
	}

	router.GET("/analytix", func(c *gin.Context) {
		data, err := controllers.Analytix()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve analytix data"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	})
}