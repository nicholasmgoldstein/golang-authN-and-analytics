package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myapp/api/dto"
	"github.com/myapp/api/models"
	"github.com/myapp/api/services"
	"github.com/myapp/api/utils"
)

type AuthController struct{}

// Register registers a new user
func (a *AuthController) Register(c *gin.Context) {
	var input dto.RegisterInput

	// validate request body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if user already exists
	if services.UserExists(input.Username) {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	// hash user password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to hash password"})
		return
	}

	// create user in database
	user := models.User{
		Username: input.Username,
		Password: hashedPassword,
	}
	if err := services.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create user"})
		return
	}

	// generate jwt token
	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to generate token"})
		return
	}

	// return response
	c.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}

// Login logs in a user
func (a *AuthController) Login(c *gin.Context) {
	var input dto.LoginInput

	// validate request body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if user exists
	user, err := services.GetUserByUsername(input.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// verify user password
	if !utils.CheckPassword(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// generate jwt token
	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to generate token"})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// Logout logs out a user
func (a *AuthController) Logout(c *gin.Context) {
	// do any necessary cleanup for logout, e.g. deleting refresh token, etc.
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
