package service

import (
	"github.com/hal-iosk/hal-cinema/model"
)

var Customer = customerImpl{}

type customerImpl struct {
}

func (c customerImpl) Create(customer model.Customer) {
	err := db.Create(&customer).Error
	if err != nil {
		panic(err)
	}
}
