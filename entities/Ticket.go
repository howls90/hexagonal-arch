package entities

import "time"

type Ticket struct {
	UserId    uint64    `json:"user_id"`
	TripId    uint64    `json:"trip_id"`
	CreatedAt time.Time `json:"created_at"`
}
