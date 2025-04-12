
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"glofox/models"
	"glofox/services"
)


type BookingHandler struct {
	bookingService services.BookingService
}

func NewBookingHandler(bookingService services.BookingService) *BookingHandler {
	return &BookingHandler{bookingService: bookingService}
}


func (h *BookingHandler) BookClass(c *gin.Context) {
	var bookingRequest models.Booking
	if err := c.ShouldBindJSON(&bookingRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	if bookingRequest.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required field: name"})
		return
	}
	if bookingRequest.Class == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required field: class"})
		return
	}
	if bookingRequest.BookedDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required field: date"})
		return
	}

	err := h.bookingService.BookClass(bookingRequest)
	if err != nil {
		if err == services.ErrClassNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to book class", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Booking successful"})
}