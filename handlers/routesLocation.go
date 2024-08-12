package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RoutesLocation() {
	r := gin.Default()

	v1 := r.Group("/api/v1/location")
	{
		// Post new user
		v1.POST("/post", postLocationHandler)
		// Get all user
		v1.GET("/locations", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will give information about locations %v", time.Now().Unix()))
		})
		// Get user by id
		v1.GET("/id", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will give information about the location %v", time.Now().Unix()))
		})
	}

}

func postLocationHandler(c *gin.Context) {
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
