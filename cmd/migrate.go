package main

import "github.com/hal-iosk/hal-cinema/model"

func main() {
	db := model.GetDBConn()

	db.DropTable(&model.Customer{})
	db.DropTable(&model.Administrator{})
	db.DropTable(&model.Earning{})
	db.DropTable(&model.Movie{})
	db.DropTable(&model.Price{})
	db.DropTable(&model.Reserve{})
	db.DropTable(&model.ScreeningSchedule{})
	db.DropTable(&model.Token{})

	//db.AutoMigrate(&model.Customer{})
	//db.AutoMigrate(&model.Administrator{})
	//db.AutoMigrate(&model.Earning{})
	//db.AutoMigrate(&model.Movie{})
	//db.AutoMigrate(&model.Price{})
	//db.AutoMigrate(&model.Reserve{})
	//db.AutoMigrate(&model.ScreeningSchedule{})
	//db.AutoMigrate(&model.Token{})
}
