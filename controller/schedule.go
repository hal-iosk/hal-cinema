package controller

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hal-iosk/hal-cinema/model"
	"github.com/hal-iosk/hal-cinema/service"
)

func GetScheduleAll(c *gin.Context) {
	movieID, err := strconv.Atoi(c.Query("movie_id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"schedules": service.Schedule.All(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"schedules": service.Schedule.FindByMovie(uint(movieID)),
	})
}

func GetSchedule(c *gin.Context) {
	scheduleID, err := strconv.Atoi(c.Param("schedule_id"))
	if err != nil {
		Batequest("数値で頼むわ", c)
		return
	}
	schedule := service.Schedule.Find(uint(scheduleID))
	if schedule == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "んなもんねーよ！笑",
		})
		return
	}
	c.JSON(http.StatusOK, schedule)
}

func CreateSchedule(c *gin.Context) {
	schedule, ok := createScheduleReq(c)
	if !ok {
		return
	}
	// createはリリース禁止
	schedule.Release = false

	schedule = service.Schedule.Create(schedule)
	c.JSON(http.StatusCreated, schedule)
}

func UpdateSchedule(c *gin.Context) {
	scheduleID, err := strconv.Atoi(c.Param("schedule_id"))
	if err != nil {
		Batequest("数値で頼むわ", c)
		return
	}
	if !service.Schedule.Exists(uint(scheduleID)) {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "んなもんねーよ！笑",
		})
	}
	if service.Schedule.IsRelease(uint(scheduleID)) {
		c.JSON(http.StatusForbidden, gin.H{
			"err": "リリースだからupdateむっりー笑笑",
		})
		return
	}
	schedule, ok := createScheduleReq(c)
	if !ok {
		return
	}
	scheduleRes := service.Schedule.Update(uint(scheduleID), schedule)
	c.JSON(http.StatusOK, scheduleRes)
}

func DeleteSchedule(c *gin.Context) {
	scheduleID, err := strconv.Atoi(c.Param("schedule_id"))
	if err != nil {
		Batequest("数値で頼むわ", c)
		return
	}
	if !service.Schedule.Exists(uint(scheduleID)) {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "んなもんねーよ！笑",
		})
	}
	service.Schedule.Delete(uint(scheduleID))
	c.Writer.WriteHeader(http.StatusNoContent)
}
func createScheduleReq(c *gin.Context) (model.ScreeningSchedule, bool) {
	schedule := model.ScreeningSchedule{Release: false}
	if c.PostForm("release") == "true" {
		schedule.Release = true
	}
	movieID, err := strconv.Atoi(c.PostForm("movie_id"))
	if err != nil {
		Batequest("movie_id 数字入れろカス〜", c)
		return schedule, false
	}
	schedule.MovieID = uint(movieID)

	TheaterNumber, err := strconv.Atoi(c.PostForm("theater_number"))
	if err != nil {
		Batequest("theater_number 数字入れろカス〜", c)
		return schedule, false
	}
	schedule.TheaterNumber = uint(TheaterNumber)

	schedule.StartTime, err = GetDateTime("start_time", c)
	if err != nil {
		panic(err)
		Batequest("start_time がちゃうで", c)
		return schedule, false
	}
	return schedule, true
}
