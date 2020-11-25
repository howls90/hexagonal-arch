package dataSource

import (
	"hexagonal/entities"
	"time"
)

type TicketMsql struct {
	Id        uint64 `gorm:"primaryKey"`
	UserId    uint64
	TripId    uint64
	CreatedAt time.Time
}

func (tm TicketMsql) FindAll() (tickets []entities.Ticket, err error) {
	if err := DB.Select("id, name, surname, email").Order("id").Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}

func (tm TicketMsql) GetById(id uint64) (t *entities.Ticket, err error) {
	if err := DB.Select("id, name, surname, email").First(&t, id).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (um TicketMsql) Create(t *entities.Ticket) (err error) {
	if err := DB.Create(&t).Error; err != nil {
		return err
	}
	return nil
}

func (tm TicketMsql) DeleteById(id uint64) (err error) {
	if DB.First(&tm, id).Error != nil {
		return
	}
	if err := DB.Unscoped().Delete(&tm).Error; err != nil {
		return err
	}

	return nil
}
