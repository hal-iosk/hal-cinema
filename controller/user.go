package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	name string
}

func NewUserController() User {
	return User{
		name: "junjun",
	}
}

func (self *User) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "junjun",
	})
}
