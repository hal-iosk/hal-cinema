package seed

import (
	"fmt"
	"time"

	"github.com/hal-iosk/hal-cinema/model"
	"github.com/hal-iosk/hal-cinema/service"
)

func CreateSchedule() {

	var year1 int
	var month1 int
	var day1 int

	var year2 int
	var month2 int
	var day2 int

	var year3 int
	var month3 int
	var day3 int

	var year4 int
	var month4 int
	var day4 int

	var year5 int
	var month5 int
	var day5 int

	var year6 int
	var month6 int
	var day6 int

	var year7 int
	var month7 int
	var day7 int

	d1 := time.Now()
	d2 := d1.Add(24 * time.Hour)
	d3 := d2.Add(24 * time.Hour)
	d4 := d3.Add(24 * time.Hour)
	d5 := d4.Add(24 * time.Hour)
	d6 := d5.Add(24 * time.Hour)
	d7 := d6.Add(24 * time.Hour)

	year1 = (d1.Year())
	month1 = int(d1.Month())
	day1 = (d1.Day())

	year2 = (d2.Year())
	month2 = int(d2.Month())
	day2 = (d2.Day())

	year3 = (d3.Year())
	month3 = int(d3.Month())
	day3 = (d3.Day())

	year4 = (d4.Year())
	month4 = int(d4.Month())
	day4 = (d4.Day())

	year5 = (d5.Year())
	month5 = int(d5.Month())
	day5 = (d5.Day())

	year6 = (d6.Year())
	month6 = int(d6.Month())
	day6 = (d6.Day())

	year7 = (d7.Year())
	month7 = int(d7.Month())
	day7 = (d7.Day())

	var baseH int
	var baseM int
	var hour int
	var minute int

	baseH = 9
	baseM = 0

	//1日目
	for i := 1; i <= 5; i++ {

		hour = baseH
		minute = baseM

		switch i {
		case 1:
			for k := 1; k <= 5; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year1, time.Month(month1), day1, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				case 3:
					hour = 17
					minute = 10
				case 4:
					hour = 20
					minute = 00
				}

			}

		case 2:
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year1, time.Month(month1), day1, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				}

			}
		case 3:
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year1, time.Month(month1), day1, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				case 3:
					hour = 17
					minute = 10
				}
			}
		case 4:
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year1, time.Month(month1), day1, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				case 3:
					hour = 17
					minute = 10
				}

			}
		case 5:
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year1, time.Month(month1), day1, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				}
			}

		}
	}
	//2日目
	for i := 1; i <= 5; i++ {

		hour = baseH
		minute = baseM

		switch i {
		case 1:
			for k := 1; k <= 5; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year2, time.Month(month2), day2, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				case 3:
					hour = 17
					minute = 10
				case 4:
					hour = 20
					minute = 00
				}

			}

		case 2:
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year2, time.Month(month2), day2, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				}

			}
		case 3:
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year2, time.Month(month2), day2, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				case 3:
					hour = 17
					minute = 10
				}
			}
		case 4:
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year2, time.Month(month2), day2, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				case 3:
					hour = 17
					minute = 10
				}

			}
		case 5:
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year2, time.Month(month2), day2, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				}
			}

		}
	}
	//3日目
	for i := 1; i <= 5; i++ {

		hour = baseH
		minute = baseM

		switch i {
		case 1:
			for k := 1; k <= 5; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year3, time.Month(month3), day3, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				case 3:
					hour = 17
					minute = 10
				case 4:
					hour = 20
					minute = 00
				}

			}

		case 2:
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year3, time.Month(month3), day3, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				}

			}
		case 3:
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year3, time.Month(month3), day3, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				case 3:
					hour = 17
					minute = 10
				}
			}
		case 4:
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year3, time.Month(month3), day3, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				case 3:
					hour = 17
					minute = 10
				}

			}
		case 5:
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year3, time.Month(month3), day3, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				}
			}

		}
	}
	//4日目
	for i := 1; i <= 5; i++ {

		hour = baseH
		minute = baseM

		switch i {
		case 1:
			for k := 1; k <= 5; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year4, time.Month(month4), day4, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				case 3:
					hour = 17
					minute = 10
				case 4:
					hour = 20
					minute = 00
				}

			}

		case 2:
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year4, time.Month(month4), day4, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				}

			}
		case 3:
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year4, time.Month(month4), day4, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				case 3:
					hour = 17
					minute = 10
				}
			}
		case 4:
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year4, time.Month(month4), day4, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				case 3:
					hour = 17
					minute = 10
				}

			}
		case 5:
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year4, time.Month(month4), day4, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				}
			}

		}
	}
	//5日目
	for i := 1; i <= 5; i++ {

		hour = baseH
		minute = baseM

		switch i {
		case 1:
			for k := 1; k <= 5; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year5, time.Month(month5), day5, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				case 3:
					hour = 17
					minute = 10
				case 4:
					hour = 20
					minute = 00
				}

			}

		case 2:
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year5, time.Month(month5), day5, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				}

			}
		case 3:
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year5, time.Month(month5), day5, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				case 3:
					hour = 17
					minute = 10
				}
			}
		case 4:
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year5, time.Month(month5), day5, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				case 3:
					hour = 17
					minute = 10
				}

			}
		case 5:
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year5, time.Month(month5), day5, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				}
			}

		}
	}
	//6日目
	for i := 1; i <= 5; i++ {

		hour = baseH
		minute = baseM

		switch i {
		case 1:
			for k := 1; k <= 5; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year6, time.Month(month6), day6, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				case 3:
					hour = 17
					minute = 10
				case 4:
					hour = 20
					minute = 00
				}

			}

		case 2:
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year6, time.Month(month6), day6, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				}

			}
		case 3:
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year6, time.Month(month6), day6, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				case 3:
					hour = 17
					minute = 10
				}
			}
		case 4:
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year6, time.Month(month6), day6, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				case 3:
					hour = 17
					minute = 10
				}

			}
		case 5:
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year6, time.Month(month6), day6, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				}
			}

		}
	}
	//7日目
	for i := 1; i <= 5; i++ {

		hour = baseH
		minute = baseM

		switch i {
		case 1:
			for k := 1; k <= 5; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year7, time.Month(month7), day7, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				case 3:
					hour = 17
					minute = 10
				case 4:
					hour = 20
					minute = 00
				}

			}

		case 2:
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year7, time.Month(month7), day7, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				}

			}
		case 3:
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year7, time.Month(month7), day7, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				case 3:
					hour = 17
					minute = 10
				}
			}
		case 4:
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year7, time.Month(month7), day7, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				case 3:
					hour = 17
					minute = 10
				}

			}
		case 5:
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year7, time.Month(month7), day7, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				switch k {
				case 1:
					hour = 12
					minute = 15
				case 2:
					hour = 14
					minute = 50
				}
			}

		}
	}

	fmt.Print("5×7日分を生成しました")
}
