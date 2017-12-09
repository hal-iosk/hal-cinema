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

func (self ReserveImpl) FindByScheduleID(id uint) []model.Reserve {
	var Reserves []model.Reserve
	db.Where("schedule_id = ?", id).Find(&Reserves)
	return Reserves
}

func (self ReserveImpl) CanCreate(Reserve model.Reserve) bool {
	var dbReserve []model.Reserve
	db.Where("schedule_id = ? AND seat_id = ?", Reserve.ScheduleID, Reserve.SeatID).Find(&dbReserve)
	return len(dbReserve) == 0
}

func (self ReserveImpl) Creates(Reserves []model.Reserve) []model.Reserve {
	for key, value := range Reserves {
		Reserves[key] = self.Create(value)
	}
	return Reserves
}

// 配列ないが重複しないかつデータベースにも存在しない
func (self ReserveImpl) CanCreates(Reserves []model.Reserve) bool {
	var seets []string
	for _, value := range Reserves {
		for _, s := range seets {
			if value.SeatID == s {
				return false
			}
		}
		seets = append(seets, value.SeatID)
		ok := self.CanCreate(value)
		if !ok {
			return false
		}
	}
	return true
}
