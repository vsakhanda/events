package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostLocationHandler(c *gin.Context) {
	var json struct {
		Name        string `json:"name"`
		Capacity    string `json:"capacity"`
		Description string `json:"description"`
		Link        string `json:"link"`
		Lat         string `json:"lat"`
		Lon         string `json:"lon"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     "Received",
		"name":        json.Name,
		"capacity":    json.Capacity,
		"description": json.Description,
		"link":        json.Link,
		"lat":         json.Lat,
		"lon":         json.Lon,
	})
}
