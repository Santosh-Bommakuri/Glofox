package services

import (
	"glofox/daos"
	"glofox/models"
)

type BookingService interface {
	BookClass(booking models.Booking) error
}

type DefaultBookingService struct {
	store datastore.DataStore
}

func NewDefaultBookingService(store datastore.DataStore) *DefaultBookingService {
	return &DefaultBookingService{store: store}
}

func (s *DefaultBookingService) BookClass(booking models.Booking) error {
	classes, err := s.store.GetClassesByNameAndDate(booking.Class, booking.BookedDate)
	if err != nil || len(classes) == 0 {
		return ErrClassNotFound
	}
	return s.store.AddBooking(booking)
}

var ErrClassNotFound = errorString("class not found on the specified date")

type errorString string

func (e errorString) Error() string {
	return string(e)
}
