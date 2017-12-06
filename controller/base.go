package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var DayFormat = "2006/01/02"
var TimeFormat = "2006-01-02 15:04:05"

func Batequest(msg string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"err": msg,
	})
	c.Abort()
}
