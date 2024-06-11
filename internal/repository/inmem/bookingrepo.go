package inmem

import (
	"booking-service/internal/domain"
	"booking-service/internal/e"
	"booking-service/internal/logger"
	"booking-service/internal/utils"
	"context"
	"math/rand"
	"sync"
	"time"
)

type BookingInMemoryRepo struct {
	orders           []domain.Order
	roomAvailability []domain.RoomAvailability
	mu               sync.RWMutex
	*logger.Logger
}

func NewBookingInMemoryRepo(logger *logger.Logger) *BookingInMemoryRepo {
	var availability = []domain.RoomAvailability{
		{"reddison", "lux", utils.Date(2024, 1, 1), 1},
		{"reddison", "lux", utils.Date(2024, 1, 2), 1},
		{"reddison", "lux", utils.Date(2024, 1, 3), 1},
		{"reddison", "lux", utils.Date(2024, 1, 4), 1},
		{"reddison", "lux", utils.Date(2024, 1, 5), 0},
	}

	return &BookingInMemoryRepo{
		orders:           make([]domain.Order, 0, 2),
		roomAvailability: availability,
		Logger:           logger,
	}
}

func (repo *BookingInMemoryRepo) CreateBooking(_ context.Context, order *domain.Order) (domain.Order, error) {
	daysToBook := utils.DaysBetween(order.From(), order.To())
	unavailableDays := make(map[time.Time]struct{})
	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()

	for _, dayToBook := range daysToBook {
		for i, availability := range repo.roomAvailability {
			if availability.HotelID != order.HotelID() || availability.RoomID != order.RoomID() ||
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

	return repo.createOrder(order)
}

func (repo *BookingInMemoryRepo) createOrder(order *domain.Order) (domain.Order, error) {
	newOrder, err := domain.NewOrder(
		order.HotelID(),
		order.RoomID(),
		order.UserEmail(),
		order.From(),
		order.To(),
	)

	//generating unique id or seq id for new entity
	err = newOrder.SetID(rand.Int())
	if err != nil {
		return domain.Order{}, err
	}

	repo.orders = append(repo.orders, *newOrder)
	return *newOrder, nil
}
