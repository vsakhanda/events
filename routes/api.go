package routes

import (
	"events_go/db"
	"events_go/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	//r.LoadHTMLGlob("front/templates/*")
	r.Static("/static", "./front/static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"version": fmt.Sprintf("%v", time.Now().Unix()),
			"title":   "Title of site",
		})
	})

	r.GET("/eventsdb", func(c *gin.Context) {
		c.IndentedJSON(
			http.StatusOK,
			db.Events,
		)
	})

	v1 := r.Group("/api/v1/")
	{
		// check health
		v1.GET("events/", func(c *gin.Context) {
			c.String(http.StatusOK, "Events ongoing!")
		})
		// Post new event with all parametrs
		v1.POST("events/post", handlers.PostEventHandler)

		// Post new event with all parametrs
		v1.POST("events/post1/", handlers.PostEventHandler1)

		// З параметром в запиті
		//v1.POST("events/post1/:id", handlers.PostEventHandler1)

		// Get all events
		v1.GET("events/events", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will give all events %v", time.Now().Unix()))
		})
		v1.GET("events/eventsdb", func(c *gin.Context) {
			c.IndentedJSON(http.StatusOK, db.Events)
		})
		// Get event by id
		v1.GET("events/id", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will give information about event %v", time.Now().Unix()))
		})
		v1.PUT("events/event", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will update information about event %v", time.Now().Unix()))
		})

		v1.POST("location/post", handlers.PostLocationHandler)
		// Get all user
		v1.GET("location/locations", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will give information about locations %v", time.Now().Unix()))
		})
		// Get user by id
		v1.GET("location/id", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will give information about the location %v", time.Now().Unix()))
		})

		// check health
		v1.GET("performer/", func(c *gin.Context) {
			c.String(http.StatusOK, "Performers ongoing!")
		})
		// Post new performer
		v1.POST("performer/post", handlers.PostPerformerHandler)

		// Get all performers
		v1.GET("performer/performers", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will give all performers %v", time.Now().Unix()))
		})
		// Get performer by id
		v1.GET("performer/id", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will give information about performer %v", time.Now().Unix()))
		})

		v1.PUT("performer/user", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will update information about performers %v", time.Now().Unix()))
		})

		v1.GET("user/", func(c *gin.Context) {
			c.String(http.StatusOK, "Events ongoing!")
		})
		v1.POST("user/post", handlers.PostUserHandler)

		v1.GET("user/users", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will give all users %v", time.Now().Unix()))
		})
		// Get user by id
		v1.GET("user/id", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will give information about user %v", time.Now().Unix()))
		})
		v1.PUT("user/user", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("In progress! API will update information about user %v", time.Now().Unix()))
		})
	}
	return r
}
