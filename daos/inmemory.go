package datastore

import (
	"sync"
	"time"
	"fmt"

	"glofox/models"
)


type InMemoryStore struct {
	classes  map[string]map[string][]models.Class   
	bookings map[string]map[string][]models.Booking 
	mu       sync.RWMutex
}

var instance *InMemoryStore
var once sync.Once


func NewInMemoryStore() *InMemoryStore {
	once.Do(func() {
		instance = &InMemoryStore{
			classes:  make(map[string]map[string][]models.Class),
			bookings: make(map[string]map[string][]models.Booking),
		}
	})
	return instance
}

func (s *InMemoryStore) CreateClass(class models.Class) error {
	s.mu.Lock() 
	defer s.mu.Unlock()
	dateKey := class.StartDate.Format("2006-01-02")
	if _, ok := s.classes[class.Name]; !ok {
		s.classes[class.Name] = make(map[string][]models.Class)
	}
	s.classes[class.Name][dateKey] = append(s.classes[class.Name][dateKey], class)
	return nil
}

func (s *InMemoryStore) GetClassesByNameAndDate(name string, date time.Time) ([]models.Class, error) {
	s.mu.RLock() 
	defer s.mu.RUnlock()
	dateKey := date.Format("2006-01-02")
	if _, ok := s.classes[name]; ok {
		return s.classes[name][dateKey], nil
	}
	return nil, nil
}

func (s *InMemoryStore) AddBooking(booking models.Booking) error {
	s.mu.Lock() 
	defer s.mu.Unlock()
	dateKey := booking.BookedDate.Format("2006-01-02")
	if _, ok := s.bookings[booking.Class]; !ok {
		s.bookings[booking.Class] = make(map[string][]models.Booking)
	}
	s.bookings[booking.Class][dateKey] = append(s.bookings[booking.Class][dateKey], booking)
	fmt.Println(s.bookings)

	return nil
}

func (s *InMemoryStore) GetBookingsByClassAndDate(class string, date time.Time) ([]models.Booking, error) {
	s.mu.RLock() 
	defer s.mu.RUnlock()
	dateKey := date.Format("2006-01-02")
	if _, ok := s.bookings[class]; ok {
		return s.bookings[class][dateKey], nil
	}
	return nil, nil
}