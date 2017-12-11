package model

func (m Movie) GetSchedule() []ScreeningSchedule {
	var Schedules []ScreeningSchedule
	db.Where("movie_id = ? AND `release` = ?", m.ID, true).Find(&Schedules)
	return Schedules
}
