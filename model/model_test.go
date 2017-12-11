package model

import (
	"testing"
)

func TestMigration(t *testing.T) {
	db.AutoMigrate(&Customer{})
	db.AutoMigrate(&Administrator{})
	db.AutoMigrate(&Movie{})
	db.AutoMigrate(&Price{})
	db.AutoMigrate(&Reserve{})
	db.AutoMigrate(&ScreeningSchedule{})
	db.AutoMigrate(&Token{})
}

func TestReserveStore(t *testing.T) {

}
