package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hal-iosk/hal-cinema/controller"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Static("/js", "./public/js")
	r.Static("/image", "./public/image")
	r.Static("/css", "./public/css")
	r.Static("/media", "./public/media")

	r.LoadHTMLGlob("view/*")

	r.GET("/live", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"live": true,
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "seat.html", nil)
	})
	r.GET("/seatSelection", func(c *gin.Context) {
		c.HTML(http.StatusOK, "seatSelection.html", nil)
	})

	api := r.Group("/api")
	{
		api.GET("/users", controller.CreateUser)
	}

	r.Run(":2000")
}
