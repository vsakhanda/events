package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RoutesUsers() {
	r := gin.Default()

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello, %s!", name)
	})

	v1 := r.Group("/api/v1/user")
	{
		// check health
		v1.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "Events ongoing!")
		})
		// Post new user
		v1.POST("/post", postUserHandler)

		// Get all user
		v1.GET("/users", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will give all users %v", time.Now().Unix()))
		})
		// Get user by id
		v1.GET("/id", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will give information about user %v", time.Now().Unix()))
		})
		v1.PUT("/user", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will update information about user %v", time.Now().Unix()))
		})
	}

}

func postUserHandler(c *gin.Context) {
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
