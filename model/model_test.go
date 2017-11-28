package model

import "testing"

func TestMigration(t *testing.T) {
	db.AutoMigrate(&Customer{})
	db.AutoMigrate(&Administrator{})
	db.AutoMigrate(&Authority{})
	db.AutoMigrate(&Earning{})
	db.AutoMigrate(&Movie{})
	db.AutoMigrate(&Price{})
	db.AutoMigrate(&Reserve{})
	db.AutoMigrate(&ScreeningSchedule{})
}

func TestReserve(t *testing.T) {
	file := &Reserve{
		ReserveNumber:  1,
		MovieName:      "君の名は。",
		ReserveTime:    "2016-06-15:12:00",
		WatchDate:      "2016-06-15",
		WatchStartTime: "15:00",
		TicketTime:     "2016-06-15:15:09",
		TheaterNumber:  1,
		PriceID:        "学生",
		SeatNumber:     "A-1",
		CustomerID:     "1",
		CancelCheck:    false,
	}
}
