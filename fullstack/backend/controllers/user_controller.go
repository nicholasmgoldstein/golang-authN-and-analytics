package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myapp/models"
	"github.com/myapp/utils"
)

// UserController handles HTTP requests for the user resource.
type UserController struct{}

// Create creates a new user.
func (ctl UserController) Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.HandleError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := user.Validate(); err != nil {
		utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Create user in database
	createdUser, err := models.CreateUser(user)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": createdUser})
}

// GetByID retrieves a user by ID.
func (ctl UserController) GetByID(c *gin.Context) {
	userID := c.Param("id")
	user, err := models.GetUserByID(userID)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, "User not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Update updates an existing user.
func (ctl UserController) Update(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.HandleError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := user.Validate(); err != nil {
		utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Update user in database
	updatedUser, err := models.UpdateUser(userID, user)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updatedUser})
}

// Delete deletes an existing user.
func (ctl UserController) Delete(c *gin.Context) {
	userID := c.Param("id")

	if err := models.DeleteUser(userID); err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
