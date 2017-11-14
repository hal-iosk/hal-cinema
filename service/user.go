package service

import "github.com/hal-iosk/hal-cinema/model"

func CreateUser(name string, password string) error {
	return db.Create(&model.Administrator{}).Error
}
