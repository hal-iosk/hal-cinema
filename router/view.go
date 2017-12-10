package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func viewRouter(r *gin.Engine) {
	// がっすん
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/campaign", func(c *gin.Context) {
		c.HTML(http.StatusOK, "campaign.html", nil)
	})
	r.GET("/cinemaDay", func(c *gin.Context) {
		c.HTML(http.StatusOK, "cinemaDay.html", nil)
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
	r.GET("/shop", func(c *gin.Context) {
		c.HTML(http.StatusOK, "shop.html", nil)
	})
	r.GET("/watchFilm", func(c *gin.Context) {
		c.HTML(http.StatusOK, "watchFilm.html", nil)
	})

	r.GET("/reserve", func(c *gin.Context) {
		c.HTML(http.StatusOK, "reserve.html", nil)
	})
	r.GET("/reserve/:path", func(c *gin.Context) {
		c.HTML(http.StatusOK, "reserve.html", nil)
	})

	// admin
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "adminCtrl.html", nil)
	})
	r.GET("/admin/:path/:hoge", func(c *gin.Context) {
		c.HTML(http.StatusOK, "adminCtrl.html", nil)
	})
	r.GET("/admin/:path", func(c *gin.Context) {
		path := c.Param("path")
		if path == "login" {
			c.HTML(http.StatusOK, "adminLogin.html", nil)
		} else {
			c.HTML(http.StatusOK, "adminCtrl.html", nil)
		}
	})

	// checkin
	r.GET("/checkin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "checkin.html", nil)
	})
	r.GET("/checkin/:path", func(c *gin.Context) {
		c.HTML(http.StatusOK, "checkin.html", nil)
	})
}
