package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RoutesPayment() {
	r := gin.Default()

	v1 := r.Group("/api/v1/user")
	{
		// Post new user
		v1.POST("/post", postPaymentHandler)
		// Get all user
		v1.GET("/payments", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will give information about payments in thse system %v", time.Now().Unix()))
		})
		// Get user by id
		v1.GET("/id", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will give information about payment %v", time.Now().Unix()))
		})
	}

}

func postPaymentHandler(c *gin.Context) {
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
