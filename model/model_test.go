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

func TestCreateAdmin(t *testing.T) {
	err := db.Create(&Administrator{
		Password:    "hogehoge",
		FirstName:   "かずき",
		LastName:    "西川",
		Phone:       "09088888888",
		Email:       "kinokoruumu@gmail.com",
		AuthorityID: "0",
	}).Error

	if err != nil {
		t.Errorf("%v", err)
	}
}
