package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hal-iosk/hal-cinema/middleware"
	cors "github.com/itsjamie/gin-cors"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Static("/js", "./public/js")
	r.Static("/image", "./public/image")
	r.Static("/css", "./public/css")

	r.LoadHTMLGlob("view/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	api := r.Group("/api")

	api.Use(cors.Middleware(middleware.CorsConfig))

	api.GET("/makki", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":  "katsuramaki taiki",
			"sex":   "man",
			"email": "llxo2_5oxll@icloud.com",
		})
	})

	r.Run(":3000")
}
