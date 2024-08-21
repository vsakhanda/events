package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func PostEventHandler(c *gin.Context) {
	var json struct {
		Id          uint          `json:"Id"`
		Name        string        `json:"name"`
		TypeId      uint          `json:"typeId"`
		LocationId  uint          `json:"locationId"`
		Capacity    uint          `json:"capacity"`
		Description string        `json:"description"`
		PerformerId uint          `json:"performerId"`
		Duration    time.Duration `json:"duration"`
		StartTime   time.Time     `json:"startTime"`
		EndTime     time.Time     `json:"endTime"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     "Received",
		"name":        json.Name,
		"Id":          json.Id,
		"typeId":      json.TypeId,
		"locationId":  json.LocationId,
		"capacity":    json.Capacity,
		"description": json.Description,
		"performerId": json.PerformerId,
		"duration":    json.Duration,
		"startTime":   json.StartTime,
		"endTime":     json.EndTime,
	})
}

func PostEventHandler1(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Received",
		"Id":      id,
		"name":    name,
	})
}
