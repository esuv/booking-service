package domain

import (
	"fmt"
	"time"
)

type Order struct {
	id        int
	hotelID   string
	roomID    string
	userEmail string
	from      time.Time
	to        time.Time
}

func NewOrder(HotelID string, RoomID string, UserEmail string, From time.Time, To time.Time) (*Order, error) {
	if HotelID == "" {
		return nil, fmt.Errorf("%w: hotel id is required", ErrRequired)
	}
	if RoomID == "" {
		return nil, fmt.Errorf("%w: room id is required", ErrRequired)
	}
	if UserEmail == "" {
		return nil, fmt.Errorf("%w: user email is required", ErrRequired)
	}
	if From.IsZero() {
		return nil, fmt.Errorf("%w: from time can not be zero", ErrNil)
	}
	if To.IsZero() {
		return nil, fmt.Errorf("%w: to time can not be zero", ErrNil)
	}

	return &Order{
		hotelID:   HotelID,
		roomID:    RoomID,
		userEmail: UserEmail,
		from:      From,
		to:        To,
	}, nil
}

// ID returns the order id.
func (p *Order) ID() int {
	return p.id
}

// SetID sets the order ID.
func (p *Order) SetID(id int) error {
	if id == 0 {
		return fmt.Errorf("%w: id is required", ErrRequired)
	}
	p.id = id
	return nil
}

// HotelID returns the hotel id.
func (p *Order) HotelID() string {
	return p.hotelID
}

// RoomID returns the room id.
func (p *Order) RoomID() string {
	return p.roomID
}

// UserEmail returns the email.
func (p *Order) UserEmail() string {
	return p.userEmail
}

// From returns the "from" time.
func (p *Order) From() time.Time {
	return p.from
}

// To returns the "to" time.
func (p *Order) To() time.Time {
	return p.to
}
