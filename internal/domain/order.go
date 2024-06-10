package domain

import "time"

type Order struct {
	ID        int
	HotelID   string
	RoomID    string
	UserEmail string
	From      time.Time
	To        time.Time
}
