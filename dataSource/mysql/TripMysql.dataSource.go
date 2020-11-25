package dataSource

import (
	"hexagonal/entities"
	"time"
)

type TripMsql struct {
	Id          uint64 `gorm:"primaryKey"`
	Title       string `gorm:"not null;varchar(50)"`
	Description string `gorm:"not null;varchar(100)"`
	Departure   time.Time
	Arrival     time.Time
	Seats       uint16
	CreatedAt   time.Time
}

func (tm TripMsql) FindAll() (trips []entities.Trip, err error) {
	if err := DB.Select("id, title, description, departure, arrival, seats").Order("id").Find(&trips).Error; err != nil {
		return nil, err
	}
	return trips, nil
}

func (tm TripMsql) GetById(id uint64) (t *entities.Trip, err error) {
	if err := DB.Select("id, title, description, departure, arrival, seats").First(&t, id).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (tm TripMsql) Create(t *entities.Trip) (err error) {
	if err := DB.Create(&t).Error; err != nil {
		return err
	}
	return nil
}

func (tm TripMsql) DeleteById(id uint64) (err error) {
	if DB.First(&tm, id).Error != nil {
		return
	}
	if err := DB.Unscoped().Delete(&tm).Error; err != nil {
		return err
	}

	return nil
}
