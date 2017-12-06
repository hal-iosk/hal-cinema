package service

import (
	"github.com/hal-iosk/hal-cinema/model"
)

var Customer = customerImpl{}

type customerImpl struct {
}

func (c customerImpl) Create(customer model.Customer) {
	customer.Password = hash(customer.Password)
	err := db.Create(&customer).Error
	if err != nil {
		panic(err)
	}
}
func (c customerImpl) Exists(email string) bool {
	var cus []model.Customer
	db.Where("email = ?", email).Find(&cus)
	return len(cus) != 0
}

func (c customerImpl) Logincheck(email, password string) bool {
	var customer model.Customer
	db.Where("email = ?", email).First(&customer)
	if customer.Email == "" {
		return false
	}
	return customer.Password == hash(password)
}
func hash(s string) string {
	return s
}
