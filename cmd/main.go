package main

import (
	"events_go/db"
	"events_go/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	handlers.RoutesEvents()
	handlers.RoutesUsers()
	handlers.RoutesPerformer()
	handlers.RoutesLocation()
	handlers.RoutesPayment()

	//r.LoadHTMLGlob("front/templates/*")
	r.Static("/static", "./front/static")

	r.GET("/1", func(c *gin.Context) {
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

	// Запуск сервера
	err := r.Run(":9090")
	if err != nil {
		return
	}

}
