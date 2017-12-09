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
	authAdmin.Use(session.AdminAuthApiMiddleware)

	api.POST("/signup", controller.CreateUser)
	api.POST("/signin", session.UserLoginHandle)
	api.GET("/movie", controller.GetMovieAll)
	api.GET("/movie/:movie_id", controller.GetMovie)
	api.GET("/schedule", controller.GetScheduleAll)
	api.GET("/schedule/:schedule_id", controller.GetSchedule)
	api.GET("/price", controller.GetPrice)
	api.GET("/reserved/:schedule_id", controller.GetRserved)

	authApi.POST("/reserve", controller.CreateReserve)
	authApi.PUT("/user", controller.UpdateUser)
	authApi.POST("/checkin", controller.Checkin)
	authApi.POST("/popcorn", controller.Popcorn)

	admin.POST("/signin", session.AdminLoginHandle)

	authAdmin.POST("/movie", controller.CreateMovie)
	authAdmin.PUT("/movie/:movie_id", controller.UpdateMovie)
	authAdmin.DELETE("/movie/:movie_id", controller.DeleteMovie)
	authAdmin.POST("/schedule", controller.CreateSchedule)
	authAdmin.PUT("/schedule/:schedule_id", controller.UpdateSchedule)
	authAdmin.DELETE("/schedule/:schedule_id", controller.DeleteSchedule)

}
