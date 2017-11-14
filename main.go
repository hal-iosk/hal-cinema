package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hal-iosk/hal-cinema/controller"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Static("/js", "./public/js")
	r.Static("/image", "./public/image")
	r.Static("/css", "./public/css")
	r.LoadHTMLGlob("view/*")

	api := r.Group("/api")
	{
		api.GET("/users", controller.CreateUser)
	}

	r.Run(":2000")
}
