package datastore

import (
	"time"
	"glofox/models"
)

type DataStore interface {
	CreateClass(class models.Class) error
	GetClassesByNameAndDate(name string, date time.Time) ([]models.Class, error)
	AddBooking(booking models.Booking) error
	GetBookingsByClassAndDate(class string, date time.Time) ([]models.Booking, error)
}
