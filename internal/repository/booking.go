package repository

import (
	"booking-service/internal/domain"
	"booking-service/internal/e"
	"booking-service/internal/logger"
	"booking-service/internal/utils"
	"math/rand"
	"sync"
	"time"
)

type BookingRepo struct {
	orders           []domain.Order
	roomAvailability []domain.RoomAvailability
	mu               sync.RWMutex
	*logger.Logger
}

func NewBookingInMemoryRepo(logger *logger.Logger) *BookingRepo {
	var availability = []domain.RoomAvailability{
		{"reddison", "lux", utils.Date(2024, 1, 1), 1},
		{"reddison", "lux", utils.Date(2024, 1, 2), 1},
		{"reddison", "lux", utils.Date(2024, 1, 3), 1},
		{"reddison", "lux", utils.Date(2024, 1, 4), 1},
		{"reddison", "lux", utils.Date(2024, 1, 5), 0},
	}

	return &BookingRepo{
		roomAvailability: availability,
		Logger:           logger,
	}
}

func (repo *BookingRepo) CreateBooking(order *domain.Order) (domain.Order, error) {
	defer repo.mu.Unlock()
	repo.mu.Lock()

	daysToBook := utils.DaysBetween(order.From, order.To)
	unavailableDays := make(map[time.Time]struct{})
	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}

	for _, dayToBook := range daysToBook {
		for i, availability := range repo.roomAvailability {
			if availability.HotelID != order.HotelID || availability.RoomID != order.RoomID ||
				!availability.Date.Equal(dayToBook) || availability.Quota < 1 {
				continue
			}

			availability.Quota -= 1
			repo.roomAvailability[i] = availability
			delete(unavailableDays, dayToBook)
		}
	}

	if len(unavailableDays) != 0 {
		repo.Logger.LogErrorf("Hotel room is not available for selected dates:\n%v\n%v", order, unavailableDays)
		return domain.Order{}, e.BookingError
	}

	return repo.createOrder(order), nil
}

func (repo *BookingRepo) createOrder(order *domain.Order) domain.Order {
	newOrder := domain.Order{
		ID:        rand.Int(),
		HotelID:   order.HotelID,
		RoomID:    order.RoomID,
		UserEmail: order.UserEmail,
		From:      order.From,
		To:        order.To,
	}

	repo.orders = append(repo.orders, newOrder)
	return newOrder
}
