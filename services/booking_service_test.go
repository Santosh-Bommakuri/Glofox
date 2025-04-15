package services

import (
    "testing"
    "time"
    "glofox/models"
    "glofox/daos"
)

func TestBookClass(t *testing.T) {
    store := datastore.NewInMemoryStore()
    classService := NewDefaultClassService(store)
    bookingService := NewDefaultBookingService(store)

    class := models.Class{
        Name:      "Zumba",
        StartDate: time.Now(),
        EndDate:   time.Now(),
        Capacity:  5,
    }


    err := classService.CreateClasses(class)
    if err != nil {
        t.Fatal(err)
    }

    booking := models.Booking{
        Name:       "Santhosh",
        Class:      "Zumba",
        BookedDate: time.Now(),
    }

    err = bookingService.BookClass(booking)
    if err != nil {
        t.Errorf("expected booking to succeed, got %v", err)
    }
}
