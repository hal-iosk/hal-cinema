package model

import (
	"time"
)

type User struct {
	Id        int64
	Name      string `sql:"size:255"`
	Password  string
	CreatedAt time.Time
	UpdatedAT time.Time
	DeletedAt time.Time
}
