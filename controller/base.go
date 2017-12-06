package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Batequest(msg string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"err": msg,
	})
	c.Abort()
}
