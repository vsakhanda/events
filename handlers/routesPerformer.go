package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RoutesPerformer() {

	r := gin.Default()

	v1 := r.Group("/api/v1/performer")
	{
		// check health
		v1.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "Performers ongoing!")
		})
		// Post new performer
		v1.POST("/post", postPerformerHandler)

		// Get all performers
		v1.GET("/performers", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will give all performers %v", time.Now().Unix()))
		})
		// Get performer by id
		v1.GET("/id", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will give information about performer %v", time.Now().Unix()))
		})

		v1.PUT("/user", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will update information about performers %v", time.Now().Unix()))
		})

	}

}

func postPerformerHandler(c *gin.Context) {
	var json struct {
		id          uint   `json:"id" :"id"`
		name        string `json:"name" :"name"`
		surname     string `json:"surname" :"surname"`
		nickname    string `json:"nickname" :"nickname"`
		description string `json:"description" :"description"`
		photoLink   string `json:"photo_link" :"photoLink"`
		instagram   string `json:"instagram" :"instagram"`
		telegram    string `json:"telegram" :"telegram"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     "Received",
		"name":        json.name,
		"id":          json.id,
		"surname":     json.surname,
		"nickname":    json.nickname,
		"instagram":   json.instagram,
		"telegram":    json.telegram,
		"description": json.description,
		"photo_link":  json.photoLink,
	})
}
