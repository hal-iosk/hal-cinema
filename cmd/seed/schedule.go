package seed

//強敵と格闘中！(#･∀･)

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hal-iosk/hal-cinema/model"
	"github.com/hal-iosk/hal-cinema/service"
)

func CreateSchedule() {

	var year1 int
	var year1s string
	var month1
	var month1s string
	var day1 int
	var day1s string

	var year2 int
	var year2s string
	var month2 int
	var month2s string
	var day2 int
	var day2s string

	var year3 int
	var year3s string
	var month3 int
	var month3s string
	var day3 int
	var day3s string

	var year4 int
	var year4s string
	var month4 int
	var month4s string
	var day4 int
	var day4s string

	var year5 int
	var year5s string
	var month5 int
	var month5s string
	var day5 int
	var day5s string

	d1 := time.Now()
	d2 := d1.Add(24 * time.Hour)
	d3 := d2.Add(24 * time.Hour)
	d4 := d3.Add(24 * time.Hour)
	d5 := d4.Add(24 * time.Hour)

	year1 = (d1.Year())
	year1s = strconv.Itoa(year1)
	month1 = int(d1.Month())
	month1s = strconv.Itoa(month1)
	day1 = (d1.Day())
	day1s = strconv.Itoa(day1)

	year2 = (d2.Year())
	year2s = strconv.Itoa(year2)
	month2 = int(d2.Month())
	month2s = strconv.Itoa(month2)
	day2 = (d2.Day())
	day2s = strconv.Itoa(day2)

	year3 = (d3.Year())
	year3s = strconv.Itoa(year3)
	month3 = int(d3.Month())
	month3s = strconv.Itoa(month3)
	day3 = (d3.Day())
	day3s = strconv.Itoa(day3)

	year4 = (d4.Year())
	year4s = strconv.Itoa(year4)
	month4 = int(d4.Month())
	month4s = strconv.Itoa(month4)
	day4 = (d4.Day())
	day4s = strconv.Itoa(day4)

	year5 = (d5.Year())
	year5s = strconv.Itoa(year5)
	month5 = int(d5.Month())
	month5s = strconv.Itoa(month5)
	day5 = (d5.Day())
	day5s = strconv.Itoa(day5)

	fmt.Print(year1s)
	fmt.Print(month1s)
	fmt.Print(day1s)

	fmt.Print("\n", year2s)
	fmt.Print(month2s)
	fmt.Print(day2s)

	fmt.Print("\n", year3s)
	fmt.Print(month3s)
	fmt.Print(day3s)

	fmt.Print("\n", year4s)
	fmt.Print(month4s)
	fmt.Print(day4s)

	fmt.Print("\n", year5s)
	fmt.Print(month5s)
	fmt.Print(day5s)

	var MovieID int
	var StartTime string
	var TheaterNumber int
	var Releas string

	var baseH int
	var baseM int
	var hour int
	var hours string
	var minute int
	var minutes string

	baseH = 9
	baseM = 0

	for i := 1; i <= 10; i++ {

		hour = baseH
		minute = baseM

		if i == 1 {
			for k := 1; k <= 5; k++ {

				hours = strconv.Itoa(hour)
				minutes = strconv.Itoa(minute)

				MovieID = i
				StartTime = year1s + "," + month1s + "," + day1s + "," + hours + "," + minutes + "," + "0" + "," + "0" + "," + "loc"
				TheaterNumber = i
				Releas = "true"

				fmt.Println("\nMovieID:", MovieID)
				fmt.Println("StartTime:", StartTime)
				fmt.Println("TheaterNumber:", TheaterNumber)
				fmt.Println("Releas:", Releas)

				service.Schedule.Create(model.ScreeningSchedule{
					MovieID:       uint(i),
					StartTime:     time.Date(year1, month1, day1, hour, minute, 0, 0, loc),
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
				hours = strconv.Itoa(hour)
				minutes = strconv.Itoa(minute)

				MovieID = i
				StartTime = year1s + "," + month1s + "," + day1s + "," + hours + "," + minutes + "," + "0" + "," + "0" + "," + "loc"
				TheaterNumber = i
				Releas = "true"

				fmt.Println("\nMovieID:", MovieID)
				fmt.Println("StartTime:", StartTime)
				fmt.Println("TheaterNumber:", TheaterNumber)
				fmt.Println("Releas:", Releas)

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
		if i == 3 {
			for k := 1; k <= 4; k++ {
				hours = strconv.Itoa(hour)
				minutes = strconv.Itoa(minute)

				MovieID = i
				StartTime = year1s + "," + month1s + "," + day1s + "," + hours + "," + minutes + "," + "0" + "," + "0" + "," + "loc"
				TheaterNumber = i
				Releas = "true"

				fmt.Println("\nMovieID:", MovieID)
				fmt.Println("StartTime:", StartTime)
				fmt.Println("TheaterNumber:", TheaterNumber)
				fmt.Println("Releas:", Releas)

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
		/*if i == 4 {
			for k := 1; k <= 4; k++ {
				MovieID = i
				StartTime = year1s + "," + month1s + "," + day1s
				TheaterNumber = i
				Releas = "true"

				fmt.Println("\nMovieID:", MovieID)
				fmt.Println("StartTime:", StartTime)
				fmt.Println("TheaterNumber:", TheaterNumber)
				fmt.Println("Releas:", Releas)
			}
		}
		if i == 5 {
			for k := 1; k <= 3; k++ {
				MovieID = i
				StartTime = year1s + "," + month1s + "," + day1s
				TheaterNumber = i
				Releas = "true"

				fmt.Println("\nMovieID:", MovieID)
				fmt.Println("StartTime:", StartTime)
				fmt.Println("TheaterNumber:", TheaterNumber)
				fmt.Println("Releas:", Releas)
			}
		}*/

	}

	service.Schedule.Create(model.ScreeningSchedule{
		MovieID:       1,
		StartTime:     time.Date(2017, 1, 18, 0, 0, 0, 0, loc),
		TheaterNumber: 2,
		Release:       true,
	})

}
