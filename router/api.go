package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hal-iosk/hal-cinema/controller"
)

func apiRouter(api *gin.RouterGroup) {
	api.GET("/users", controller.CreateUser)
}
