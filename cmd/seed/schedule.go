package seed

//強敵と格闘中！(#･∀･)

import (
	"time"

	"github.com/hal-iosk/hal-cinema/model"
	"github.com/hal-iosk/hal-cinema/service"
)

func CreateSchedule() {

	var year1 int
	var month1 int
	var day1 int

	//var year2 int
	//var month2 int
	//var day2 int

	//var year3 int
	//var month3 int
	//var day3 int

	//var year4 int
	//var month4 int
	//var day4 int

	//var year5 int
	//var month5 int
	//var day5 int

	d1 := time.Now()
	//d2 := d1.Add(24 * time.Hour)
	//d3 := d2.Add(24 * time.Hour)
	//d4 := d3.Add(24 * time.Hour)
	//d5 := d4.Add(24 * time.Hour)

	year1 = (d1.Year())
	month1 = int(d1.Month())
	day1 = (d1.Day())

	//year2 = (d2.Year())
	//month2 = int(d2.Month())
	//day2 = (d2.Day())

	//year3 = (d3.Year())
	//month3 = int(d3.Month())
	//day3 = (d3.Day())

	//year4 = (d4.Year())
	//month4 = int(d4.Month())
	//day4 = (d4.Day())

	//year5 = (d5.Year())
	//month5 = int(d5.Month())
	//day5 = (d5.Day())

	var baseH int
	var baseM int
	var hour int
	var minute int

	baseH = 9
	baseM = 0

	//1日目
	for i := 1; i <= 10; i++ {

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
		case 6:
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
		case 7:
			for k := 1; k <= 2; k++ {

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
				}
			}
		case 8:
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
		case 9:
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
		case 10:
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
		}
	}
	//2日目
	/*for i := 1; i <= 10; i++ {

		hour = baseH
		minute = baseM

		if i == 1 {
			for k := 1; k <= 5; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year2, time.Month(month2), day2, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

				if k == 3 {

					hour = 17
					minute = 10
				}

				if k == 4 {

					hour = 20
					minute = 00
				}

			}
		}
		if i == 2 {
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year2, time.Month(month2), day2, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

			}
		}
		if i == 3 {
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year2, time.Month(month2), day2, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

				if k == 3 {

					hour = 17
					minute = 10
				}
			}
		}
		if i == 4 {
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year2, time.Month(month2), day2, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

				if k == 3 {

					hour = 17
					minute = 10
				}

			}
		}
		if i == 5 {
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year2, time.Month(month2), day2, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}
			}
		}
		if i == 6 {
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year2, time.Month(month2), day2, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

				if k == 3 {

					hour = 17
					minute = 10
				}
			}
		}
		if i == 7 {
			for k := 1; k <= 2; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year2, time.Month(month2), day2, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

			}
		}
		if i == 8 {
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year2, time.Month(month2), day2, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

			}
		}
		if i == 9 {
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year2, time.Month(month2), day2, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

			}
		}
		if i == 10 {
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year2, time.Month(month2), day2, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

				if k == 3 {

					hour = 17
					minute = 10
				}

			}
		}
	}

	//3日目
	for i := 1; i <= 10; i++ {

		hour = baseH
		minute = baseM

		if i == 1 {
			for k := 1; k <= 5; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year3, time.Month(month3), day3, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

				if k == 3 {

					hour = 17
					minute = 10
				}

				if k == 4 {

					hour = 20
					minute = 00
				}

			}
		}
		if i == 2 {
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year3, time.Month(month3), day3, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

			}
		}
		if i == 3 {
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year3, time.Month(month3), day3, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

				if k == 3 {

					hour = 17
					minute = 10
				}
			}
		}
		if i == 4 {
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year3, time.Month(month3), day3, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

				if k == 3 {

					hour = 17
					minute = 10
				}

			}
		}
		if i == 5 {
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year3, time.Month(month3), day3, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}
			}
		}
		if i == 6 {
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year3, time.Month(month3), day3, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

				if k == 3 {

					hour = 17
					minute = 10
				}
			}
		}
		if i == 7 {
			for k := 1; k <= 2; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year3, time.Month(month3), day3, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

			}
		}
		if i == 8 {
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year3, time.Month(month3), day3, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

			}
		}
		if i == 9 {
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year3, time.Month(month3), day3, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

			}
		}
		if i == 10 {
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year3, time.Month(month3), day3, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

				if k == 3 {

					hour = 17
					minute = 10
				}

			}
		}
	}

	//4日目
	for i := 1; i <= 10; i++ {

		hour = baseH
		minute = baseM

		if i == 1 {
			for k := 1; k <= 5; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year4, time.Month(month4), day4, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

				if k == 3 {

					hour = 17
					minute = 10
				}

				if k == 4 {

					hour = 20
					minute = 00
				}

			}
		}
		if i == 2 {
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year4, time.Month(month4), day4, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

			}
		}
		if i == 3 {
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year4, time.Month(month4), day4, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

				if k == 3 {

					hour = 17
					minute = 10
				}
			}
		}
		if i == 4 {
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year4, time.Month(month4), day4, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

				if k == 3 {

					hour = 17
					minute = 10
				}

			}
		}
		if i == 5 {
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year4, time.Month(month4), day4, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}
			}
		}
		if i == 6 {
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year4, time.Month(month4), day4, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

				if k == 3 {

					hour = 17
					minute = 10
				}
			}
		}
		if i == 7 {
			for k := 1; k <= 2; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year4, time.Month(month4), day4, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

			}
		}
		if i == 8 {
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year4, time.Month(month4), day4, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

			}
		}
		if i == 9 {
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year4, time.Month(month4), day4, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

			}
		}
		if i == 10 {
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year4, time.Month(month4), day4, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

				if k == 3 {

					hour = 17
					minute = 10
				}

			}
		}
	}

	//5日目
	for i := 1; i <= 10; i++ {

		hour = baseH
		minute = baseM

		if i == 1 {
			for k := 1; k <= 5; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year5, time.Month(month5), day5, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

				if k == 3 {

					hour = 17
					minute = 10
				}

				if k == 4 {

					hour = 20
					minute = 00
				}

			}
		}
		if i == 2 {
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year5, time.Month(month5), day5, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

			}
		}
		if i == 3 {
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year5, time.Month(month5), day5, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

				if k == 3 {

					hour = 17
					minute = 10
				}
			}
		}
		if i == 4 {
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year5, time.Month(month5), day5, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

				if k == 3 {

					hour = 17
					minute = 10
				}

			}
		}
		if i == 5 {
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year5, time.Month(month5), day5, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}
			}
		}
		if i == 6 {
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year5, time.Month(month5), day5, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

				if k == 3 {

					hour = 17
					minute = 10
				}
			}
		}
		if i == 7 {
			for k := 1; k <= 2; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year5, time.Month(month5), day5, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

			}
		}
		if i == 8 {
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year5, time.Month(month5), day5, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

			}
		}
		if i == 9 {
			for k := 1; k <= 3; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year5, time.Month(month5), day5, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

			}
		}
		if i == 10 {
			for k := 1; k <= 4; k++ {

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year5, time.Month(month5), day5, hour, minute, 0, 0, loc),
					TheaterNumber: uint(i),
					Release:       true,
				})

				if k == 1 {

					hour = 12
					minute = 15
				}

				if k == 2 {

					hour = 14
					minute = 50
				}

				if k == 3 {

					hour = 17
					minute = 10
				}

			}
		}
	}

	fmt.Print("5日分を生成しました")*/
}
