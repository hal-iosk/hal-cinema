package controller

import (
	"net/http"

	"time"

	"strconv"

	"github.com/gin-gonic/gin"
)

var DayFormat = "2006/01/02"
var TimeFormat = "2006/01/02 15:04:05"

func Batequest(msg string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"err": msg,
	})
	c.Abort()
}

func GetDate(key string, c *gin.Context) (time.Time, error) {
	return time.Parse(DayFormat, c.PostForm(key))
}

func GetDateTime(key string, c *gin.Context) (time.Time, error) {
	return time.Parse(TimeFormat, c.PostForm(key))
}

func GetInt(key string, c *gin.Context) (uint, bool) {
	num, err := strconv.Atoi(c.PostForm(key))
	if err != nil {
		Batequest(key+" がちがう希ガス", c)
		return 0, false
	}
	return uint(num), true
}
