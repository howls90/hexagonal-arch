package repositories

import "hexagonal/entities"

type UserRepo interface {
	FindAll() (users []entities.User, err error)
	GetById(id uint64) (user entities.User, err error)
	DeleteById(id uint64) (err error)
	Create(u *entities.User) (err error)
}
