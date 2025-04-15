package services

import (
    "testing"
    "time"
    "glofox/models"
    "glofox/daos"
)

func TestCreateClasses(t *testing.T) {
    store := datastore.NewInMemoryStore()
    service := NewDefaultClassService(store)

    class := models.Class{
        Name:      "Yoga",
        StartDate: time.Now(),
        EndDate:   time.Now().AddDate(0, 0, 2),
        Capacity:  10,
    }

    err := service.CreateClasses(class)
    if err != nil {
        t.Errorf("expected no error, got %v", err)
    }

    
    for i := 0; i < 3; i++ {
        date := class.StartDate.AddDate(0, 0, i)
        result, _ := store.GetClassesByNameAndDate("Yoga", date)
        if len(result) != 1 {
            t.Errorf("expected 1 class on %v, got %d", date, len(result))
        }
    }
}
