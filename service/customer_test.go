package service

import (
	"testing"

	"time"

	"fmt"

	"github.com/hal-iosk/hal-cinema/model"
)

func TestCustomerImpl_Create(t *testing.T) {
	time := time.Now()
	Customer.Create(model.Customer{
		Email:            "llxo2_5oxll@icloud.com2",
		Password:         "hoge",
		FirstName:        "葛巻",
		LastName:         "大樹",
		FirstNameRead:    "かつらまき",
		LastNameRead:     "たいき",
		Phone:            "09096220908",
		Address:          "高槻",
		Birthdate:        time,
		Magazine:         true,
		PointCount:       0,
		CreditCardLimit:  "20/15",
		CreditCardNumber: "4242 4242 4242 4242",
		SecurityCode:     "564",
	})
}

func TestCustomerImpl_Logincheck(t *testing.T) {
	user, ok := Customer.Logincheck("llxo2_5oxll@icloud.com", "hoge")
	fmt.Println(user, ok)
}
func TestCustomerImpl_UpdatePoint(t *testing.T) {
	Customer.UpdatePoint(1, -12)
}
