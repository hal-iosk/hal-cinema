package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// 会員テーブル
type Customer struct {
	gorm.Model
	Email            string    `gorm:"not null;unique" json:"email"`
	Password         string    `gorm:"not null" json:"password"`
	FirstName        string    `gorm:"not null" json:"first_name"`
	LastName         string    `gorm:"not null" json:"last_name"`
	FirstNameRead    string    `gorm:"not null" json:"first_name_read"`
	LastNameRead     string    `gorm:"not null" json:"last_name_read"`
	Phone            string    `gorm:"not null" json:"phone"`
	Address          string    `gorm:"not null" json:"address"`
	Birthdate        time.Time `gorm:"not null" json:"birthday"`
	Magazine         bool      `gorm:"not null" json:"magazine"`
	PointCount       int64     `gorm:"not null" json:"point_count"`
	CreditCardLimit  string    `gorm:"not null" json:"credit_card_limit"`
	CreditCardNumber string    `gorm:"not null" json:"credit_card_number"`
	SecurityCode     string    `gorm:"not null" json:"security_code"`
}

// 管理者テーブル
type Administrator struct {
	gorm.Model
	Password  string `gorm:"not null" json:"password"`
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Phone     string `gorm:"not null" json:"phone"`
	Email     string `gorm:"not null" json:"email"`
}

// 売上テーブル
type Earning struct {
	gorm.Model
	ReserveID   string `gorm:"not null;unique" json:"reserve_id"`
	EarningDate string `gorm:"not null" json:"earning_date"`
}

// 映画テーブル
type Movie struct {
	gorm.Model
	ImagePath       string    `gorm:"not null" json:"image_path"`
	MovieName       string    `gorm:"not null" json:"movie_name"`
	Details         string    `gorm:"not null" json:"details"`
	StartDate       time.Time `gorm:"not null" json:"start_date"`
	EndDate         time.Time `gorm:"not null" json:"end_date"`
	WatchTime       int64     `gorm:"not null" json:"watch_time"`
	AdministratorID string    `gorm:"not null" json:"administrator_id"`
}

// 料金テーブル
type Price struct {
	gorm.Model
	CustomerType               string `gorm:"not null" json:"customer_type"`
	Price                      int64  `gorm:"not null" json:"price"`
	LastUpdatedAdministratorID string `gorm:"not null" json:"last_updated_administrator_id"`
}

// 予約テーブル
type Reserve struct {
	gorm.Model
	PriceID     string `gorm:"not null;unique" json:"price_id"`
	SeatNumber  string `gorm:"not null" json:"seat_number"`
	CustomerID  string `gorm:"not null;unique" json:"customer_id"`
	CancelCheck bool   `gorm:"not null" json:"cancel_check"`
}

// 上映スケジュールテーブル
type ScreeningSchedule struct {
	gorm.Model
	MovieID                    string    `gorm:"not null;unique" json:"movie_id"`
	StartTime                  time.Time `gorm:"not null" json:"start_time"`
	TheaterNumber              int64     `gorm:"not null" json:"theater_number"`
	LastUpdatedAdministratorID string    `gorm:"not null" json:"last_updated_administrator_id"`
}

type Token struct {
	gorm.Model
	AdminFlag bool `gorm:"not null" json:"admin_flag"`
	UserID    uint `gorm:"not null" json:"user_id"`
}

type UserType uint
