package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostPerformerHandler(c *gin.Context) {
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
