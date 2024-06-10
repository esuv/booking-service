package rest

import (
	"booking-service/internal/domain"
	"booking-service/internal/transport/rest/dto"
)

func OrderHttpToDomain(dto *dto.Order) *domain.Order {
	return &domain.Order{
		HotelID:   dto.HotelID,
		RoomID:    dto.RoomID,
		UserEmail: dto.UserEmail,
		From:      dto.From,
		To:        dto.To,
	}
}

func OrderModelToHttp(domain *domain.Order) *dto.Order {
	return &dto.Order{
		ID:        domain.ID,
		HotelID:   domain.HotelID,
		RoomID:    domain.RoomID,
		UserEmail: domain.UserEmail,
		From:      domain.From,
		To:        domain.To,
	}
}
