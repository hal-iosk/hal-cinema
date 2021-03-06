package seed

import (
	"time"

	"github.com/hal-iosk/hal-cinema/model"
	"github.com/hal-iosk/hal-cinema/service"
)

var db = model.GetDBConn()
var loc, _ = time.LoadLocation("Asia/Tokyo")

func Seed() {

	//db.Exec("INSERT INTO `authorities` (`authority_name`, `authority_code`) VALUES ('なし', 404);")
	//db.Exec("INSERT INTO `authorities` (`authority_name`, `authority_code`) VALUES ('ルート', 777);")
	//db.Exec("INSERT INTO `authorities` (`authority_name`, `authority_code`) VALUES ('管理者登録', 400);")
	//db.Exec("INSERT INTO `authorities` (`authority_name`, `authority_code`) VALUES ('メルマガ送信', 123);")
	//db.Exec("INSERT INTO `authorities` (`authority_name`, `authority_code`) VALUES ('スケジュール管理', 345);")
	//db.Exec("INSERT INTO `authorities` (`authority_name`, `authority_code`) VALUES ('チケット取り消し', 500);")
	//db.Exec("INSERT INTO `authorities` (`authority_name`, `authority_code`) VALUES ('売上管理', 100);")
	//
	//db.Exec("INSERT INTO `prices` (`customer_type`, `price`, `last_updated_administrator_id`) VALUES ('一般', 1800, '0');")
	//db.Exec("INSERT INTO `prices` (`customer_type`, `price`, `last_updated_administrator_id`) VALUES ('学生（大高生）', 1500, '0');")
	//db.Exec("INSERT INTO `prices` (`customer_type`, `price`, `last_updated_administrator_id`) VALUES ('小中学生', 800, '0');")
	//db.Exec("INSERT INTO `prices` (`customer_type`, `price`, `last_updated_administrator_id`) VALUES ('幼児', 500, '0');")
	//db.Exec("INSERT INTO `prices` (`customer_type`, `price`, `last_updated_administrator_id`) VALUES ('3D専用メガネ代', 300, '0');")

	service.Customer.Create(model.Customer{
		Email:            "hoge@gmail.com",
		Password:         "hoge",
		FirstName:        "葛巻",
		LastName:         "大樹",
		FirstNameRead:    "かつらまき",
		LastNameRead:     "たいき",
		Phone:            "09096220908",
		Address:          "高槻",
		Birthdate:        time.Now(),
		Magazine:         true,
		PointCount:       0,
		CreditCardLimit:  "20/15",
		CreditCardNumber: "4242 4242 4242 4242",
		SecurityCode:     "564",
	})
	service.Admin.Create(model.Administrator{
		Email:     "hoge@gmail.com",
		Password:  "hoge",
		FirstName: "葛巻",
		LastName:  "大樹",
		Phone:     "09096220908",
	})
	CreateMovie()
	CreateSchedule()
}
