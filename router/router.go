package router

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {

	r := gin.Default()
	r.Static("/js", "./public/js")
	r.Static("/image", "./public/image")
	r.Static("/css", "./public/css")
	r.Static("/media", "./public/media")

	r.LoadHTMLGlob("view/*")
	viewRouter(r)
	api := r.Group("/api")
	apiRouter(api)
	return r
}
