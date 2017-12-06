package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hal-iosk/hal-cinema/service"
)

func GetMovieAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"movies": service.Movie.All(),
	})
}
