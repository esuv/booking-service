package handler

import (
	"booking-service/internal/domain"
	"booking-service/internal/logger"
	"booking-service/internal/service"
	"bytes"
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

type HttpTestSuite struct {
	suite.Suite
	service    BookingService
	httpServer BookingHandler
}

func NewHttpTestSuite() *HttpTestSuite {
	suite := &HttpTestSuite{}
	l := logger.NewLogger()

	// create order mock repository
	repository := new(BookingMockRepo)
	repository.On("CreateBooking", mock.Anything, mock.Anything).Return(stubOrder(), nil)

	// create order service
	suite.service = service.NewBookingService(repository, l)

	// create http server with application injected
	suite.httpServer = NewBookingHandler(suite.service, l)

	return suite
}

func TestHttpTestSuite(t *testing.T) {
	suite.Run(t, NewHttpTestSuite())
}

func (suite *HttpTestSuite) TestCreateBooking() {
	createOrderRequest, err := os.ReadFile("testfixtures/create_order_request.json")
	require.NoError(suite.T(), err)

	createOrderResponse, err := os.ReadFile("testfixtures/create_order_response.json")
	require.NoError(suite.T(), err)

	// create POST /orders request
	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(createOrderRequest))

	w := httptest.NewRecorder()

	// run request
	suite.httpServer.Post(w, req)

	res := w.Result()
	defer res.Body.Close()

	// read response body
	data, err := io.ReadAll(res.Body)
	require.NoError(suite.T(), err)

	require.Equal(suite.T(), http.StatusCreated, res.StatusCode)
	require.Equal(suite.T(), createOrderResponse, data)
}

type BookingMockRepo struct {
	mock.Mock
}

func (m *BookingMockRepo) CreateBooking(_ context.Context, order *domain.Order) (domain.Order, error) {
	args := m.Called(order)
	return stubOrder(), args.Error(1)
}

func stubOrder() domain.Order {
	from, _ := time.Parse(time.RFC3339, "2024-01-02T00:00:00Z")
	to, _ := time.Parse(time.RFC3339, "2024-01-04T00:00:00Z")
	order, _ := domain.NewOrder("reddison",
		"lux",
		"guest@mail.ru",
		from,
		to,
	)
	_ = order.SetID(4423797040789215841)
	return *order
}
