package service

import (
	"booking-service/internal/domain"
	"booking-service/internal/logger"
)

type Booking interface {
	CreateBooking(order *domain.Order) (domain.Order, error)
}

type BookingServiceImpl struct {
	repository Booking
	*logger.Logger
}

func NewBookingService(repository Booking, log *logger.Logger) BookingServiceImpl {
	return BookingServiceImpl{repository: repository, Logger: log}
}

func (bs *BookingServiceImpl) Book(order *domain.Order) (domain.Order, error) {
	return bs.repository.CreateBooking(order)
}
