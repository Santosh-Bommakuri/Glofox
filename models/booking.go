package models

import "time"

type Booking struct {
	Name       string    `json:"name"`
	Class      string    `json:"class"`
	BookedDate time.Time `json:"date"`
}
