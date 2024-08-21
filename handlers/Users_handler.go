package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostUserHandler(c *gin.Context) {
	var json struct {
		Id        uint   `json:"id"`
		Name      string `json:"name"`
		Surname   string `json:"surname"`
		Nickname  string `json:"nickname"`
		Phone     string `json:"phone"`
		Email     string `json:"email"`
		Instagram string `json:"instagram"`
		Telegram  string `json:"telegram"`
		StatusId  string `json:"statusId"`
		Blocked   bool   `json:"blocked"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   "Received",
		"name":      json.Name,
		"id":        json.Id,
		"surname":   json.Surname,
		"nickname":  json.Nickname,
		"phone":     json.Phone,
		"email":     json.Email,
		"instagram": json.Instagram,
		"telegram":  json.Telegram,
		"statusId":  json.StatusId,
		"blocked":   json.Blocked,
	})
}
