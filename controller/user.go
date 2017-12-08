package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hal-iosk/hal-cinema/model"
	"github.com/hal-iosk/hal-cinema/service"
)

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

func Checkin(c *gin.Context) {
	userID, _ := c.Get("userID")
	p := service.Customer.UpdatePoint(userID.(uint), 1)
	c.JSON(http.StatusOK, gin.H{
		"point": p,
	})
}
func Popcorn(c *gin.Context) {
	userID, _ := c.Get("userID")
	User := service.Customer.Find(userID.(uint))
	if User.PointCount < 3 {
		Batequest("ぽっぷこーんかえまてん＾ー＾", c)
		return
	}
	p := service.Customer.UpdatePoint(userID.(uint), -3)
	c.JSON(http.StatusOK, gin.H{
		"point": p,
	})
}
func UpdateUser(c *gin.Context) {
	userID, _ := c.Get("userID")
	req, err := CustomerReq{}.create(c)
	if err != nil {
		return
	}
	user := service.Customer.Find(userID.(uint))

	//変更させたくないもの
	req.PointCount = user.PointCount
	res := service.Customer.Update(userID.(uint), *req)
	c.JSON(http.StatusOK, res)
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
		Birthdate:        req.Birthdate,
		Magazine:         req.Magazine,
		PointCount:       0,
		CreditCardLimit:  req.CreditCardLimit,
		CreditCardNumber: req.CreditCardNumber,
		SecurityCode:     req.SecurityCode,
	})
	c.JSON(http.StatusCreated, "ok!!")
}

func (self CustomerReq) create(c *gin.Context) (*model.Customer, error) {
	var err error
	self.Email = c.PostForm("email")
	self.Password = c.PostForm("password")
	self.FirstName = c.PostForm("first_name")
	self.LastName = c.PostForm("last_name")
	self.FirstNameRead = c.PostForm("first_name_read")
	self.LastNameRead = c.PostForm("last_name_read")
	self.Phone = c.PostForm("phone")
	self.Address = c.PostForm("address")
	self.Birthdate, err = GetDate("birthday", c)
	if err != nil {
		Batequest("日付けのフォーマットがおかしいよん", c)
		return nil, err
	}
	self.Magazine = toBool(c.PostForm("magazine"))
	self.CreditCardLimit = c.PostForm("credit_card_limit")
	self.CreditCardNumber = c.PostForm("credit_card_number")
	self.SecurityCode = c.PostForm("security_code")

	Customer := model.Customer{
		Email:            self.Email,
		Password:         self.Password,
		FirstName:        self.FirstName,
		LastName:         self.LastName,
		FirstNameRead:    self.FirstNameRead,
		LastNameRead:     self.LastNameRead,
		Phone:            self.Phone,
		Address:          self.Address,
		Birthdate:        self.Birthdate,
		Magazine:         self.Magazine,
		PointCount:       0,
		CreditCardLimit:  self.CreditCardLimit,
		CreditCardNumber: self.CreditCardNumber,
		SecurityCode:     self.SecurityCode,
	}
	return &Customer, nil
}

func toBool(string string) bool {
	return string == "true"
}
