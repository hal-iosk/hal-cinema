package model

import "testing"

func TestMigration(t *testing.T) {
	db.AutoMigrate(&Customer{})
	db.AutoMigrate(&Administrator{})
	db.AutoMigrate(&Authority{})
	db.AutoMigrate(&Earning{})
	db.AutoMigrate(&Movie{})
	db.AutoMigrate(&Price{})
	db.AutoMigrate(&Reserve{})
	db.AutoMigrate(&ScreeningSchedule{})
}
