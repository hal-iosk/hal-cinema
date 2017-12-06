package controller

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hal-iosk/hal-cinema/model"
	"github.com/hal-iosk/hal-cinema/service"
)

func GetMovieAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"movies": service.Movie.All(),
	})
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
	c.JSON(http.StatusOK, movie)
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

func DeleteMovie() {

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
