package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexView(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
