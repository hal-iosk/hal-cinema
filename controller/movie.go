package controller

import (
	"net/http"

	"strconv"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/hal-iosk/hal-cinema/model"
	"github.com/hal-iosk/hal-cinema/service"
)

const (
	OnAir = iota
	NotOnAir
)

type MovieRes struct {
	model.Movie
	Schedules []model.ScreeningSchedule `json:"schedules"`
}

func GetMovieAll(c *gin.Context) {
	MoviesAll := service.Movie.All()
	var Movies []model.Movie
	var Day *time.Time
	status, err := strconv.Atoi(c.Query("status"))
	strDay, ok := c.GetQuery("day")
	if ok {
		d, err := time.Parse(DayFormat, strDay)
		if err != nil {
			Batequest("dayがおかしいよ", c)
			return
		}
		Day = &d
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"movies": createMovieRes(MoviesAll, Day),
		})
		return
	}
	switch status {
	case OnAir:
		for _, value := range MoviesAll {
			if value.EndDate.After(time.Now()) {
				if isOnAir(value.StartDate) {
					Movies = append(Movies, value)
				}
			}
		}
		break
	case NotOnAir:
		for _, value := range MoviesAll {
			if value.EndDate.After(time.Now()) {
				if !isOnAir(value.StartDate) {
					Movies = append(Movies, value)
				}
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"movies": createMovieRes(Movies, Day),
	})
}
func isOnAir(startTime time.Time) bool {
	return startTime.Before(time.Now())
}
func createMovieRes(m []model.Movie, day *time.Time) []MovieRes {
	var res []MovieRes
	for _, value := range m {
		res = append(res, MovieRes{
			Movie:     value,
			Schedules: value.GetSchedule(),
		})
	}
	if day == nil {
		return res
	}
	for key, r := range res {
		var Schedules []model.ScreeningSchedule
		for _, s := range r.Schedules {
			if isSomeDay(*day, s.StartTime) {
				Schedules = append(Schedules, s)
			}
		}
		res[key].Schedules = Schedules
	}
	return res
}
func isSomeDay(time1 time.Time, time2 time.Time) bool {
	return time1.Year() == time2.Year() && time1.Day() == time2.Day() && time1.Month() == time2.Month()
}
func GetMovie(c *gin.Context) {
	movieID, err := strconv.Atoi(c.Param("movie_id"))
	if err != nil {
		Batequest("数値で頼むわ", c)
		return
	}
	movie := service.Movie.Find(uint(movieID))
	if movie == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "んなもんねーよ！笑",
		})
		return
	}
	c.JSON(http.StatusOK, movie)
}

func CreateMovie(c *gin.Context) {
	movie, ok := createMovieReq(c)
	if !ok {
		return
	}
	movie = service.Movie.Create(movie)
	c.JSON(http.StatusCreated, movie)
}

func UpdateMovie(c *gin.Context) {
	movieID, err := strconv.Atoi(c.Param("movie_id"))
	if err != nil {
		Batequest("数値で頼むわ", c)
		return
	}
	if !service.Movie.Exists(uint(movieID)) {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "んなもんねーよ！笑",
		})
	}
	movie, ok := createMovieReq(c)
	if !ok {
		return
	}
	movieRes := service.Movie.Update(uint(movieID), movie)
	c.JSON(http.StatusOK, movieRes)
}

func DeleteMovie(c *gin.Context) {
	movieID, err := strconv.Atoi(c.Param("movie_id"))
	if err != nil {
		Batequest("数値で頼むわ", c)
		return
	}
	if !service.Movie.Exists(uint(movieID)) {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "んなもんねーよ！笑",
		})
	}
	service.Movie.Delete(uint(movieID))
	c.Writer.WriteHeader(http.StatusNoContent)
}
func createMovieReq(c *gin.Context) (model.Movie, bool) {
	movie := model.Movie{
		MovieName: c.PostForm("movie_name"),
		Details:   c.PostForm("details"),
		ImagePath: c.PostForm("image_path"),
	}
	watchTime, err := strconv.Atoi(c.PostForm("watch_time"))
	if err != nil {
		Batequest("watch_time 数字入れろカス〜", c)
		return movie, false
	}
	movie.WatchTime = uint(watchTime)

	movie.StartDate, err = GetDate("start_date", c)
	if err != nil {
		panic(err)
		Batequest("start_date がちゃうで", c)
		return movie, false
	}
	movie.EndDate, err = GetDate("end_date", c)
	if err != nil {
		Batequest("end_date がちゃうで", c)
		return movie, false
	}

	return movie, true
}
