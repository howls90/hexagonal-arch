package entities

import "time"

type Trip struct {
	Id          uint64    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Departure   time.Time `json:"departure"`
	Arrival     time.Time `json:"arrival"`
	Seats       uint16    `json:"seats"`
	CreatedAt   time.Time `json:"created_at"`
}
