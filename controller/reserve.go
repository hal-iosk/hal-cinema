package controller

import (
	"net/http"

	"strconv"

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
func GetRserved(c *gin.Context) {
	var seats []string
	ScheduleID, err := strconv.Atoi(c.Param("schedule_id"))
	if err != nil {
		Batequest("数値でおなしゃす", c)
		return
	}
	Reserves := service.Reserve.FindByScheduleID(uint(ScheduleID))
	for _, value := range Reserves {
		seats = append(seats, value.SeatID)
	}
	c.JSON(http.StatusOK, gin.H{
		"seats": seats,
	})
}
func GetUserRserved(c *gin.Context) {
	UserID := GetUserID(c)
	reserved := service.Reserve.FindByUserID(UserID)
	c.JSON(http.StatusOK, gin.H{
		"reserved": reserved,
	})
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
