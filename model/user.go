package model

import (
	"github.com/jinzhu/gorm"
)

// 会員テーブル
type Customer struct {
	gorm.Model
	CustomerID       string `gorm:"not null;unique" json:"customer_id"`
	Password         string `gorm:"not null" json:"password"`
	FirstName        string `gorm:"not null" json:"first_name"`
	LastName         string `gorm:"not null" json:"last_name"`
	FirstNameRead    string `gorm:"not null" json:"first_name_read"`
	LastNameRead     string `gorm:"not null" json:"last_name_read"`
	Phone            string `gorm:"not null" json:"phone"`
	Email            string `gorm:"not null;unique" json:"email"`
	Address          string `gorm:"not null" json:"address"`
	Birthdate        string `gorm:"not null" json:"birthday"`
	Magazine         bool   `gorm:"not null" json:"magazine"`
	PointCount       int64  `gorm:"not null" json:"point_count"`
	CreditCardName   string `gorm:"not null" json:"credit_card_name"`
	CreditCardLimit  string `gorm:"not null" json:"credit_card_limit"`
	CreditCardNumber string `gorm:"not null" json:"credit_card_number"`
	SecurityCode     string `gorm:"not null" json:"security_code"`
}

// 管理者テーブル
type Administrator struct {
	gorm.Model
	AdminID     string `gorm:"not null;unique" json:"admin_id"`
	Password    string `gorm:"not null" json:"password"`
	FirstName   string `gorm:"not null" json:"first_name"`
	LastName    string `gorm:"not null" json:"last_name"`
	Phone       string `gorm:"not null" json:"phone"`
	Email       string `gorm:"not null" json:"email"`
	AuthorityID string `gorm:"not null;unique" json:"authority_id"`
}

// 権限テーブル
type Authority struct {
	gorm.Model
	AuthorityID   string `gorm:"not null;unique" json:"authority_id"`
	AuthorityName string `gorm:"not null" json:"authority_name"`
	AuthorityCode int64  `gorm:"not null" json:"authority_code"`
}

// 売上テーブル
type Earning struct {
	gorm.Model
	EarningID   string `gorm:"not null;unique" json:"earning_id"`
	ReserveID   string `gorm:"not null;unique" json:"reserve_id"`
	EarningDate string `gorm:"not null" json:"earning_date"`
}

// 映画テーブル
type Movie struct {
	gorm.Model
	MovieID         string `gorm:"not null;unique" json:"movie_id"`
	MovieName       string `gorm:"not null" json:"movie_name"`
	StartDate       string `gorm:"not null" json:"start_date"`
	EndDate         string `gorm:"not null" json:"end_date"`
	WatchTime       int64  `gorm:"not null" json:"watch_time"`
	AdministratorID string `gorm:"not null" json:"administrator_id"`
}

// 料金テーブル
type Price struct {
	gorm.Model
	PriceID                    string `gorm:"not null;unique" json:"price_id"`
	CustomerType               string `gorm:"not null" json:"customer_type"`
	Price                      int64  `gorm:"not null" json:"price"`
	LastUpdatedAdministratorID string `gorm:"not null" json:"last_updated_administrator_id"`
}

// 予約テーブル
type Reserve struct {
	gorm.Model
	ReserveID      string `gorm:"not null;unique" json:"reserve_id"`
	ReserveNumber  int64  `gorm:"not null" json:"reserve_number"`
	MovieName      string `gorm:"not null" json:"movie_name"`
	ReserveTime    string `gorm:"not null" json:"reserve_time"`
	WatchDate      string `gorm:"not null" json:"watch_date"`
	WatchStartTime string `gorm:"not null" json:"watch_start_time"`
	TicketTime     string `gorm:"not null" json:"ticket_time"`
	TheaterNumber  int64  `gorm:"not null" json:"theater_number"`
	PriceID        string `gorm:"not null;unique" json:"price_id"`
	SeatNumber     string `gorm:"not null" json:"seat_number"`
	CustomerID     string `gorm:"not null;unique" json:"customer_id"`
	CancelCheck    bool   `gorm:"not null" json:"cancel_check"`
}

// 上映スケジュールテーブル
type ScreeningSchedule struct {
	gorm.Model
	ScheduleID                 string `gorm:"not null;unique" json:"schedule_id"`
	MovieID                    string `gorm:"not null;unique" json:"movie_id"`
	WatchStartTime             string `gorm:"not null" json:"watch_start_time"`
	WatchDate                  string `gorm:"not null" json:"watch_date"`
	TheaterNumber              int64  `gorm:"not null" json:"theater_number"`
	LastUpdatedAdministratorID string `gorm:"not null" json:"last_updated_administrator_id"`
}
