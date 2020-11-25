package dataSource

import (
	"hexagonal/entities"
	"time"

	"github.com/prometheus/common/log"
)

type UserMsql struct {
	Id        uint64 `gorm:"primaryKey"`
	Name      string `gorm:"not null;varchar(50)"`
	Surname   string `gorm:"not null;varchar(50)"`
	Email     string `gorm:"not null;varchar(50)"`
	CreatedAt time.Time
}

func (um UserMsql) FindAll() (users []entities.User, err error) {
	if err := DB.Select("id, name, surname, email").Order("id").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (um UserMsql) GetById(id uint64) (user entities.User, err error) {
	if err := DB.Select("id, name, surname, email").First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (um UserMsql) Create(u *entities.User) (err error) {
	if err := DB.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (um UserMsql) DeleteById(id uint64) (err error) {
	if DB.First(&um, id).Error != nil {
		log.Error(err)
		return
	}
	if err := DB.Unscoped().Delete(&um).Error; err != nil {
		log.Error(err)
		return err
	}

	return nil
}
