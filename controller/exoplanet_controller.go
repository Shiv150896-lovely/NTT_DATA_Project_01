package controller

import (
	"net/http"
	"strconv"

	"Ntt_DATA/service"

	"github.com/gin-gonic/gin"
)

func AddExoplanet(c *gin.Context) {
	var newExoplanet service.Exoplanet
	if err := c.ShouldBindJSON(&newExoplanet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.ValidateExoplanet(newExoplanet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdExoplanet := service.AddExoplanet(newExoplanet)
	c.JSON(http.StatusCreated, createdExoplanet)
}

func ListExoplanets(c *gin.Context) {
	exoplanets := service.ListExoplanets()
	c.JSON(http.StatusOK, exoplanets)
}

func GetExoplanetByID(c *gin.Context) {
	id := c.Param("id")
	exoplanet, err := service.GetExoplanetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exoplanet not found"})
		return
	}
	c.JSON(http.StatusOK, exoplanet)
}

func UpdateExoplanet(c *gin.Context) {
	id := c.Param("id")
	var updatedExoplanet service.Exoplanet
	if err := c.ShouldBindJSON(&updatedExoplanet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exoplanet, err := service.UpdateExoplanet(id, updatedExoplanet)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exoplanet not found"})
		return
	}

	c.JSON(http.StatusOK, exoplanet)
}

func DeleteExoplanet(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteExoplanet(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exoplanet not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func FuelEstimation(c *gin.Context) {
	id := c.Param("id")
	crewCapacityStr := c.Query("crewCapacity")
	crewCapacity, err := strconv.Atoi(crewCapacityStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid crew capacity"})
		return
	}

	fuel, err := service.FuelEstimation(id, crewCapacity)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"fuelEstimation": fuel})
}
