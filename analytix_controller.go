package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myapp/models"
	"github.com/myapp/utils"
)

type AnalytixController struct{}

var Analytix AnalytixController

// POST /analytix
func (a *AnalytixController) CreateAnalytix(c *gin.Context) {
	var analytix models.Analytix

	err := json.NewDecoder(c.Request.Body).Decode(&analytix)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := analytix.Validate(); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := analytix.CreateAnalytix(); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusCreated, analytix)
}

// GET /analytix/:id
func (a *AnalytixController) GetAnalytix(c *gin.Context) {
	id := c.Param("id")

	analytix := models.Analytix{ID: id}
	if err := analytix.GetAnalytix(); err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Analytix not found")
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, analytix)
}

// PUT /analytix/:id
func (a *AnalytixController) UpdateAnalytix(c *gin.Context) {
	id := c.Param("id")

	analytix := models.Analytix{ID: id}
	err := json.NewDecoder(c.Request.Body).Decode(&analytix)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := analytix.Validate(); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := analytix.UpdateAnalytix(); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, analytix)
}

// DELETE /analytix/:id
func (a *AnalytixController) DeleteAnalytix(c *gin.Context) {
	id := c.Param("id")

	analytix := models.Analytix{ID: id}
	if err := analytix.DeleteAnalytix(); err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Analytix not found")
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, gin.H{"result": "success"})
}