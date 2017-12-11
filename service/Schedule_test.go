package service

import (
	"testing"
	"time"

	"fmt"

	"github.com/hal-iosk/hal-cinema/model"
)

func TestFoo(t *testing.T) {
	startTime, err := time.Parse("2006/01/02 15:04:05", "2017/12/11 10:00:00")
	startTime.Add()
	//startTime = time.Date(2017, 12, 11, 10, 0, 0, 0, time.Local)
	fmt.Println(time.Local.String())
	if err != nil {
		panic(err)
	}
	//fmt.Println(startTime)
	Schedule.Create(model.ScreeningSchedule{
		MovieID:       88,
		StartTime:     startTime,
		TheaterNumber: 7,
	})
	fmt.Println(Schedule.FindByMovie(88))
}
func TestHoge(t *testing.T) {
	//startTime, err := time.Parse("2006/01/02 15:04:05", "2017/12/11 10:00:00")
	//startTime.Add()
	////startTime = time.Date(2017, 12, 11, 10, 0, 0, 0, time.Local)
	//fmt.Println(time.Local.String())
	//if err != nil {
	//	panic(err)
	//}
	////fmt.Println(startTime)
	//Schedule.Create(model.ScreeningSchedule{
	//	MovieID:       88,
	//	StartTime:     startTime,
	//	TheaterNumber: 7,
	//})
	fmt.Println(Schedule.FindByMovie(88))
}
