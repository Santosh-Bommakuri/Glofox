package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"glofox/daos"
	"glofox/handlers"
	"glofox/services"
)

func main() {
	r := gin.Default()

	store := datastore.NewInMemoryStore()

	classService := services.NewDefaultClassService(store)
	bookingService := services.NewDefaultBookingService(store)

	classHandler := handlers.NewClassHandler(classService)
	bookingHandler := handlers.NewBookingHandler(bookingService)

	r.POST("/classes", classHandler.CreateClass)
	r.POST("/bookings", bookingHandler.BookClass)

	fmt.Println("Server listening on port 8080")
	log.Fatal(r.Run(":9091"))
}
