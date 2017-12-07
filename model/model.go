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
	Birthdate        time.Time `json:"birthday"`
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

// 映画テーブル
type Movie struct {
	gorm.Model
	MovieName string    `gorm:"not null" json:"movie_name"`
	Details   string    `gorm:"not null" json:"details"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	ImagePath string    `gorm:"not null" json:"image_path"`
	WatchTime uint      `gorm:"not null" json:"watch_time"`
}

// 料金テーブル
type Price struct {
	gorm.Model
	CustomerType               string `gorm:"not null" json:"customer_type"`
	Price                      int64  `gorm:"not null" json:"price"`
	LastUpdatedAdministratorID uint   `gorm:"not null" json:"last_updated_administrator_id"`
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
	MovieID       uint      `gorm:"not null" json:"movie_id"`
	StartTime     time.Time `json:"start_time"`
	TheaterNumber uint      `gorm:"not null" json:"theater_number"`
	Release       bool      `json:"release"`
}

type Token struct {
	gorm.Model
	AdminFlag bool   `gorm:"not null" json:"admin_flag"`
	UserID    uint   `gorm:"not null" json:"user_id"`
	Token     string `gorm:"not null;unique" json:"token"`
}

type UserType uint
