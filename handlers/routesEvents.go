package handlers

import (
	"events_go/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RoutesEvents() {
	r := gin.Default()

	r.GET("/eventsdb", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, db.Events)
	})

	v1 := r.Group("/api/v1/events")
	{
		// check health
		v1.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "Events ongoing!")
		})
		// Post new event with all parametrs
		v1.POST("/post", postEventHandler)

		// Get all events
		v1.GET("/events", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will give all events %v", time.Now().Unix()))
		})
		v1.GET("/eventsdb", func(c *gin.Context) {
			c.IndentedJSON(http.StatusOK, db.Events)
		})

		// Get event by id
		v1.GET("/id", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will give information about event %v", time.Now().Unix()))
		})

		v1.PUT("/event", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will update information about event %v", time.Now().Unix()))
		})
	}
}

func postEventHandler(c *gin.Context) {
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
