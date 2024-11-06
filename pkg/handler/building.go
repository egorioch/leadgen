package handler

import (
	"github.com/gin-gonic/gin"
	"leadgen/pkg/model"
	"net/http"
	"strconv"
)

func (h *ApiHandler) CreateBuilding(c *gin.Context) {
	var building model.Building

	if err := c.ShouldBindJSON(&building); err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't bind json to object"})
		return
	}

	if err := h.s.CreateBuilding(&building); err != nil {
		h.log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can't create building)"})
		return
	}

	c.JSON(http.StatusCreated, building)
	return
}

func (h *ApiHandler) ListBuildings(c *gin.Context) {
	city := c.Query("city")
	yearBuiltStr := c.Query("year_built")
	floorsStr := c.Query("floors")

	// Преобразуем параметры `year_built` и `floors` в int, если они указаны
	var yearBuilt, floors int
	var err error
	if yearBuiltStr != "" {
		yearBuilt, err = strconv.Atoi(yearBuiltStr)
		if err != nil {
			h.log.Error(err.Error())

			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year_built parameter"})
			return
		}
	}
	if floorsStr != "" {
		floors, err = strconv.Atoi(floorsStr)
		if err != nil {
			h.log.Error(err.Error())

			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid floors parameter"})
			return
		}
	}

	buildings, err := h.s.ListBuildings(city, yearBuilt, floors)
	if err != nil {
		h.log.Error(err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch buildings"})
		return
	}

	// Возвращаем список строений
	c.JSON(http.StatusOK, buildings)
	return
}
