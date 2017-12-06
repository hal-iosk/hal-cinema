package service

import (
	"testing"

	"time"

	"fmt"

	"github.com/hal-iosk/hal-cinema/model"
)

var loc, _ = time.LoadLocation("Asia/Tokyo")

func TestMovieImpl_Create(t *testing.T) {
	Movie.Create(model.Movie{
		MovieName: "君の名は",
		Details:   "君の名はだよー",
		StartDate: time.Date(2017, 12, 2, 0, 0, 0, 0, loc),
		EndDate:   time.Date(2018, 2, 3, 0, 0, 0, 0, loc),
		ImagePath: "https://rr.img.naver.jp/mig?src=http%3A%2F%2Fwww.kiminona.com%2Fimages%2Fcommon%2Fog_image.jpg&twidth=1000&theight=0&qlt=80&res_format=jpg&op=r",
	})
}

func TestMovieImpl_Find(t *testing.T) {
	movie := Movie.Find(1)
	fmt.Println(movie)
}

func TestMovieImpl_Update(t *testing.T) {
	movie := Movie.Update(1, model.Movie{
		MovieName: "君の名は?",
		Details:   "君の名はだよー",
		StartDate: time.Date(2017, 12, 2, 0, 0, 0, 0, loc),
		EndDate:   time.Date(2018, 2, 5, 0, 0, 0, 0, loc),
		ImagePath: "https://rr.img.naver.jp/mig?src=http%3A%2F%2Fwww.kiminona.com%2Fimages%2Fcommon%2Fog_image.jpg&twidth=1000&theight=0&qlt=80&res_format=jpg&op=r",
	})
	fmt.Println(movie)
}

func TestMovieImpl_Delete(t *testing.T) {
	Movie.Delete(1)
}

func TestMovieImpl_All(t *testing.T) {
	movies := Movie.All()
	fmt.Println(movies)
}
