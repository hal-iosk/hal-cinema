package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hal-iosk/hal-cinema/controller"
	"github.com/hal-iosk/hal-cinema/session"
)

func apiRouter(api *gin.RouterGroup) {
	api.POST("/signup", controller.CreateUser)
	api.POST("/signin", session.UserLoginHandle)
	authApi := api.Group("")
	authApi.Use(session.AuthApiMiddleware)
	authApi.POST("/foo", func(c *gin.Context) {
		userID, _ := c.Get("userID")
		c.JSON(200, userID)
	})
}
