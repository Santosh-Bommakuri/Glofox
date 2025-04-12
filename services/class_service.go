package services

import (
	"glofox/daos"
	"glofox/models"
)

type ClassService interface {
	CreateClasses(class models.Class) error
}

type DefaultClassService struct {
	store datastore.DataStore
}

func NewDefaultClassService(store datastore.DataStore) *DefaultClassService {
	return &DefaultClassService{store: store}
}

func (s *DefaultClassService) CreateClasses(class models.Class) error {
	currentDate := class.StartDate
	for !currentDate.After(class.EndDate) {
		classForDate := models.Class{
			Name:      class.Name,
			StartDate: currentDate,
			EndDate:   currentDate,
			Capacity:  class.Capacity,
		}
		if err := s.store.CreateClass(classForDate); err != nil {
			return err
		}
		currentDate = currentDate.AddDate(0, 0, 1)
	}
	return nil
}
