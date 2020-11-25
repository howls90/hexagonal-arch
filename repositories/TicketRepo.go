package repositories

import "hexagonal/entities"

type TicketRepo interface {
	FindAll() (tickets []entities.Ticket, err error)
	GetById(id uint64) (t *entities.Ticket, err error)
	DeleteById(id uint64) (err error)
	Create(t *entities.Ticket) (err error)
}
