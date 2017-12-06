package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hal-iosk/hal-cinema/controller"
	"github.com/hal-iosk/hal-cinema/session"
)

func apiRouter(api *gin.RouterGroup) {
	authApi := api.Group("")
	authApi.Use(session.AuthApiMiddleware)
	admin := api.Group("/admin")
	authAdmin := admin.Group("")
	authAdmin.Use(session.AuthApiMiddleware)

	api.POST("/signup", controller.CreateUser)
	api.POST("/signin", session.UserLoginHandle)

	authApi.POST("/foo", func(c *gin.Context) {
		userID, _ := c.Get("userID")
		c.JSON(200, userID)
	})

	admin.POST("/signin", session.AdminLoginHandle)

}
