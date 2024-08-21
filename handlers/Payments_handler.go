package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func PostPaymentHandler(c *gin.Context) {
	var json struct {
		id          uint      `json:"id" :"id"`
		date        time.Time `json:"date" :"date"`
		userid      string    `json:"userid" :"userid"`
		eventId     string    `json:"event_id" :"event_id"`
		description string    `json:"description" :"description"`
		account     string    `json:"account" :"account"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     "Received",
		"id":          json.id,
		"date":        json.date,
		"userid":      json.userid,
		"eventId":     json.eventId,
		"description": json.description,
		"account":     json.account,
	})
}
