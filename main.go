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

	// がっすん
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/campaign", func(c *gin.Context) {
		c.HTML(http.StatusOK, "campaign.html", nil)
	})
	r.GET("/comingSoon", func(c *gin.Context) {
		c.HTML(http.StatusOK, "comingSoon.html", nil)
	})
	r.GET("/coupleDay", func(c *gin.Context) {
		c.HTML(http.StatusOK, "coupleDay.html", nil)
	})
	r.GET("/customerEdit", func(c *gin.Context) {
		c.HTML(http.StatusOK, "customerEdit.html", nil)
	})
	r.GET("/ladysDay", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ladysDay.html", nil)
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.GET("/mypage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "mypage.html", nil)
	})
	r.GET("/registerComp", func(c *gin.Context) {
		c.HTML(http.StatusOK, "registerComp.html", nil)
	})
	r.GET("/registerConfirm", func(c *gin.Context) {
		c.HTML(http.StatusOK, "registerConfirm.html", nil)
	})
	r.GET("/registerCustomer", func(c *gin.Context) {
		c.HTML(http.StatusOK, "registerCustomer.html", nil)
	})
	r.GET("/seniorDay", func(c *gin.Context) {
		c.HTML(http.StatusOK, "seniorDay.html", nil)
	})
	r.GET("/watchFirm", func(c *gin.Context) {
		c.HTML(http.StatusOK, "watchFirm.html", nil)
	})

	// おーえっち
	r.GET("/seat", func(c *gin.Context) {
		c.HTML(http.StatusOK, "seat.html", nil)
	})
	r.GET("/seatSelection", func(c *gin.Context) {
		c.HTML(http.StatusOK, "seatSelection.html", nil)
	})
	r.GET("/ticket", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ticket.html", nil)
	})
	r.GET("/payment", func(c *gin.Context) {
		c.HTML(http.StatusOK, "payment.html", nil)
	})
	r.GET("/confirm", func(c *gin.Context) {
		c.HTML(http.StatusOK, "confirm.html", nil)
	})
	r.GET("/complete", func(c *gin.Context) {
		c.HTML(http.StatusOK, "complete.html", nil)
	})

	api := r.Group("/api")
	{
		api.GET("/users", controller.CreateUser)
	}

	r.Run(":2000")
}
