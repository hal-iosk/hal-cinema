package service

import (
	"testing"

	"fmt"

	"github.com/hal-iosk/hal-cinema/model"
)

func TestAdminImpl_Create(t *testing.T) {
	Admin.Create(model.Administrator{
		Email:     "llxo2_5oxll@icloud.com",
		Password:  "hoge",
		FirstName: "葛巻",
		LastName:  "大樹",
		Phone:     "09096220908",
	})
}

func TestAdminImpl_Logincheck(t *testing.T) {
	_, ok := Admin.Logincheck("llxo2_5oxll@icloud.com", "hoge")
	fmt.Println(ok)
}
