package main

import "github.com/hal-iosk/hal-cinema/model"

func main() {
	db := model.GetDBConn()
	//db.LogMode(true)
	db.DropTableIfExists(&model.Customer{})
	db.DropTableIfExists(&model.Administrator{})
	db.DropTableIfExists(&model.Movie{})
	db.DropTableIfExists(&model.Price{})
	db.DropTableIfExists(&model.Reserve{})
	db.DropTableIfExists(&model.ScreeningSchedule{})
	db.DropTableIfExists(&model.Token{})

	db.AutoMigrate(&model.Customer{})
	db.AutoMigrate(&model.Administrator{})
	db.AutoMigrate(&model.Movie{})
	db.AutoMigrate(&model.Price{})
	db.AutoMigrate(&model.Reserve{})
	db.AutoMigrate(&model.ScreeningSchedule{})
	db.AutoMigrate(&model.Token{})
}
