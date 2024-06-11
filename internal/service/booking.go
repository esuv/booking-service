package service

import (
	"booking-service/internal/domain"
	"booking-service/internal/logger"
	"context"
)

type BookingRepo interface {
	CreateBooking(ctx context.Context, order *domain.Order) (domain.Order, error)
}

type BookingServiceImpl struct {
	repository BookingRepo
	*logger.Logger
}

func NewBookingService(repository BookingRepo, log *logger.Logger) BookingServiceImpl {
	return BookingServiceImpl{repository: repository, Logger: log}
}

func (bs BookingServiceImpl) Book(ctx context.Context, order *domain.Order) (domain.Order, error) {
	return bs.repository.CreateBooking(ctx, order)
}
