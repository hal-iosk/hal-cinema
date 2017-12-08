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
func (self customerImpl) Update(id uint, Customer model.Customer) *model.Customer {
	Customer.ID = id
	db.Model(&Customer).Update(&Customer)
	return &Customer
}

func (self customerImpl) Find(id uint) *model.Customer {
	Customer := []model.Customer{}
	db.Find(&Customer, id)
	if len(Customer) == 0 {
		return nil
	}
	return &Customer[0]
}
func (c customerImpl) Exists(email string) bool {
	var cus []model.Customer
	db.Where("email = ?", email).Find(&cus)
	return len(cus) != 0
}

func (c customerImpl) Logincheck(email, password string) (*model.Customer, bool) {
	var customer model.Customer
	db.Where("email = ?", email).First(&customer)
	if customer.Email == "" {
		return nil, false
	}
	return &customer, customer.Password == hash(password)
}
func hash(s string) string {
	return s
}
