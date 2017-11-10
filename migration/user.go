package main

import (
	"github.com/hal-iosk/hal-cinema/model"
)

func main() {
	db := model.GetDBConn()

	// user table
	db.DropTable(&model.User{})
	db.CreateTable(&model.User{})
}
