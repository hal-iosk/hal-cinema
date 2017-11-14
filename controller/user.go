package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	name string
}

func CreateUser(c *gin.Context) {
	// _ := c.PostForm("name")
	// _ := c.PostForm("password")

	c.JSON(http.StatusCreated, nil)
}
