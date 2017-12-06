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

func CreateMovie(c *gin.Context) model.Movie {
	movie := model.Movie{
		MovieName: c.PostForm("movie_name"),
		Details:   c.PostForm("details"),
		ImagePath: c.PostForm("image_path"),
	}
	//movie.WatchTime = strconv.Atoi(c.PostForm(""))
	//movie.StartDate
	//movie.EndDate
	return movie
}

func MovieCreate(c *gin.Context) {

}
