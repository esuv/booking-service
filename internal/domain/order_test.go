package domain

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewOrder(t *testing.T) {
	t.Parallel()

	hotelID := "hotel id"
	roomID := "room id"
	userEmail := "email"
	from := time.Now()
	to := time.Now()

	t.Run("valid", func(t *testing.T) {
		order, err := NewOrder(hotelID, roomID, userEmail, from, to)
		require.NoError(t, err)

		require.Equal(t, hotelID, order.HotelID())
		require.Equal(t, roomID, order.RoomID())
		require.Equal(t, userEmail, order.UserEmail())
		require.Equal(t, from, order.From())
		require.Equal(t, to, order.To())
	})

	t.Run("missing hotel id", func(t *testing.T) {
		_, err := NewOrder("", roomID, userEmail, from, to)
		require.Error(t, err)
	})

	t.Run("missing room id", func(t *testing.T) {
		_, err := NewOrder(hotelID, "", userEmail, from, to)
		require.Error(t, err)
	})

	t.Run("missing email", func(t *testing.T) {
		_, err := NewOrder(hotelID, roomID, "", from, to)
		require.Error(t, err)
	})

	t.Run("missing 'from' time", func(t *testing.T) {
		_, err := NewOrder(hotelID, roomID, userEmail, time.Time{}, to)
		require.Error(t, err)
	})

	t.Run("missing 'to' time", func(t *testing.T) {
		_, err := NewOrder(hotelID, roomID, userEmail, from, time.Time{})
		require.Error(t, err)
	})
}
