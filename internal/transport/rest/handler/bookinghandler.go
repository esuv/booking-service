package handler

import (
	"booking-service/internal/domain"
	"booking-service/internal/logger"
	"booking-service/internal/transport/rest"
	"booking-service/internal/transport/rest/dto"
	"encoding/json"
	"fmt"
	"net/http"
)

type BookingService interface {
	Book(order *domain.Order) (domain.Order, error)
}

// BookingHandler is a booking handler.
type BookingHandler struct {
	service BookingService
	logger  *logger.Logger
}

// NewBookingHandler creates a new booking handler.
func NewBookingHandler(service BookingService, log *logger.Logger) *BookingHandler {
	return &BookingHandler{service: service, logger: log}
}

// Post handles POST /orders requests
func (h *BookingHandler) Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("method %s unsupported", r.Method), http.StatusMethodNotAllowed)
		return
	}

	var request dto.Order
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("invalid request: %s", err), http.StatusBadRequest)
		return
	}

	order := rest.OrderHttpToDomain(&request)
	newOrder, err := h.service.Book(order)
	if err != nil {
		http.Error(w, fmt.Sprintf("booking error: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rest.OrderModelToHttp(&newOrder))

	h.logger.LogInfo("Order successfully created: %v", newOrder)
}
