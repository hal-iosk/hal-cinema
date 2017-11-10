package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hal-iosk/hal-cinema/controller"
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

	r.GET("/live", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"live": true,
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	userController := controller.NewUserController()
	api := r.Group("/api")
	api.Use(cors.Middleware(middleware.CorsConfig))
	api.GET("/users", userController.Create)

	r.Run(":2000")
}
