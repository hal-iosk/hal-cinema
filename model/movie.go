package model

func (m Movie) GetSchedule() []ScreeningSchedule {
	var Schedules []ScreeningSchedule
	db.Where("movie_id = ?", m.ID).Find(&Schedules)
	return Schedules
}
