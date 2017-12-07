package service

import "github.com/hal-iosk/hal-cinema/model"

var Reserve = ReserveImpl{}

type ReserveImpl struct {
}

func (self ReserveImpl) Create(Reserve model.Reserve) model.Reserve {
	err := db.Create(&Reserve).Error
	if err != nil {
		panic(err)
	}
	return Reserve
}

func (self ReserveImpl) Find(id uint) *model.Reserve {
	Reserve := []model.Reserve{}
	db.Find(&Reserve, id)
	if len(Reserve) == 0 {
		return nil
	}
	return &Reserve[0]
}
func (self ReserveImpl) Update(id uint, Reserve model.Reserve) *model.Reserve {
	Reserve.ID = id
	db.Model(&Reserve).Update(&Reserve)
	return &Reserve
}

func (self ReserveImpl) Delete(id uint) {
	Reserve := model.Reserve{}
	db.Delete(&Reserve, id)
}

func (self ReserveImpl) All() []model.Reserve {
	var Reserves []model.Reserve
	db.Find(&Reserves)
	return Reserves
}

func (self ReserveImpl) Exists(id uint) bool {
	return self.Find(id) != nil
}

func (self ReserveImpl) CanCreate(Reserve model.Reserve) bool {
	var dbReserve []model.Reserve
	db.Where("schedule_id = ? AND seat_id = ?", Reserve.ScheduleID, Reserve.SeatID).Find(&dbReserve)
	return len(dbReserve) == 0
}
