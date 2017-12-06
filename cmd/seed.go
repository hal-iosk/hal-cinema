package main

import (
	"time"

	"github.com/hal-iosk/hal-cinema/model"
	"github.com/hal-iosk/hal-cinema/service"
)

var db = model.GetDBConn()

func main() {
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

	service.Admin.Create(model.Administrator{
		Email:     "hoge@gmail.com",
		Password:  "hoge",
		FirstName: "葛巻",
		LastName:  "大樹",
		Phone:     "09096220908",
	})

	service.Movie.Create(model.Movie{
		MovieName: "君の名は",
		Details:   "君の名はだよー",
		StartDate: time.Date(2017, 12, 2, 0, 0, 0, 0, loc),
		EndDate:   time.Date(2018, 2, 3, 0, 0, 0, 0, loc),
		ImagePath: "https://rr.img.naver.jp/mig?src=http%3A%2F%2Fwww.kiminona.com%2Fimages%2Fcommon%2Fog_image.jpg&twidth=1000&theight=0&qlt=80&res_format=jpg&op=r",
	})

	service.Movie.Create(model.Movie{
		MovieName: "君の名は",
		Details:   "君の名はだよー",
		StartDate: time.Date(2017, 12, 2, 0, 0, 0, 0, loc),
		EndDate:   time.Date(2018, 2, 3, 0, 0, 0, 0, loc),
		ImagePath: "https://rr.img.naver.jp/mig?src=http%3A%2F%2Fwww.kiminona.com%2Fimages%2Fcommon%2Fog_image.jpg&twidth=1000&theight=0&qlt=80&res_format=jpg&op=r",
	})

	service.Movie.Create(model.Movie{
		MovieName: "君の名は",
		Details:   "君の名はだよー",
		StartDate: time.Date(2017, 12, 2, 0, 0, 0, 0, loc),
		EndDate:   time.Date(2018, 2, 3, 0, 0, 0, 0, loc),
		ImagePath: "https://rr.img.naver.jp/mig?src=http%3A%2F%2Fwww.kiminona.com%2Fimages%2Fcommon%2Fog_image.jpg&twidth=1000&theight=0&qlt=80&res_format=jpg&op=r",
	})

	service.Movie.Create(model.Movie{
		MovieName: "君の名は",
		Details:   "君の名はだよー",
		StartDate: time.Date(2017, 12, 2, 0, 0, 0, 0, loc),
		EndDate:   time.Date(2018, 2, 3, 0, 0, 0, 0, loc),
		ImagePath: "https://rr.img.naver.jp/mig?src=http%3A%2F%2Fwww.kiminona.com%2Fimages%2Fcommon%2Fog_image.jpg&twidth=1000&theight=0&qlt=80&res_format=jpg&op=r",
	})

}
