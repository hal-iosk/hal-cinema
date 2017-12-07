package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hal-iosk/hal-cinema/model"
	"github.com/hal-iosk/hal-cinema/service"
)

func CreateReserve(c *gin.Context) {
	Reserve, ok := createReserveReq(c)
	if !ok {
		return
	}
	if !service.Reserve.CanCreate(Reserve) {
		Batequest("そのよやくうまっちゃってますね＾＾；", c)
		return
	}
	Reserve = service.Reserve.Create(Reserve)
	c.JSON(http.StatusCreated, Reserve)
}

func createReserveReq(c *gin.Context) (model.Reserve, bool) {
	Reserve := model.Reserve{SeatID: c.PostForm("seat_id")}
	ScheduleID, ok := GetInt("schedule_id", c)
	if !ok {
		return Reserve, false
	}
	CustomerID, _ := c.Get("userID")
	PriceID, ok := GetInt("price_id", c)
	if !ok {
		return Reserve, false
	}
	Reserve.ScheduleID = ScheduleID
	Reserve.CustomerID = CustomerID.(uint)
	Reserve.PriceID = PriceID
	return Reserve, true
}
