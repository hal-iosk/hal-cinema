package service

import "github.com/hal-iosk/hal-cinema/model"

var Movie = movieImpl{}

type movieImpl struct {
}

func (self movieImpl) Create(movie model.Movie) model.Movie {
	err := db.Create(&movie).Error
	if err != nil {
		panic(err)
	}
	return movie
}

func (self movieImpl) Find(id uint) *model.Movie {
	movie := []model.Movie{}
	db.Find(&movie, id)
	if len(movie) == 0 {
		return nil
	}
	return &movie[0]
}
func (self movieImpl) Update(id uint, movie model.Movie) *model.Movie {
	movie.ID = id
	db.Model(&movie).Update(&movie)
	return &movie
}

func (self movieImpl) Delete(id uint) {
	movie := model.Movie{}
	db.Delete(&movie, id)
}

func (self movieImpl) All() []model.Movie {
	var movies []model.Movie
	db.Find(&movies)
	return movies
}

func (self movieImpl) Exists(id uint) bool {
	return self.Find(id) != nil
}
