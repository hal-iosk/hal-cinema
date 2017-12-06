package service

import (
	"github.com/hal-iosk/hal-cinema/model"
)

var Admin = adminImpl{}

type adminImpl struct {
}

func (c adminImpl) Create(admin model.Administrator) {
	admin.Password = hash(admin.Password)
	err := db.Create(&admin).Error
	if err != nil {
		panic(err)
	}
}
func (c adminImpl) Exists(email string) bool {
	var admins []model.Administrator
	db.Where("email = ?", email).Find(&admins)
	return len(admins) != 0
}

func (c adminImpl) Logincheck(email, password string) (*model.Administrator, bool) {
	var admin model.Administrator
	db.Where("email = ?", email).First(&admin)
	if admin.Email == "" {
		return nil, false
	}
	return &admin, admin.Password == hash(password)
}
