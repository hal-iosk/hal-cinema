package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hal-iosk/hal-cinema/model"
	"github.com/hal-iosk/hal-cinema/service"
)

func CreateReserve(c *gin.Context) {
	Reserves, ok := createReserveReq(c)
	if !ok {
		return
	}
	ok = service.Reserve.CanCreates(Reserves)
	if !ok {
		Batequest("その席はだめなの＞＜", c)
		return
	}
	service.Reserve.Creates(Reserves)
	c.JSON(http.StatusCreated, Reserves)
}

func createReserveReq(c *gin.Context) ([]model.Reserve, bool) {
	UserID := GetUserID(c)
	var Reserves []model.Reserve
	err := c.BindJSON(&Reserves)
	if err != nil {
		Batequest(err.Error(), c)
		return nil, false
	}

	for key, _ := range Reserves {
		Reserves[key].CustomerID = UserID
	}
	return Reserves, true
}
