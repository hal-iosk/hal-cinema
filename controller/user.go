package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hal-iosk/hal-cinema/model"
	"github.com/hal-iosk/hal-cinema/service"
)

var DayFormat = "2006/01/02"
var TimeFormat = "2006-01-02 15:04:05"

type CustomerReq struct {
	Email            string    `gorm:"not null;unique" json:"email"`
	Password         string    `gorm:"not null" json:"password"`
	FirstName        string    `gorm:"not null" json:"first_name"`
	LastName         string    `gorm:"not null" json:"last_name"`
	FirstNameRead    string    `gorm:"not null" json:"first_name_read"`
	LastNameRead     string    `gorm:"not null" json:"last_name_read"`
	Phone            string    `gorm:"not null" json:"phone"`
	Address          string    `gorm:"not null" json:"address"`
	Birthdate        time.Time `json:"birthday"`
	Magazine         bool      `gorm:"not null" json:"magazine"`
	CreditCardLimit  string    `gorm:"not null" json:"credit_card_limit"`
	CreditCardNumber string    `gorm:"not null" json:"credit_card_number"`
	SecurityCode     string    `gorm:"not null" json:"security_code"`
}

func CreateUser(c *gin.Context) {
	req, err := CustomerReq{}.create(c)
	if err != nil {
		return
	}
	if service.Customer.Exists(req.Email) {
		Batequest("メアド使用済み", c)
		return
	}
	service.Customer.Create(model.Customer{
		Email:            req.Email,
		Password:         req.Password,
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		FirstNameRead:    req.FirstNameRead,
		LastNameRead:     req.LastNameRead,
		Phone:            req.Phone,
		Address:          req.Address,
		Birthdate:        &req.Birthdate,
		Magazine:         req.Magazine,
		PointCount:       0,
		CreditCardLimit:  req.CreditCardLimit,
		CreditCardNumber: req.CreditCardNumber,
		SecurityCode:     req.SecurityCode,
	})
	c.JSON(200, "ok!!")
}

func (self CustomerReq) create(c *gin.Context) (CustomerReq, error) {
	var err error
	self.Email = c.Query("email")
	self.Password = c.Query("password")
	self.FirstName = c.Query("first_name")
	self.LastName = c.Query("last_name")
	self.FirstNameRead = c.Query("first_name_read")
	self.LastNameRead = c.Query("last_name_read")
	self.Phone = c.Query("phone")
	self.Address = c.Query("address")
	self.Birthdate, err = time.Parse(DayFormat, c.Query("birthday"))
	if err != nil {
		Batequest("日付けのフォーマットがおかしいよん", c)
		return self, err
	}
	self.Magazine = toBool(c.Query("magazine"))
	self.CreditCardLimit = c.Query("credit_card_limit")
	self.CreditCardNumber = c.Query("credit_card_number")
	self.SecurityCode = c.Query("security_code")
	return self, nil
}

func toBool(string string) bool {
	return string == "true"
}
