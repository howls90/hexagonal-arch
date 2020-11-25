package repositories

import "hexagonal/entities"

type TripRepo interface {
	FindAll() (trps []entities.Trip, err error)
	GetById(id uint64) (t *entities.Trip, err error)
	DeleteById(id uint64) (err error)
	Create(t *entities.Trip) (err error)
}
