package service

import "github.com/hal-iosk/hal-cinema/model"

var Schedule = scheduleImpl{}

type scheduleImpl struct {
}

func (self scheduleImpl) Create(Schedule model.ScreeningSchedule) model.ScreeningSchedule {
	err := db.Create(&Schedule).Error
	if err != nil {
		panic(err)
	}
	return Schedule
}

func (self scheduleImpl) Find(id uint) *model.ScreeningSchedule {
	Schedule := []model.ScreeningSchedule{}
	db.Find(&Schedule, id)
	if len(Schedule) == 0 {
		return nil
	}
	return &Schedule[0]
}
func (self scheduleImpl) Update(id uint, Schedule model.ScreeningSchedule) *model.ScreeningSchedule {
	Schedule.ID = id
	db.Model(&Schedule).Update(&Schedule)
	return &Schedule
}

func (self scheduleImpl) Delete(id uint) {
	Schedule := model.ScreeningSchedule{}
	db.Delete(&Schedule, id)
}

func (self scheduleImpl) All() []model.ScreeningSchedule {
	var Schedules []model.ScreeningSchedule
	db.Find(&Schedules)
	return Schedules
}

func (self scheduleImpl) Exists(id uint) bool {
	return self.Find(id) != nil
}

func (self scheduleImpl) FindByMovie(id uint) []model.ScreeningSchedule {
	var Schedules []model.ScreeningSchedule
	db.Where("movie_id = ?", id).Find(&Schedules)
	return Schedules
}
