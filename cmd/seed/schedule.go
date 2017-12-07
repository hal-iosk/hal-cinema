package seed

import (
	"time"

	"github.com/hal-iosk/hal-cinema/model"
	"github.com/hal-iosk/hal-cinema/service"
)

func CreateSchedule() {
	service.Schedule.Create(model.ScreeningSchedule{
		MovieID:       1,
		StartTime:     time.Date(2017, 1, 18, 0, 0, 0, 0, loc),
		TheaterNumber: 2,
		Release:       true,
	})
}
