package service

import (
	"testing"

	"github.com/hal-iosk/hal-cinema/model"
)

func TestCustomerImpl_Create(t *testing.T) {
	Customer.Create(model.Customer{})
}
